package main

import(
	"os"
	"fmt"
)
func main(){
	text := readInput()
	fmt.Println(countWord(text))
}

func readInput() string{
	args := os.Args
	if len(args) != 2 {
		fmt.Println("wordcount 'your string'")
	}
	return args[1]
}

func countWord(text string) uint{
	var count uint
	for _, ch := range text{
		if ch == rune(' '){
			count++
		}
	}
	// +1 last word in text
	return count + 1
}
