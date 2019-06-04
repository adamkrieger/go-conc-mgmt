package main

import "fmt"

func main() {

	process()
	fmt.Println("returned normally from processing.")
}

type specificError struct {
	code int
	msg  string
}

func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered into the defer in process()", r)

			if _, is := r.(*specificError); is {
				fmt.Println("-is of type error")
			} else {
				fmt.Println("-is some other type")
			}
		}
	}()

	fmt.Println("processing")
	usefulFunction(0, "thing")
	fmt.Println("returned normally")
}

func usefulFunction(i int, some string) {
	if i > 3 {
		fmt.Println("panicking!")

		panic(&specificError{code: 404, msg: "not found"})
	}

	defer fmt.Println("defer in g", i, some)
	fmt.Println("printing in g", i, some)
	usefulFunction(i+1, some)
}
