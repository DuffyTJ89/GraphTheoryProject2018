package main

import (
	"fmt"
)

func intopost(infix string) string { //function to convert inflix regExp to postfix regExp

	specials := map[rune]int{'*': 10, '.': 9, '|': 8} //create a map for special characters which will be used in this program

	pofix := []rune{}
	s := []rune{}

	return string(pofix)
}

func main() {

	//answer ab.c*.
	fmt.Println("Infix :  ", "a.b.c*") //A followed by B, followed by zero or more C
	fmt.Println("Postfix : ", intopost("a.b.c*"))

	//answer abd|.*
	fmt.Println("Infix :  ", "(a.(b|d))*") //Zero or more of A followed by B or D
	fmt.Println("Postfix : ", intopost("(a.(b|d))*"))

	//answer abd|.c*.
	fmt.Println("Infix :  ", "a.(b|d).c*")
	fmt.Println("Postfix : ", intopost("(a.(b|d).c*"))

	//answer abb.+.c.
	fmt.Println("Infix :  ", "a.(b.b)+.c")
	fmt.Println("Postfix : ", intopost("a.(b.b)+.c"))
}
