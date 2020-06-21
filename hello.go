package main

import (
	"fmt"
	"ferrugem.com.br/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, world."))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
