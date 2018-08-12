package main

import (
	"fmt"
	"os"
)

func main() {
	// var 선언은 문자열 타입의 두변수 s와 sep를 선언한다. 변수는 선언 중에 초기화할 수 있다
	// 명시적으로 초기화되지 않은 경우에는 묵시적으로 해당 타입의 제로 값으로 초기화되며, 숫자형의 제로 값은 0이고 문자형의 제로 값은 ""이다
	// 따라서 이 예에서 선언은 묵시적으로 s와 sep를 빈 문자열로 초기화한다.
	// := 기호는 짧은 변수 선언으로 하나 이상의 변수를 선언하고 초기화 값에 기초한 타입을 이 변수에 할당하는 문장이다
	// 후치 연산자만 사용가능하다 --i는 허용되지 않는다
	// for 루프는 Go의 유일한 루프 문이다
	// while 루프 표현 for 조건 { //todo }
	// 무한 루프 for { //todo }
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
