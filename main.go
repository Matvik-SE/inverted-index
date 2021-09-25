package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var threads = 6
var dataDir = "./data/"
var index []string

func main() {

}

func checkFileArray(arr []string, search string) {
	for _, element := range arr {
		if fileContains(element, search) {
			index = append(index, element)
			fmt.Println(element)
		}
	}
}

func fileContains(filePath string, search string) bool {
	b, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Print(err)
	}
	content := string(b)

	return strings.Contains(strings.ToLower(content), strings.ToLower(search))
}

//import (
//	"fmt"
//	"github.com/gorilla/mux"
//	"html"
//	"log"
//	"net/http"
//)
//
//func main() {
//	router := mux.NewRouter().StrictSlash(true)
//
//	router.HandleFunc("/find", Index)
//
//	log.Fatal(http.ListenAndServe(":3000", router))
//}
//
//func Index(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
//}
