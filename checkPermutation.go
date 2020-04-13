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

		fmt.Printf("input string: ")
		str2, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error accured while reading user input")
		}

		str = strings.Replace(str, "\n", "", -1)
		str2 = strings.Replace(str2, "\n", "", -1)

		fmt.Println(checkInclusion(str, str2))
	}
}

func checkInclusion(s1 string, s2 string) bool {
	checker := 0

	for _, char := range s1 {
		checker ^= int(char - '0')
	}

	for _, char := range s2 {
		checker ^= int(char - '0')
	}

	return checker == 0
}
