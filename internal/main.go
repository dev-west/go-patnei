package main

import (
	"fmt"
	. "github.com/dev-west/go-patnei"
)

func main() {
	var p *ExObj
	var e error

	expected := "ex"
	secret := "sec"

	// generate new ExObj
	p = NewObj()
	fmt.Printf("NewObj() =>\t\t\t\tp.S: \"%s\"\n", p.S)
	// set p to internal eEmpty
	if p, e = p.Append(secret); e != nil {
		fmt.Printf("Append(secret) expects error =>\t\tp.S: \"%s\"; e: \"%s\"\n", p.S, e)
	} else {
		fmt.Printf("Append(secret) unexpected no error =>\t\tp.S: \"%s\"\n", p.S)
	}
	// write secret to internal eEmpty
	if p, e = p.Add(secret); e != nil {
		fmt.Printf("Add(secret) unexpected error =>\t\tp.S: \"%s\"; e: \"%s\"\n", p.S, e)
	} else {
		fmt.Printf("Add(secret) expects \"%s\" =>\t\tp.S: \"%s\"\n", secret, p.S)
	}

	// generate new ExObj
	p = NewObj()
	fmt.Printf("NewObj() =>\t\t\t\tp.S: \"%s\"\n", p.S)
	// set p to internal eEmpty
	if p, e = p.Append(expected); e != nil {
		fmt.Printf("Append(expected) expects error =>\tp.S: \"%s\"; e: \"%s\"\n", p.S, e)
	} else {
		fmt.Printf("Append(expected) unexpected no error =>\tp.S: \"%s\"\n", p.S)
	}
	// read value of eEmpty
	fmt.Printf("Object points to internal eEmpty => \tp.S: \"%s\"\n\n", p.S)
}
