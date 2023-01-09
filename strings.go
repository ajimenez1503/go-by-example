package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Contains:", strings.Contains("Test", "es"))
	fmt.Println("Count:", strings.Count("test", "t"))
	fmt.Println("HasPrefix: ", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", strings.HasSuffix("test", "st"))
	fmt.Println("Index: ", strings.Index("test", "e"))
	fmt.Println("Join: ", strings.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat: ", strings.Repeat("a", 5))
	fmt.Println("Replace: ", strings.Replace("foo", "o", "O", -1))
	fmt.Println("Replace: ", strings.Replace("foo", "o", "O", 1))
	fmt.Println("split: ", strings.Split("a-b-c", "-"))
	fmt.Println("ToLower: ", strings.ToLower("TEST"))
	fmt.Println("ToUpper: ", strings.ToUpper("test"))
}
