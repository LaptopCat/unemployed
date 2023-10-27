package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ZERO  = "currently"
	ONE   = "unemployed"
	SPACE = "i am"
)

func binaryConverter(what string) string {
	return strings.ReplaceAll(strings.ReplaceAll(what, "0", ZERO+" "), "1", ONE+" ")
}

func encoder(what string) (result string) {
	for _, v := range []byte(what) {
		result += binaryConverter(strconv.FormatInt(int64(v), 2)) + SPACE + " "
	}
	return strings.Trim(strings.TrimSuffix(result, SPACE+" "), " ")
}

func decoder(what string) (result string) {
	if len(what) == 0 {
		return ""
	}
	chunks := []byte{}
	for _, v := range strings.Split(what, SPACE) {
		result, err := strconv.ParseInt(
			strings.ReplaceAll(
				strings.ReplaceAll(strings.ReplaceAll(v, ZERO, "0"), ONE, "1"), " ", ""), 2, 64)
		if err != nil {
			panic(err)
		}
		chunks = append(chunks, byte(result))
	}
	return string(chunks)
}

func main() {
	encoded := encoder(strings.Join(os.Args[1:], " "))
	fmt.Println("encoded:", encoded)
	fmt.Println()
	fmt.Println("decoded:", decoder(encoded))
}
