package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	routers := mux.NewRouter()
	routers.HandleFunc("/router", RouterFunction).Methods("GET")
	fmt.Println("server starting at port :9090")
	if err := http.ListenAndServe(":9090", routers); err == nil {
		fmt.Println("server starts at port :9090")
	} else {
		fmt.Println("Error := ", err.Error())
	}

}

func RouterFunction(h http.ResponseWriter, r *http.Request) {
	func() {
		fmt.Println("****************Printing the 1 function********************")
	}()
	func() {
		fmt.Println("****************Printing the 2 function********************")
	}()
	func() {
		fmt.Println("****************Printing the 3 function********************")
	}()
	func() {
		fmt.Println("****************Printing the 4 function********************")
	}()
	func() {
		fmt.Println("****************Printing the 5 function********************")
	}()

	go Function3()

	func() {
		fmt.Println("****************Printing the second function*******************")
	}()

	h.Header().Set("Content-Type", "application/json")
	json.NewEncoder(h).Encode("Api Executed Successfully")
}

func Function3() {
	i := 5
	for i == 5 {
		time.Sleep(1 * time.Second)
		fmt.Printf("Function 3 : %d\n", i)
	}
	// done <- true // Signal that the goroutine is done
}
