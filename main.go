package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
	"sync"
	"time"
)

var threads = 6
var dataDir = "./data/"
var index []string

func main() {
	res := getDirFiles(dataDir)
	length := len(res)
	part := int(math.Ceil(float64(length) / float64(threads)))

	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)

	for i := 0; i < threads; i++ {
		start := i * part
		stop := start + part
		part := res[start:stop]

		go func() {
			defer wg.Done()
			checkFileArray(part, "police")
		}()
	}

	wg.Wait()

	//fmt.Printf("%v", index)

	elapsed := time.Since(start)
	log.Printf("Process took %s", elapsed)

	//foo := fileContains(res[0], "lower rating")
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
