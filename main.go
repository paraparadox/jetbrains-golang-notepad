package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	const n = 5
	var i int
	var notes [n]string
	for {
		fmt.Print("Enter command and data: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		commandAndData := strings.SplitN(line, " ", 2)
		command := commandAndData[0]
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		if command == "exit" {
			break
		}
		switch command {
		case "create":
			if i == n {
				fmt.Println("[Error] Notepad is full")
			} else {
				data := commandAndData[1]
				notes[i] = data
				i++
				fmt.Println("[OK] The note was successfully created")
			}
		case "list":
			for j := 0; j < n; j++ {
				if notes[j] != "" {
					fmt.Printf("[INFO] %d: %s\n", j+1, notes[j])
				}
			}
		case "clear":
			for j := 0; j < n; j++ {
				notes[j] = ""
			}
			i = 0
			fmt.Println("[OK] All notes were successfully deleted")
		}
	}
	fmt.Println("[Info] Bye!")
}
