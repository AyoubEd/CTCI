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
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error reading user input")
		}

		str = strings.Replace(str, "\n", "", -1)

		fmt.Println(URLify(str))
	}
}

func URLify(str string) string {
	strByteArr := []byte(str)
	res := make([]byte, len(str))

	i := len(str) - 1
	for strByteArr[i] == ' ' {
		i--
	}
	margin := len(str) - i - 1

	for ; i >= 0; i-- {
		if strByteArr[i] == ' ' {
			res[i+margin] = '0'
			res[i+margin-1] = '2'
			res[i+margin-2] = '%'
			margin -= 3
		} else {
			res[i+margin] = strByteArr[i]
		}
	}

	return string(res)
}
