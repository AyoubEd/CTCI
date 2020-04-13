package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("input string: ")
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error accured while reading user input")
		}

		fmt.Println(checkIsUnique(strings.Replace(str, "\n", "", -1)))
	}
}

func checkIsUnique(s1 string) bool {
	checker := 0

	for _, char := range s1 {
		if checker&(1<<int(char-'a')) > 0 {
			return false
		}
		checker |= 1 << int(char-'a')
	}

	return true
}
