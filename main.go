package main

import (
	"fmt"
	"strings"
)

func main() {
    fmt.Print("Hello, World!")
}

func cleanText(text string) []string {

    var cleanWords []string
    if strings.TrimSpace(text) == "" {
        return cleanWords
    }

    cleanWords = strings.Split(text, " ")
    for i := range cleanWords {
        cleanWords[i] = strings.ToLower(cleanWords[i])
    }

    return cleanWords

}
