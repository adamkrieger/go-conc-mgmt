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
			// for a, b := range r {
			// 	fmt.Println("each: ", a, " ", b)
			// }
			if _, is := r.(*specificError); is {
				fmt.Println("-is of type error")
			} else {
				fmt.Println("-is some other type")
			}
		}
	}()

	fmt.Println("processing")
	g(0, "thing")
	fmt.Println("returned normally")
}

func g(i int, some string) {
	if i > 3 {
		fmt.Println("panicking!")
		//panic(fmt.Sprintf("%v %s", i, some))
		//arr := []string{some}

		panic(&specificError{code: 404, msg: "not found"})

		//fmt.Println(arr[100])
	}

	defer fmt.Println("defer in g", i, some)
	fmt.Println("printing in g", i, some)
	g(i+1, some)
}
