package main

import "fmt"

// TODO don't use global vars!
var (
	identifier_str string
	num_val        float64
)

// all possible lexical tokens
const (
	tok_eof = iota

	// commands
	tok_def
	tok_extern

	// primary
	tok_identifier
	tok_number
)

func get_token() rune {
	last_char = ' '

  
}

// debug only
func main() {
	fmt.Printf("%v\n%v\n", tok_def, tok_number)
}
