package main

import (
	"fmt"
	"strings"
	"text/scanner"
	"unicode"
)

func main() {
	const src = `
// comment
if %a > 10 {
	hello = "world"
}`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "example"
	s.Mode ^= scanner.SkipComments // don't skip comments

	s.IsIdentRune = func(ch rune, i int) bool {
		return ch == '%' && i == 0 || unicode.IsLetter(ch) || unicode.IsDigit(ch) && i > 0
	}

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %v, %s\n", s.Position, tok, s.TokenText())
	}
}
