package main

import (
	"fmt"
	"io"
	"net/http"
)

const keyServerAddr = "serverAddr"

func getStl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DBG: got /stl request")

	myName := r.PostFormValue("file")
	fmt.Printf("DBG: %s\n", myName)
	if myName == "" {
		w.Header().Set("x-missing-field", "file")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("DBG: %+v\n", r)
	io.WriteString(w, fmt.Sprintf("Hello, %+v!\n", r))
}

func getStlOld(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintln(w, "Failed to parse form data")
		panic(err)
	}
	fmt.Printf("DEBUG1: %+v\n", r.Form)
	x := r.Form.Get("file")
	fmt.Printf("DEBUG2: %+v\n", r)
	fmt.Printf("DEBUG3: %s\n", x)
	fmt.Fprintln(w, "Form data parsed.")
}

func main() {
	http.HandleFunc("/stl", getStl)

	err := http.ListenAndServe("0.0.0.0:9999", nil)
	if err != nil {
		panic(err)
	}
}
