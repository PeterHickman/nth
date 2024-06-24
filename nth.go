package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var nth = flag.Int("n", 0, "print out every nth line")
	flag.Parse()

	if *nth <= 1 {
		fmt.Println("A value > 1 would be desirable")
		os.Exit(2)
	}

	line := *nth

	scanner := bufio.NewScanner(os.Stdin)
	c := 0

	for scanner.Scan() {
		text := scanner.Text()

		c++
		if c == line {
			fmt.Println(text)
			line += *nth
		}
	}
}
