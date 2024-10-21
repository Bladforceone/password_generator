package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Bladforceone/password_generator/generator"
)

func main() {
	length := flag.Int("l", 8, "Length of the password")
	letter := flag.Bool("let", true, "Include letters in the password")
	number := flag.Bool("numb", false, "Include numbers in the password")
	special := flag.Bool("spec", false, "Include special in the password")
	count := flag.Int("count", 1, "Number of passwords to generate")

	flag.Parse()

	ch := make(chan string)
	for i := 0; i < *count; i++ {
		go func() {
			password, err := generator.GeneratePassword(*length, *letter, *number, *special)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			ch <- password
		}()
	}

	save := []string{}
	for i := range *count {
		save = append(save, <-ch)
		fmt.Printf("Password: %s \n", save[i])
	}

	fmt.Println("Save the result in Documents/password.txt? y/n")

	var input string
	fmt.Scan(&input)

	if input == "Y" || input == "y" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		documentsPath := filepath.Join(homeDir, "Documents", "password.txt")
		file, err := os.OpenFile(documentsPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		defer file.Close()

		for i := range save {
			_, err := file.WriteString(save[i] + "\n")
			if err != nil {
				fmt.Println(err)
				os.Exit(4)
			}
		}

		fmt.Println("Saved")
	}
	os.Exit(0)
}
