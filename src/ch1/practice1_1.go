package main

import (
	"fmt"
	"os"
	"strings"
)

/*
	echo 프로그램을 수정해 호출한 명령인 os.Args[0]도 같이 출력하라.
*/

func main() {
	/*
		s, sep := "", ""
		for _, args := range os.Args[0:] {
			s += sep + args
			sep = " "
		}
		fmt.Println(s)
	*/
	fmt.Println(strings.Join(os.Args[0:], " "))
}
