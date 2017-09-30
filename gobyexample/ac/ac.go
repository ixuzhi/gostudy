package main

import (
	"fmt"

	"github.com/cloudflare/ahocorasick"
)

func main() {
	fmt.Println("AC")
	m := ahocorasick.NewStringMatcher([]string{"Superman", "Superma", "Superm", "Super"})
	hits := m.Match([]byte("The Man Of Steel: Superman"))
	fmt.Println(hits)
}
