package main

import (
	"fmt"
	"bufio"
	"os"
)

type Cypher struct {
	dict map[rune]rune
	backdict map[rune]rune
	input string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	m := map[rune]rune{}
	b := map[rune]rune{}
	c := Cypher{m,b,  input}
	fmt.Print(c.input)
	cont := true
	for cont {
		fmt.Print("Enter character to replace or 'quit':")
		k, _ := reader.ReadString('\n')
		if k == "quit\n" {
			cont = false
			break
		}
		fmt.Print("Enter replacement or 'quit':")
		v, _ := reader.ReadString('\n')
		if v == "quit\n" {
			cont = false
			break
		}
		krunes := []rune(k)
		vrunes := []rune(v)
		if vrunes[0] == krunes[0] {
			fmt.Printf("Removing mapping for %s\n", string(vrunes[0]))
			oldVal, _ := c.dict[krunes[0]]
			delete(c.backdict, oldVal)
			delete(c.dict, krunes[0])
			continue
		}
		if key, ok := c.backdict[vrunes[0]]; ok {
			fmt.Printf("%s is already mapped to %s\n", string(vrunes[0]), string(key))
			continue

		}
		c.dict[krunes[0]] = vrunes[0]
		c.backdict[vrunes[0]] = krunes[0]
		fmt.Print("Your result is : \n")
		for _, r := range c.input {
			if sub, ok := c.dict[r]; ok {
				fmt.Printf("\x1b[31;1m%s\x1b[0m", string(sub))
			} else {
				fmt.Print(string(r))
			}
		}
		fmt.Printf("Your original is : \n%s \n", c.input)
	}
}