package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type token struct {
	data      string
	attrs     map[string]string
	tokenType uint8
}

type node struct {
	name     string
	attrs    map[string]string
	body     string
	parent   *node
	children []*node
}

const (
	tokenSTag = iota
	tokenContent
	tokenETag
	tokenSETag
	tokenEof
)

const (
	stateUnknown = iota
	stateInSTag
	stateInSTagName
	stateInSTagNamed
	stateInETag
	stateInETagName
	stateInETagNamed
	stateInSETag
	stateInAttrName
	stateInAttrNamed
	stateInAttrValue
	stateInContent
)

func readToken(r io.RuneScanner) (res *token, err error) {
	state := stateUnknown

	var data, attrName, attrValue strings.Builder
	attrs := make(map[string]string)

outer:
	for {
		var ch rune
		var size int
		ch, size, err = r.ReadRune()
		if err == io.EOF {
			switch state {
			case stateInContent:
				res = &token{data: data.String(), tokenType: tokenEof}
			default:
				res = &token{tokenType: tokenEof}
			}
			err = nil
			break
		} else if size != 1 {
			err = fmt.Errorf("unexpected character size")
			break
		} else if err != nil {
			break
		}

		switch ch {
		case '<':
			switch state {
			case stateUnknown:
				state = stateInSTag
			case stateInContent:
				err = r.UnreadRune()

				res = &token{data: strings.TrimSpace(data.String()), tokenType: tokenContent}
				break outer
			default:
				err = fmt.Errorf("invalid start tag")
				break outer
			}

		case '>':
			switch state {
			case stateInSTag:
				err = fmt.Errorf("empty tag")
				break outer
			case stateInSTagName, stateInSTagNamed:
				res = &token{data: data.String(), attrs: attrs, tokenType: tokenSTag}
				break outer
			case stateInETagName, stateInETagNamed:
				res = &token{data: data.String(), attrs: attrs, tokenType: tokenETag}
				break outer
			case stateInSETag:
				res = &token{data: data.String(), attrs: attrs, tokenType: tokenSETag}
				break outer
			case stateInAttrName, stateInAttrNamed:
				attrs[attrName.String()] = ""
				res = &token{data: data.String(), attrs: attrs, tokenType: tokenSTag}
				break outer
			default:
				err = fmt.Errorf("invalid end tag")
				break outer
			}

		case '"':
			switch state {
			case stateUnknown, stateInContent:
				state = stateInContent
				data.WriteRune(ch)
			case stateInAttrNamed:
				state = stateInAttrValue
			case stateInAttrValue:
				attrs[attrName.String()] = attrValue.String()
				attrName.Reset()
				attrValue.Reset()
				state = stateInSTagNamed
			default:
				err = fmt.Errorf("unexpected quote")
				break outer
			}

		case '/':
			switch state {
			case stateUnknown, stateInContent:
				state = stateInContent
				data.WriteRune(ch)
			case stateInSTag:
				state = stateInETag
			case stateInSTagName:
				state = stateInSTagNamed
			case stateInETagName:
				state = stateInETagNamed
			case stateInSTagNamed:
				state = stateInSETag
			case stateInAttrValue:
				attrValue.WriteRune(ch)
			default:
				err = fmt.Errorf("invalid slash")
				break outer
			}

		case '=':
			switch state {
			case stateUnknown:
				state = stateInContent
				data.WriteRune(ch)
			case stateInContent:
				data.WriteRune(ch)
			case stateInAttrValue:
				attrValue.WriteRune(ch)
			case stateInAttrName:
				state = stateInAttrNamed
			default:
				err = fmt.Errorf("invalid equal")
				break outer
			}

		case ' ':
			switch state {
			case stateUnknown, stateInSTag, stateInETag, stateInSTagNamed, stateInAttrNamed:
			case stateInContent:
				data.WriteRune(ch)
			case stateInAttrValue:
				attrValue.WriteRune(ch)
			case stateInSTagName:
				state = stateInSTagNamed
			case stateInETagName:
				state = stateInETagNamed
			case stateInAttrName:
				state = stateInAttrNamed
			default:
				err = fmt.Errorf("invalid space")
				break outer
			}

		default:
			switch state {
			case stateUnknown:
				state = stateInContent
				data.WriteRune(ch)
			case stateInSTag:
				state = stateInSTagName
				data.WriteRune(ch)
			case stateInETag:
				state = stateInETagName
				data.WriteRune(ch)
			case stateInContent, stateInSTagName, stateInETagName:
				data.WriteRune(ch)
			case stateInSTagNamed, stateInAttrName:
				state = stateInAttrName
				attrName.WriteRune(ch)
			case stateInAttrValue:
				attrValue.WriteRune(ch)
			default:
				err = fmt.Errorf("unexpected state %d", state)
				break outer
			}
		}
	}

	return
}

func main() {
	in := bufio.NewReader(os.Stdin)
	//in := bytes.NewBufferString(`<html><body><a href="/hello"><h1>Hello There!</a><p>I love parsing stuff</p><img src="homerbushes.gif"></body></html>`)

	root := &node{name: "root"}
	current := root

	log.Print("BOF")

	satags := map[string]struct{}{
		"img":  {},
		"meta": {},
		"link": {},
	}

outer:
	for {
		token, err := readToken(in)
		if err != nil {
			log.Fatal(err)
		}

		switch token.tokenType {
		case tokenSTag:
			log.Printf("start tag %s", token.data)
			current = &node{name: token.data, attrs: token.attrs, parent: current}

		case tokenContent:
			if len(token.data) == 0 {
				log.Print("no content")
			} else {
				log.Printf("content %s", token.data)
				current.body = token.data
			}

		case tokenETag:
			if token.data != current.name {
				if _, ok := satags[current.name]; !ok {
					log.Fatalf("invalid end tag %s != %s", token.data, current.name)
				}
				current = current.parent
			}

			log.Printf("end tag %s", token.data)
			current = current.parent

		case tokenSETag:
			log.Printf("start end tag %s", token.data)
			current.children = append(current.children,
				&node{name: token.data, attrs: token.attrs, parent: current})

		case tokenEof:
			if len(token.data) != 0 {
				log.Printf("content: %s", token.data)
				current = &node{name: token.data, attrs: token.attrs, parent: current}
			}
			break outer
		}
	}

	log.Print("EOF")
}
