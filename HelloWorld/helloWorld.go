package main

import (
	"fmt"
	"rsc.io/quote"
	"example.com/greetings"
)

func main() {
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
	fmt.Println(greetings.CustomHello("Matt"))
}
