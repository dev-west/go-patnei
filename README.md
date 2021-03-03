# go-patnei
Bug reminder docs: pointer access to non-exported identifier (patnei) not caught by compiler

This repo should NOT be used for its package. It is intended as an example to new gophers regarding isolating package identifiers intended to contain non-exported variables.

This bug occurs when a pointer to a non-exported identifier is returned by an exported method. This violates the specification for exported identifiers: https://golang.org/ref/spec#Exported_identifiers

The go implementation (as of Feb 10, 2021) appears to prevent direct access to non-exported identifiers, but this can by bypassed by the developer by returning a pointer via an exported method. In my opinion, this should be flagged by the compiler.

Hope it helps. -- Regards.
