package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Bladforceone/password_generator/generator"
)

func main() {
	length := flag.Int("l", 8, "Length of the password")
	letter := flag.Bool("let", true, "Include letters in the password")
	number := flag.Bool("numb", false, "Include numbers in the password")
	special := flag.Bool("spec", false, "Include special in the password")
	count := flag.Int("count", 1, "Number of passwords to generate")

	flag.Parse()

	for i := 0; i < *count; i++ {
		password, err := generator.GeneratePassword(*length, *letter, *number, *special)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Password:", password)
	}
}
