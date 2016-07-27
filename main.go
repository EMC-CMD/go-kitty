package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const FILEPATH = "/mnt/data/goapp.txt"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
	fmt.Fprintln(w, "VCAP_SERVICES:", os.Getenv("VCAP_SERVICES"))
}

func read(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile(FILEPATH)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, "file content: %s", string(content))
}

func write(w http.ResponseWriter, r *http.Request) {
	s := "Chris and I were here"
	err := ioutil.WriteFile(FILEPATH, []byte(s), os.ModePerm)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, "saved %s to %s successfully", s, FILEPATH)
}

func main() {
	http.HandleFunc("/read", read)
	http.HandleFunc("/write", write)
	http.HandleFunc("/", handler)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
