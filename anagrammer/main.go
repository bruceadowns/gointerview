package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
	"sync"
	"unicode"
)

var wordDict map[string]struct{}

func initDict(s string) {
	wordDict = make(map[string]struct{})
	f, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wordDict[strings.ToLower(scanner.Text())] = struct{}{}
	}
}

func permute(iterable []rune) (res [][]rune) {
	pool := iterable
	n := len(pool)

	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	cycles := make([]int, n)
	for i := range cycles {
		cycles[i] = n - i
	}

	result := make([]rune, n)
	for i, el := range indices[:n] {
		result[i] = pool[el]
	}

	results := [][]rune{result}

	for n > 0 {
		i := n - 1
		for ; i >= 0; i -= 1 {
			cycles[i] -= 1
			if cycles[i] == 0 {
				index := indices[i]
				for j := i; j < n-1; j += 1 {
					indices[j] = indices[j+1]
				}
				indices[n-1] = index
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indices[i], indices[n-j] = indices[n-j], indices[i]

				result := make([]rune, n)
				for k := 0; k < n; k += 1 {
					result[k] = pool[indices[k]]
				}

				results = append(results, result)

				break
			}
		}

		if i < 0 {
			return results
		}
	}

	return nil
}

func strip(s string) (res []rune) {
	res = make([]rune, 0)
	for _, v := range s {
		if !unicode.IsSpace(v) {
			res = append(res, v)
		}
	}

	return
}

func genLetterPerms(s string) (out chan []rune) {
	out = make(chan []rune, 100)

	go func() {
		defer close(out)

		p := strip(s)
		log.Printf("phrase: %s [%d]", string(p), len(p))
		for _, v := range permute(p) {
			out <- v
		}
	}()

	return out
}

func words(s []rune) (res [][]rune) {
	res = make([][]rune, 0)
	for i := 0; i < len(s); i++ {
		w1 := s[:i+1]
		w2 := s[i+1:]

		if _, ok := wordDict[string(w1)]; ok {
			if len(w2) == 0 {
				log.Printf("0: %d %s %s %d", i, string(w1), string(w2), len(res))

				res = append(res, w1)
			} else {
				log.Printf("1: %d %s %s %d", i, string(w1), string(w2), len(res))

				for _, v := range words(w2) {
					l1 := len(w1)
					l2 := len(v)

					buf := make([]rune, l1+l2+1)
					for j := 0; j < l1; j++ {
						buf[j] = s[:i+1][j]
					}
					buf[l1+1] = ' '
					for j := 0; j < l2; j++ {
						buf[l1+j+1] = v[j]
					}

					log.Printf("add %s", string(buf))
					res = append(res, buf)
				}
			}
		}
	}

	return
}

func genWordPerms(in chan []rune) (out chan []rune) {
	out = make(chan []rune, 100)
	wg := sync.WaitGroup{}

	//for cpu := 0; cpu < runtime.NumCPU()-1; cpu++ {
	for cpu := 0; cpu < 1; cpu++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for p := range in {
				for _, w := range words(p) {
					out <- w
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func processResults(perms <-chan []rune) {
	count := 0
	for p := range perms {
		count++
		log.Print(string(p))
	}
	log.Printf("count: %d", count)

	return
}

func main() {
	var dictFile, phrase string
	flag.StringVar(&dictFile, "dict", "/usr/share/dict/words", "Dictionary of words")
	flag.StringVar(&phrase, "phrase", "new relic jive", "Phrase to anagram")
	flag.Parse()

	log.Printf("Dictionary: %s", dictFile)
	log.Printf("Phrase: %s", phrase)

	initDict(dictFile)
	log.Printf("Number of dictionary words: %d", len(wordDict))

	s1 := strip("bl ack")
	log.Printf("%s", string(s1))
	for k, v := range words(s1) {
		log.Printf("%d: %s", k, string(v))
	}
	return

	// pipeline
	lperms := genLetterPerms(phrase)
	wperms := genWordPerms(lperms)
	processResults(wperms)

	log.Printf("Done")
}
