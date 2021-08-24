package main

import "fmt"

func main() {

	//x := 3
	//b := string(x)
	// s := []string{"5", "6"}
	// s = append(s, "1", "2", "3", "fuck u, beach")
	// if s != nil {
	// 	s = s[2:len(s)]
	// 	fmt.Printf("%v\n", s)
	// 	fmt.Printf("Length is %v \n", len(s))
	// }
	f := []int{1, 2, 3}
	s := make([]int, len(f), cap(f))
	copy(s, f)
	s[0] = 3
	fmt.Printf("%v\n%v\n", s, f)
}
