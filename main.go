package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		case "update":
			input := strings.Split(line, " ")
			if len(input) < 2 || input[1] == "" {
				fmt.Println("[Error] Missing position argument")
				break
			}
			if len(input) < 3 || input[2] == "" {
				fmt.Println("[Error] Missing note argument")
				break
			}
			numberAndData := strings.SplitN(commandAndData[1], " ", 2)
			numberStr := numberAndData[0]
			number, err := strconv.ParseInt(numberStr, 10, 64)
			if err != nil {
				fmt.Printf("[Error] Invalid position: %s\n", numberStr)
				break
			}
			if number < 1 || number > int64(n) {
				fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", number, n)
				break
			}
			if notes[number-1] == "" {
				fmt.Println("[Error] There is nothing to update")
				break
			}
			notes[number-1] = numberAndData[1]
			fmt.Printf("[OK] The note at position %d was successfully updated\n", number)
		default:
			fmt.Println("[Error] Unknown command")
		}
	}
	fmt.Println("[Info] Bye!")
}
