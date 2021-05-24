package main

import (
	"fmt"

	"github.com/Heongilee/learngo/something"
)

// fmt package
//
// 포맷팅을 위한 패키지.

// Go에서는 너가 어떤 패키지를 사용하는지 작성해줘야 해
// 다른 오픈소스 프로젝트에 main.go가 없다는 것은 오픈 소스 프로젝트 개발자가
// 그것을 컴파일하기 원치 않기 때문이라고 볼 수 있지.

// 왜 하필 함수 이름 첫 글자가 대문자(Uppercase)일까?
//
// Go에서는 function을 export 시킬때 대문자로 시작해주면 되기 때문이다.
func main() {
	fmt.Println("Hello world!")

	something.SayHello()
}
