package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/RouterFunction", RouterFunction).Methods("GET")
	fmt.Println("server starting at port :9090")
	if err := http.ListenAndServe(":9090", nil); err == nil {
		fmt.Println("server starts at port :9090")
	} else {
		fmt.Println("Error := ", err.Error())
	}

}

func RouterFunction(h http.ResponseWriter, r *http.Request) {

	// done := make(chan bool) // Create a channel to signal when goroutine is done

	func() {
		fmt.Println("****************Printing the first function********************")
	}()

	Function3()

	func() {
		fmt.Println("****************Printing the second function*******************")
	}()

	h.Write([]byte("welcome to the chi"))

	// <-done

}

func Function3() {
	// for i := 1000; i < 1100; i++ {
	// 	fmt.Printf("Function 3 : %d\n", i)
	// }
	i := 5
	if i == 5 {
		fmt.Printf("Function 3 : %d\n", i)
	}
	// done <- true // Signal that the goroutine is done
}
