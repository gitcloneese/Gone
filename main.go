package main

import (
	"fmt"
	"github.com/gitcloneese/gomodone/v3"
)

func main() {
	g, err := gomodone.SayHi("Roberto", "pt", "long time no see")
	if err != nil {
		panic(err)
	}
	fmt.Println(g)
}
