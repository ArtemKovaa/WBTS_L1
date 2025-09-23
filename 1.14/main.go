package main

import "log"

func main() {
	describeType(5)
	describeType("sap")
	describeType(make(chan int))
	describeType(true)
	describeType(9.0)
	describeType(make(chan bool))

	ch := make(chan<- string)
	describeType(ch)
}

func describeType(v interface{}) {
	switch v.(type) {
	case int:
		log.Println("It's integer")
	case string:
		log.Println("It's string")
	case bool:
		log.Println("It's bool")
	case chan int:
		log.Println("It's int channel")
	case chan string:
		log.Println("It's string channel")
	case chan bool:
		log.Println("It's bool channel")
	default:
		log.Println("It's unknown type")
	}
}
