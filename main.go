package main

import (
	"fmt"
	"rsc.io/quote"
	"github.com/cherishAchange/go_demo2"
	"go_demo1/branch"
)

func main()  {
	message := go_demo2.Hello("Gladys")
	message1 := branch.Branch2("niubi")
	fmt.Println(message1)
	fmt.Println(message)
	fmt.Println("Hello World!")
	fmt.Println(quote.Go())
}