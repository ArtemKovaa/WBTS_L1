	package main

	import "fmt"

	func main() {
		var num int64
		var bitPos int
		fmt.Print("Enter number and bit for reverse: ")
		fmt.Scanln(&num, &bitPos)

		fmt.Printf("Representation of %d in binary form: %b\n", num, num)

		mask := int64(1) << bitPos
		res := num ^ mask
		fmt.Printf("Representation of %d in binary form: %b", res, res)
	}
