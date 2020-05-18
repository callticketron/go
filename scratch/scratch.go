package main

import (
	"fmt"
	"github.com/callticketron/requesties"
)

func main() {
	a := requesties.Get("http://www.google.com")
	fmt.Println(a)
}
