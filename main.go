package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	for {
		fmt.Print("Enter command and data: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		scanner.Scan()
		command := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		if command == "exit" {
			break
		}
		fmt.Println(command)
	}
	fmt.Println("[Info] Bye!")
}
