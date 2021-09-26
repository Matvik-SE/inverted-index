package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"sync"
	"time"
)

var threads = 3
var word = "and"

var dataDir = "./data/"
var dataExt = ".txt"
var resArr []string
var resMap sync.Map

func main() {
	filesArray := getAllFiles(dataDir, dataExt)
	arrayLength := len(filesArray)

	if arrayLength == 0 {
		log.Panicln("Files folder is empty")
	}
	if threads > arrayLength {
		log.Panicln("Maximum possible threads:", arrayLength)
	}

	sliceLength := int(math.Floor(float64(arrayLength) / float64(threads)))
	startTime := time.Now()

	var wg sync.WaitGroup

	for i := 0; i < arrayLength; i += sliceLength {
		wg.Add(1)
		go func(from int) {
			defer wg.Done()
			to := from + sliceLength

			if to > arrayLength {
				to = arrayLength
			}

			buildInvertedIndex(filesArray[from:to], word)
		}(i)
	}

	wg.Wait()
	elapsedTime := time.Since(startTime)

	log.Printf("Process took %s", elapsedTime)
	log.Println("Total array records:", len(resArr))
	log.Println("Total sync.Map records:", getMapLen(resMap, false))
}

func buildInvertedIndex(arr []string, search string) {
	for _, element := range arr {
		substrCounter := fileSubstrCount(element, search)

		if substrCounter > 0 {
			resMap.Store(element, strconv.Itoa(substrCounter))
			resArr = append(resArr, element+" - "+strconv.Itoa(substrCounter))
		}
	}
}

func getMapLen(syncMap sync.Map, printContent bool) int {
	counter := 0
	record := make(map[interface{}]interface{})

	syncMap.Range(func(k, v interface{}) bool {
		record[k] = v
		counter++
		return true
	})

	if printContent {
		fmt.Println(record)
	}

	return counter
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
