package main

import "fmt"

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{5}
	fmt.Println(CanJump(input3))

	input4 := []uint{7, 1, 1, 1, 1}
	fmt.Println(CanJump(input4))

	input5 := []uint{1, 1, 1, 1, 1}
	fmt.Println(CanJump(input5))

	input6 := []uint{}
	fmt.Println(CanJump(input6))
}

// CanJump()
func CanJump(s []uint) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}

	indx := uint(0)
	for i := len(s); i > 0; i-- {
		steps := s[indx]
		if indx+steps < uint(len(s)-1) {
			indx = indx + steps
		}
		if indx+steps == uint(len(s)-1) {
			return true
		}
		if indx+steps > uint(len(s)-1) {
			break
		}
	}
	return false
}
