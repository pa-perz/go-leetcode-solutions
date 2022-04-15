package main

import (
	"errors"
	"log"
)

type Parenthesis struct {
	open  rune
	close rune
}

const p1_open = '{'
const p1_close = '}'
const p2_open = '('
const p2_close = ')'
const p3_open = '['
const p3_close = ']'

var parenthesisDict = map[int]Parenthesis{
	0: {open: p1_open, close: p1_close},
	1: {open: p2_open, close: p2_close},
	2: {open: p3_open, close: p3_close},
}

var openPar = [3]rune{p1_open, p2_open, p3_open}
var closePar = [3]rune{p1_close, p2_close, p3_close}

func isValid(s string) bool {
	var orphanPar []rune

	for _, ch := range s {
		if contains(closePar[:], ch) {
			p, err := getParenthesisFromClose(ch)
			if err != nil {
				log.Fatal(err)
			}
			if len(orphanPar) == 0 {
				return false
			}
			if p.open != orphanPar[len(orphanPar)-1:][0] {
				return false
			} else {
				if len(orphanPar) > 0 {
					orphanPar = orphanPar[:len(orphanPar)-1]
				}
			}

		} else if contains(openPar[:], ch) {
			orphanPar = append(orphanPar, ch)
		}
	}

	return len(orphanPar) <= 0
}

func contains(arr []rune, s rune) bool {
	for _, p := range arr {
		if s == p {
			return true
		}
	}
	return false
}

func getParenthesisFromClose(close rune) (Parenthesis, error) {
	switch close {
	case p1_close:
		return parenthesisDict[0], nil
	case p2_close:
		return parenthesisDict[1], nil
	case p3_close:
		return parenthesisDict[2], nil
	}
	return Parenthesis{}, errors.New("Parenthesis not found")

}
