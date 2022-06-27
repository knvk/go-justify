package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	numWords   int
	maxLineLen int = 80
	word           = []rune{}
	line           = []rune{}
)

func readWord(r *bufio.Reader) (int, error) {
	var l int
	for {
		ch, err := readChar(r)
		if (err == nil) && (ch != ' ') {
			word = append(word, ch)
			l++
			//fmt.Println(word)
		} else {
			return l, err
		}
	}
}

func addWord() {
	if numWords > 0 {
		line = append(line, ' ')
	}
	line = append(line, word...)
	numWords++
}

func readChar(r *bufio.Reader) (rune, error) {
	ch, _, err := r.ReadRune()
	if err != nil {
		if err == io.EOF {
			return 0, io.EOF
		}
		return -1, err
	}
	if (ch == '\t') || (ch == '\n') {
		return ' ', nil
	}
	return ch, nil
}

func writeLine() {
	extraSpaces := maxLineLen - len([]rune(line))
	for _, ch := range line {
		if ch != ' ' {
			fmt.Printf(string(ch))
		} else {
			insertSpaces := extraSpaces / (numWords - 1)
			if extraSpaces > 0 && insertSpaces == 0 {
				insertSpaces++
			}
			for j := 0; j <= insertSpaces; j++ {
				fmt.Printf(string(' '))
			}
			extraSpaces -= insertSpaces
			numWords--
		}
	}
	fmt.Printf("\n")
}

func main() {
	r := bufio.NewReader(os.Stdin)
	if len(os.Args) > 1 {
		var err error
		maxLineLen, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("usage: ./justify <line_len>\n")
			os.Exit(1)
		}
	}
	for {
		// reset slice
		word = []rune{}
		_, err := readWord(r)
		if err == io.EOF {
			fmt.Println(string(line), string(' '))
			break
		}
		if len([]rune(line))+len([]rune(word))+1 > maxLineLen {
			writeLine()
			// reset line
			line = []rune{}
			numWords = 0
		}
		addWord()
	}
}
