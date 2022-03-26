package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var n, i int
	fmt.Print("Enter the maximum number of notes: ")
	fmt.Scan(&n)
	notes := make([]string, n)
	for {
		fmt.Print("Enter command and data: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		line := scanner.Text()
		commandAndData := strings.SplitN(line, " ", 2)
		command := commandAndData[0]
		if command == "exit" {
			break
		}
		switch command {
		case "create":
			if i == n {
				fmt.Println("[Error] Notepad is full")
				break
			}
			input := strings.Split(line, " ")
			if len(input) == 1 || input[1] == "" {
				fmt.Println("[Error] Missing note argument")
				break
			}
			data := commandAndData[1]
			notes[i] = data
			i++
			fmt.Println("[OK] The note was successfully created")
		case "list":
			if i == 0 {
				fmt.Println("[Info] Notepad is empty")
			} else {
				for j := 0; j < i; j++ {
					fmt.Printf("[INFO] %d: %s\n", j+1, notes[j])
				}
			}
		case "clear":
			for j := 0; j < n; j++ {
				notes[j] = ""
			}
			i = 0
			fmt.Println("[OK] All notes were successfully deleted")
		default:
			fmt.Println("[Error] Unknown command")
		}
	}
	fmt.Println("[Info] Bye!")
}
