package main

import (
	"fmt"
	"log"
)

func main() {

	//web server
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Go!"))
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }

	fmt.Println("start")
	paniker()
	fmt.Println("end")
}

func paniker() {
	fmt.Println("about tp panic")

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			// panic(err)
		}
	}()

	panic("Something bad happen!")
}
