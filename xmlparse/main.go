package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type stack struct {
	items []string
}

func (s *stack) push(in string) {
	log.Printf("push %s", in)

	s.items = append(s.items, in)
}

func (s *stack) pop() (res string) {
	if len(s.items) == 0 {
		log.Fatal("pop empty stack")
	}

	res, s.items = s.items[len(s.items)-1], s.items[:len(s.items)-1]

	log.Printf("pop %s", res)
	return
}

const (
	unknown = iota
	stateSTag
	stateETag
	stateIgnoreAttr
)

func main() {
	in := bufio.NewReader(os.Stdin)
	//in := bytes.NewBufferString(`<CD><TITLE area="greatest-hits">Greatest Hits</TITLE><ARTIST>Dolly Parton</ARTIST><COUNTRY>USA</COUNTRY><COMPANY>RCA</COMPANY><PRICE>9.90</PRICE><YEAR>1982</YEAR></CD>`)

	var current strings.Builder
	var s stack
	state := unknown
	log.Print("BOF")

	for {
		ch, size, err := in.ReadRune()
		if err == io.EOF {
			log.Print("EOF")
			break
		} else if size != 1 {
			log.Fatalf("unexpected size %d", size)
		} else if err != nil {
			log.Fatal(err)
		}

		switch ch {
		case '<':
			if state == stateSTag {
				log.Fatal("invalid start tag")
			}
			state = stateSTag

		case '>':
			switch state {
			case stateSTag, stateIgnoreAttr:
				s.push(current.String())
			case stateETag:
				pop := s.pop()
				if pop != current.String() {
					log.Fatalf("unmatched tag %s != %s", pop, current.String())
				}
			default:
				log.Fatalf("unknown state %d", state)
			}

			state = unknown
			current.Reset()

		case '/':
			state = stateETag

		case ' ':
			switch state {
			case stateSTag:
				state = stateIgnoreAttr
			}

		default:
			switch state {
			case stateSTag, stateETag:
				current.WriteRune(ch)
			}
		}
	}
}
