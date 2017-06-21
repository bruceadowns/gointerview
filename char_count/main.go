package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	/*
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic("error: " + err.Error())
		}
		//fmt.Println(string(b))
	*/

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		c := make(map[int]int)
		for _, v := range line {
			if i, ok := c[int(v)]; ok {
				c[int(v)] = i + 1
			} else {
				c[int(v)] = 1
			}
		}

		//s := []int{}
		var i int
		//s := make([]int)
		//s := make([]int, len(c))
		l := 0
		ccap := len(c)
		s := make([]int, l, ccap)
		//s := make([]int, 0, len(c))
		for k := range c {
			s = append(s, k)
			//s[i] = k
			i++
			//fmt.Printf("%s %d\n", string(k), v)
		}

		sort.Ints(s)
		for _, v := range s {
			fmt.Printf("%s%d", string(v), c[v])
		}
		fmt.Println()

	}

}
