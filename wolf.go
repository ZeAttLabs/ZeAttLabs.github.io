package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// ShowFile - display html file
func ShowFile(w http.ResponseWriter, fileName string) {
	var htmlString string
	htmlBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		htmlString = "File Not Found : " + fileName
	} else {
		htmlString = string(htmlBytes)
	}
	fmt.Fprintf(w, htmlString)
}

func test(w http.ResponseWriter, r *http.Request) {
	ShowFile(w, "index.html")
}

func kill(w http.ResponseWriter, r *http.Request) {
	ShowFile(w, "kill.html")
}

func js(w http.ResponseWriter, r *http.Request) {
	ShowFile(w, "web3.min.js")
}

func main() {
	http.Handle("/inc/", http.StripPrefix("/inc/", http.FileServer(http.Dir("inc"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.HandleFunc("/", test)
	http.HandleFunc("/kill", kill)
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println(err)
	}
}
