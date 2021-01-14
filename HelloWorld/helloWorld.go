package main

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())

	message := "Hello, world! (From Matt)."
	fmt.Println(message)
}
