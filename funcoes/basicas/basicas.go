package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("%s %s \n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("%s, %s", p1, p2)
}

func f5() (string, string) {
	return "Return1", "Return2"
}

func main() {
	f1()
	f2("Param1", "Param2")

	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)

	_, r51 := f5()
	fmt.Println(r51)

	r50, r51 := f5()
	fmt.Println(r50, r51)

}
