package main

import (
	"fmt"
	//"strings"
	"time"
)

func SaySth(word string) {
	fmt.Println(word)
}

func main() {
	SaySth("hello 1!")
	time.AfterFunc(time.Duration(1)*time.Second, func() { SaySth("hello 2!") })
	time.AfterFunc(time.Duration(2)*time.Second, func() { SaySth("hello 3!") })
	time.AfterFunc(time.Duration(3)*time.Second, func() { SaySth("hello 4!") })
	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println("hello 5!")
}
