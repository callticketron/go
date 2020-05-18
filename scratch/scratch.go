package main

import (
	"fmt"
	"github.com/callticketron/requesties"
)

func main() {
	a := requesties.Get("http://www.reddit.com")
	fmt.Println(a)
}
