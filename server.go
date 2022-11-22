package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(":3000", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	hostname, error := os.Hostname()

	if error != nil {
		fmt.Println(error)
		w.Write([]byte(string(error.Error())))
	}

	w.Write([]byte(fmt.Sprintf("<h1>K8s - GitActions Flows - Hostname: %s</h1>", hostname)))
}
