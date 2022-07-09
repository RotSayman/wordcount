package main

import(
	"os"
	"fmt"
	"errors"
)
func main(){
	text, err := readInput()
	if err != nil{
		fmt.Println("->", err)
	}else{
		fmt.Println(countWord(text))
	}
}

func readInput() (string,error){
	args := os.Args
	if len(args) != 2 {
		return "", errors.New("wordcount 'your string'")
	}
	if args[1] == ""{
		return "", nil
	}
	return args[1], nil
}

func countWord(text string) uint{
	var count uint
	if text == ""{
		return 0
	}
	for _, ch := range text{
		if ch == rune(' '){
			count++
		}
	}
	// +1 last word in text
	return count + 1
}
