package main

import (
	"fmt"
	"rin/libs/rope"
)

func main() {
	r := rope.NewRope("Hello world")
	i := rope.NewRope("I am john")
	k := r.Concatenate(rope.NewRope(" "))
	f := k.Concatenate(i)
	fmt.Printf(f.String())
}
