package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func dbg(s string) {
	fmt.Printf("\033[1;32mDBG\033[0m: \033[1;37m%s\033[0m\n", s)
}

func getStlGoSlice(w http.ResponseWriter, r *http.Request) {
	// parse the multipart form data
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		dbg("Error parsing multipart form")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	f, h, err := r.FormFile("file")
	if err != nil {
		dbg("Error getting file from form")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer f.Close()

	dst, err := os.Create(h.Filename)
	if err != nil {
		dbg("Error creating file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	s, err := io.Copy(dst, f)
	if err != nil {
		dbg("Error copying file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dbg(fmt.Sprintf("Wrote %d bytes to %s", s, h.Filename))

	dbg(fmt.Sprintf("%+v\n", r))
	fmt.Fprintln(w, "Upload successful")
}

func getStlCura(w http.ResponseWriter, r *http.Request) {}

func getStlPrusa(w http.ResponseWriter, r *http.Request) {}

func main() {
	http.HandleFunc("/stl/goslice", getStlGoSlice)
	http.HandleFunc("/stl/cura", getStlCura)
	http.HandleFunc("/stl/prusa", getStlPrusa)

	err := http.ListenAndServe("0.0.0.0:9999", nil)
	if err != nil {
		panic(err)
	}
}
