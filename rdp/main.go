package main

// Lexical Analyzer and Recursive Descent Parser
// https://en.wikipedia.org/wiki/Recursive_descent_parser
// https://en.wikipedia.org/wiki/PL/0

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

type itemType int

const (
	itemEOF itemType = iota

	itemBecomes
	itemComma
	itemDiv
	itemDot
	itemEqual
	itemGt
	itemGte
	itemHash
	itemIdent
	itemLparan
	itemLt
	itemLte
	itemMinus
	itemMul
	itemNot
	itemNumber
	itemPlus
	itemQuestion
	itemRparan
	itemSemicolon

	itemBegin
	itemCall
	itemConst
	itemDo
	itemEnd
	itemIf
	itemOdd
	itemProcedure
	itemThen
	itemVar
	itemWhile
	itemWrite
)

var keywords = map[string]itemType{
	"begin":     itemBegin,
	"call":      itemCall,
	"const":     itemConst,
	"do":        itemDo,
	"end":       itemEnd,
	"if":        itemIf,
	"odd":       itemOdd,
	"procedure": itemProcedure,
	"then":      itemThen,
	"var":       itemVar,
	"while":     itemWhile,
	"write":     itemWrite,
}

type lexItem struct {
	t itemType
	v string
}

type lexer struct {
	bb    []byte
	start int
	pos   int
	ch    chan lexItem
}

func (l *lexer) accept(b byte) bool {
	if !l.next() {
		return false
	}

	if l.curr() == b {
		return true
	}

	l.backup()
	return false
}

func (l *lexer) acceptRun(s string) {
	eol := false
	for strings.IndexByte(s, l.curr()) != -1 {
		if !l.next() {
			eol = true
			break
		}
	}

	if !eol {
		l.backup()
	}
}

func (l *lexer) ignore() bool {
	if l.pos == len(l.bb)-1 {
		return false
	}

	l.start = l.pos + 1
	return true
}

func (l *lexer) next() bool {
	if l.pos == len(l.bb)-1 {
		return false
	}

	l.pos++
	return true
}

func (l *lexer) curr() byte {
	if l.pos > len(l.bb)-1 {
		log.Fatal("invalid logic")
	}

	return l.bb[l.pos]
}

func (l *lexer) backup() {
	if l.pos == 0 {
		log.Fatalf("backup through 0")
	}

	l.pos--
}

func (l *lexer) peek() byte {
	if l.pos+1 > len(l.bb)-1 {
		log.Fatalf("peek passed: %d", l.pos)
	}

	return l.bb[l.pos+1]
}

func (l *lexer) emit(t itemType) {
	l.ch <- lexItem{t, string(l.bb[l.start : l.pos+1])}

	l.start = l.pos + 1
}

type lexStateFn func(*lexer) lexStateFn

func nextFn(l *lexer) (res lexStateFn) {
	if !l.next() {
		return
	}

	switch l.curr() {
	case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
		res = wordFn
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		res = numberFn
	case '<':
		res = ltFn
	case '>':
		res = gtFn
	case ':':
		res = becomesFn
	case '=', ';', '.', ',', '!', '?', '#', '+', '-', '*', '/', '(', ')':
		res = unaryOpFn
	case ' ':
		l.ignore()
		res = nextFn
	default:
		log.Fatalf("invalid char: %c", l.curr())
	}

	return
}

func wordFn(l *lexer) lexStateFn {
	l.acceptRun("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return idWordFn
}

func idWordFn(l *lexer) lexStateFn {
	word := string(l.bb[l.start : l.pos+1])
	if t, ok := keywords[word]; ok {
		l.emit(t)
	} else {
		l.emit(itemIdent)
	}

	return nextFn
}

func numberFn(l *lexer) lexStateFn {
	l.acceptRun("0123456789")
	l.emit(itemNumber)

	return nextFn
}

func gtFn(l *lexer) lexStateFn {
	if l.peek() == '=' {
		l.next()
		l.emit(itemGte)
	} else {
		l.emit(itemGt)
	}

	return nextFn
}

func ltFn(l *lexer) lexStateFn {
	if l.peek() == '=' {
		l.next()
		l.emit(itemLte)
	} else {
		l.emit(itemLt)
	}

	return nextFn
}

func becomesFn(l *lexer) lexStateFn {
	if !l.next() {
		log.Fatal("invalid becomes")
	}
	if l.curr() != '=' {
		log.Fatal("invalid becomes")
	}

	l.emit(itemBecomes)
	return nextFn
}

func unaryOpFn(l *lexer) lexStateFn {
	switch l.curr() {
	case '=':
		l.emit(itemEqual)
	case ';':
		l.emit(itemSemicolon)
	case '.':
		l.emit(itemDot)
	case ',':
		l.emit(itemComma)
	case '!':
		l.emit(itemNot)
	case '?':
		l.emit(itemQuestion)
	case '#':
		l.emit(itemHash)
	case '+':
		l.emit(itemPlus)
	case '-':
		l.emit(itemMinus)
	case '*':
		l.emit(itemMul)
	case '/':
		l.emit(itemDiv)
	case '(':
		l.emit(itemLparan)
	case ')':
		l.emit(itemRparan)
	}

	return nextFn
}

func in(r io.Reader) (res chan []byte) {
	var wg sync.WaitGroup
	res = make(chan []byte)

	wg.Add(1)
	go func() {
		defer wg.Done()

		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			//log.Printf("line: %s", scanner.Text())
			bb := bytes.ToLower(bytes.TrimSpace(scanner.Bytes()))
			if len(bb) > 0 {
				res <- bb
			}
		}
	}()

	go func() {
		wg.Wait()
		close(res)
	}()

	return
}

func lex(ch <-chan []byte) (res chan lexItem) {
	var wg sync.WaitGroup
	res = make(chan lexItem)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for bb := range ch {
			//log.Printf("lex line: %s", string(bb))
			l := &lexer{pos: -1, bb: bb, ch: res}
			for state := nextFn; state != nil; {
				state = state(l)
			}
		}

		res <- lexItem{itemEOF, ""}
	}()

	go func() {
		wg.Wait()
		close(res)
	}()

	return
}

type parser struct {
	sym lexItem
	ch  <-chan lexItem
}

func (p *parser) accept(li itemType) bool {
	if p.sym.t == li {
		p.next()
		return true
	}
	return false
}

func (p *parser) expect(li itemType) bool {
	if p.sym.t == li {
		p.next()
		return true
	}
	log.Fatal("expect: unexpected symbol")
	return false
}

func (p *parser) next() {
	p.sym = <-p.ch
}

func (p *parser) factor() {
	if p.accept(itemIdent) {
	} else if p.accept(itemNumber) {
	} else if p.accept(itemLparan) {
		p.expression()
		p.expect(itemRparan)
	} else {
		log.Fatal("factor: syntax error")
		p.next()
	}
}

func (p *parser) term() {
	p.factor()
	for p.sym.t == itemMul || p.sym.t == itemDiv {
		p.next()
		p.factor()
	}
}

func (p *parser) expression() {
	if p.sym.t == itemPlus || p.sym.t == itemMinus {
		p.next()
	}
	p.term()
	for p.sym.t == itemPlus || p.sym.t == itemMinus {
		p.next()
		p.term()
	}
}

func (p *parser) condition() {
	if p.accept(itemOdd) {
		p.expression()
	} else {
		p.expression()
		switch p.sym.t {
		case itemEqual, itemHash, itemLt, itemLte, itemGt, itemGte:
			p.next()
			p.expression()
		default:
			log.Fatal("condition: invalid operator")
		}
	}
}

func (p *parser) statement() {
	if p.accept(itemIdent) {
		p.expect(itemBecomes)
		p.expression()
	} else if p.accept(itemCall) {
		p.expect(itemIdent)
	} else if p.accept(itemQuestion) {
		p.expect(itemIdent)
	} else if p.accept(itemNot) {
		p.expression()
	} else if p.accept(itemBegin) {
		p.statement()
		for p.accept(itemSemicolon) {
			p.statement()
		}
		p.expect(itemEnd)
	} else if p.accept(itemIf) {
		p.condition()
		p.expect(itemThen)
		p.statement()
	} else if p.accept(itemWhile) {
		p.condition()
		p.expect(itemDo)
		p.statement()
	} else if p.accept(itemWrite) {
		p.expect(itemIdent)
	} else {
		log.Fatal("statement: syntax error")
		p.next()
	}
}

func (p *parser) block() {
	if p.accept(itemConst) {
		p.expect(itemIdent)
		p.expect(itemEqual)
		p.expect(itemNumber)
		for p.accept(itemComma) {
			p.expect(itemIdent)
			p.expect(itemEqual)
			p.expect(itemNumber)
		}
		p.expect(itemSemicolon)
	}
	if p.accept(itemVar) {
		p.expect(itemIdent)
		for p.accept(itemComma) {
			p.expect(itemIdent)
		}
		p.expect(itemSemicolon)
	}
	for p.accept(itemProcedure) {
		p.expect(itemIdent)
		p.expect(itemSemicolon)
		p.block()
		p.expect(itemSemicolon)
	}
	p.statement()
}

func parse(ch <-chan lexItem) bool {
	p := parser{ch: ch}

	p.next()
	p.block()
	p.expect(itemDot)
	p.expect(itemEOF)

	return true
}

func main() {
	//prog := `<debug>`
	//bb := bytes.NewBufferString(prog)
	//chLine := in(bb)

	// pipeline in -> lex -> parse -> eval -> print
	chLine := in(os.Stdin)
	chLex := lex(chLine)
	valid := parse(chLex)

	if valid {
		log.Print("Format is valid")
	} else {
		log.Print("Format is not valid")
	}
}
