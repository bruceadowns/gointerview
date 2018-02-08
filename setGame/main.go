package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// https://www.amazon.com/SET-Family-Game-Visual-Perception/dp/B00000IV34
// https://www.setgame.com/set/puzzle

type color int

const (
	red = iota
	purple
	green
)

type shape int

const (
	circle = iota
	wave
	diamond
)

type number int

type shade int

const (
	empty = iota
	half
	full
)

type card struct {
	color  color
	shape  shape
	number number
	shade  shade
}

func (c card) String() string {
	buf := bytes.Buffer{}

	switch c.color {
	case red:
		buf.WriteString("red,")
	case purple:
		buf.WriteString("purple,")
	case green:
		buf.WriteString("green,")
	}

	switch c.shape {
	case circle:
		buf.WriteString("circle,")
	case wave:
		buf.WriteString("wave,")
	case diamond:
		buf.WriteString("diamond,")
	}

	buf.WriteString(fmt.Sprintf("%d,", c.number))

	switch c.shade {
	case empty:
		buf.WriteString("clean")
	case half:
		buf.WriteString("half")
	case full:
		buf.WriteString("full")
	}

	return buf.String()
}

func build(s string) (res card) {
	fields := strings.Split(s, ",")
	if len(fields) != 4 {
		log.Fatalf("invalid input")
	}

	switch fields[0] {
	case "red":
		res.color = red
	case "purple":
		res.color = purple
	case "green":
		res.color = green
	default:
		log.Fatal("invalid input")
	}

	switch fields[1] {
	case "circle":
		res.shape = circle
	case "wave":
		res.shape = wave
	case "diamond":
		res.shape = diamond
	default:
		log.Fatal("invalid input")
	}

	n, err := strconv.Atoi(fields[2])
	if err != nil {
		log.Fatal("invalid input")
	}
	res.number = number(n)

	switch fields[3] {
	case "empty":
		res.shade = empty
	case "half":
		res.shade = half
	case "full":
		res.shade = full
	default:
		log.Fatal("invalid input")
	}

	return
}

func isSet(c1, c2, c3 card) bool {
	if ((c1.color == c2.color && c1.color == c3.color) || (c1.color != c2.color && c1.color != c3.color && c2.color != c3.color)) &&
		((c1.shape == c2.shape && c1.shape == c3.shape) || (c1.shape != c2.shape && c1.shape != c3.shape && c2.shape != c3.shape)) &&
		((c1.number == c2.number && c1.number == c3.number) || (c1.number != c2.number && c1.number != c3.number && c2.number != c3.number)) &&
		((c1.shade == c2.shade && c1.shade == c3.shade) || (c1.shade != c2.shade && c1.shade != c3.shade && c2.shade != c3.shade)) {
		return true
	}

	return false
}

func main() {
	buf := bytes.NewBufferString(`
green,circle,2,full
green,circle,1,empty
red,circle,1,full
green,circle,2,empty
green,circle,1,full
purple,circle,3,full
red,wave,1,full
purple,diamond,2,full
red,wave,3,full
green,circle,2,half
purple,wave,3,full
green,circle,3,full
`)

	cards := make([]card, 0)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			cards = append(cards, build(line))
		}
	}
	//log.Printf("cards: %v", cards)

	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			for k := j + 1; k < len(cards); k++ {
				c1, c2, c3 := cards[i], cards[j], cards[k]
				if isSet(c1, c2, c3) {
					log.Printf("Set: %s %s %s", c1, c2, c3)
				}
			}
		}
	}
}
