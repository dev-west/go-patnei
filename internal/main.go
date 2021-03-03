package main

import (
	"errors"
	"fmt"
	"github.com/dev-west/go-patnei"
)

func main() {
	var p *ExObj
	var e error

	expected := "ex"
	secret := "sec"

	p = NewObj() *ExObj {
	fmt.Printf("p.s: \"%s\"\n", p.s)
	p, e = p.Append(secret)
	fmt.Printf("p.s: \"%s\"\n", p.s)
	p, e = p.Add(secret)
	fmt.Printf("p.s: \"%s\"\n", p.s)

	p = NewObj() *ExObj {
	fmt.Printf("p.s: \"%s\"\n", p.s)
	p, e = p.Append(expected)
	fmt.Printf("p.s: \"%s\"\n", p.s)
	p, e = p.Add(expected)
	fmt.Printf("p.s: \"%s\"\n", p.s)
}
