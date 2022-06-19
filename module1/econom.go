package main

import (
	"fmt"

	"strings"
)

func main() {

	//examples := []string{"x", "($xy)", "($(@ab)c)", "(#i($jk))", "(#($ab)($ab))", "(@(#ab)($ab))", "(#($a($b($cd)))(@($b($cd))($a($b($cd)))))", "(#($(#xy)($(#ab)(#ab)))(@z($(#ab)(#ab))))"}
	var input string
	fmt.Scanln(&input)

	s := ([]byte)(input)

	p := strings.Index(input, ")")

	k := 0

	for p != -1 {

		p -= 4

		k++

		s = ([]byte)(strings.Replace(string(s), string(s[p:p+5]), string(k), -1))

		p = strings.Index(string(s), ")")

	}

	fmt.Printf("%d\n", k)

}
