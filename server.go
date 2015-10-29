package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server started listening at port 3000")
	panic(http.ListenAndServe(":3000", http.FileServer(http.Dir("./"))))
}
