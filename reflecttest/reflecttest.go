package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(&x)
	fmt.Println("interface:", v.Interface())
	fmt.Println("settability of v:", v.Elem().CanSet())
	v.Elem().SetFloat(6.8)
	fmt.Println("x:", x)
	fmt.Println("v:", v.Elem().Interface())

	type T struct {
		A int
		B string
	}
	t := T{32, "itegel"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

}
