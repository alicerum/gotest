package main

import "fmt"

type MyInterface interface {
	method1(val1 int) int
	method2(val1, val2 int) int
}

type defaultImplementation struct {
	Addition int
}

func (d *defaultImplementation) method1(val1 int) int {
	return val1 + d.Addition
}

func (d *defaultImplementation) method2(val1, val2 int) int {
	return val1 + val2 + d.Addition
}

// Now this is fun
type VeryMuchMyStruct struct {
	MyInterface
}

func (m *VeryMuchMyStruct) method1(val1 int) int {
	return val1 + 1
}

func main() {
	var mi MyInterface = &VeryMuchMyStruct{
		MyInterface: &defaultImplementation{Addition: 3},
	}

	fmt.Printf("method1 (should add 1 because overrided, so result is 6): %d\n", mi.method1(5))
	fmt.Printf("method2 (taken from default implementation, so result is 6):%d\n", mi.method2(1, 2))
}
