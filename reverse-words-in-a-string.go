package main

import (
	"fmt"
	"strings"
)

func main_reverse_words_in_string() {
	s := " a good   example  "
	// s = strings.Trim(s, " ")
	s = strings.Join(strings.Fields(s), " ")
	fmt.Println(s)

	words := strings.Split(s, " ")

	for i := 0; i < len(words)/2; i++ {
		fmt.Println(i, len(words)-1)
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	fmt.Println(strings.Join(words, " "), len(words))

}
