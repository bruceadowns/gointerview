package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
)

func generateCombos(sl []uint8) (res []map[uint8]int) {
	if len(sl) == 0 {
		return nil
	}

	children := generateCombos(sl[1:])
	if children == nil {
		for i := 0; i < 10; i++ {
			m := make(map[uint8]int)
			m[sl[0]] = i
			res = append(res, m)
		}
	} else {
		for i := 0; i < 10; i++ {
			for _, v := range children {
				m := make(map[uint8]int)
				for kk, vv := range v {
					m[kk] = vv
				}
				m[sl[0]] = i
				res = append(res, m)
			}
		}
	}

	return
}

type equation struct {
	diff, lhs, rhs string
}

func (e *equation) uniq() (res []uint8) {
	m := make(map[uint8]struct{})
	m[e.diff[0]] = struct{}{}
	m[e.diff[1]] = struct{}{}
	m[e.lhs[0]] = struct{}{}
	m[e.lhs[1]] = struct{}{}
	m[e.rhs[0]] = struct{}{}
	m[e.rhs[1]] = struct{}{}

	for k := range m {
		res = append(res, k)
	}

	return
}

func (e *equation) true(m map[uint8]int) bool {
	d := m[e.diff[0]]*10 + m[e.diff[1]]
	l := m[e.lhs[0]]*10 + m[e.lhs[1]]
	r := m[e.rhs[0]]*10 + m[e.rhs[1]]
	return d == l-r
}

func (e *equation) resolve(m map[uint8]int) string {
	d := m[e.diff[0]]*10 + m[e.diff[1]]
	l := m[e.lhs[0]]*10 + m[e.lhs[1]]
	r := m[e.rhs[0]]*10 + m[e.rhs[1]]
	return fmt.Sprintf("%d = %d - %d", d, l, r)
}

func in(r io.Reader) (res []equation) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var d, l, r string
		num, err := fmt.Sscanf(scanner.Text(),
			"%s = %s - %s", &d, &l, &r)
		if err != nil {
			log.Fatal(err)
		}
		if num != 3 {
			log.Fatalf("error scanning actual %d expect 3", num)
		}

		res = append(res, equation{d, l, r})
	}

	return
}

func main() {
	eqs := in(bytes.NewBufferString(`AB = CA - DE
FG = BE - HG
DB = AF - IA
JC = KC - JE
JL = LJ - DK`))
	//HE = CG - BI`))
	//eqs := in(os.Stdin)

	var accUniq []uint8
	var prevCombos []map[uint8]int
	for _, eq := range eqs {
		uniq := eq.uniq()
		log.Printf("unique letters: %v for %s", uniq, eq)

		var diffUniq []uint8
	outerUniq:
		for _, v := range uniq {
			for _, vv := range accUniq {
				if v == vv {
					continue outerUniq
				}
			}

			diffUniq = append(diffUniq, v)
		}
		log.Printf("diff unique letters: %v for %s", diffUniq, eq)

		var workingCombos []map[uint8]int
		genCombos := generateCombos(diffUniq)
		if len(genCombos) == 0 {
			workingCombos = prevCombos
		} else if len(prevCombos) == 0 {
			workingCombos = genCombos
		} else {
			for _, v := range prevCombos {
				for _, vv := range genCombos {
					m := make(map[uint8]int)
					for kkk, vvv := range v {
						m[kkk] = vvv
					}
					for kkk, vvv := range vv {
						m[kkk] = vvv
					}
					workingCombos = append(workingCombos, m)
				}
			}
		}
		log.Printf("total combos: %d", len(workingCombos))

		var potentialCombos []map[uint8]int
		for _, combo := range workingCombos {
			if eq.true(combo) {
				potentialCombos = append(potentialCombos, combo)
			}
		}
		log.Printf("potentials: %d", len(potentialCombos))

		prevCombos = potentialCombos
		accUniq = append(accUniq, diffUniq...)

		log.Printf("accumulated uniques: %v", accUniq)
		log.Printf("accumulated combinations: %v", prevCombos)
	}

	for _, v := range prevCombos {
		m := make(map[int]struct{})
		for _, vv := range v {
			m[vv] = struct{}{}
		}

		if len(m) == 8 {
			for _, vv := range eqs {
				log.Print(vv.resolve(v))
			}
			log.Print()
		}
	}
}
