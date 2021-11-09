package main

import "fmt"

func swap(x string, y string, z int) (int, string, string) {
	return z, y, x
}

func main() {
	a, b, c := swap("hello", "world", 23)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// string,intなど型の違った返り値も複数返すことができる