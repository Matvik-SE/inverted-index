package main

import (
	"log"
	"math"
	"strconv"
	"sync"
	"time"
)

var resArr []string
var resMap sync.Map

func runIndexThreads(filesArray []string, threads int, word string) map[string]interface{} {
	arrayLength := len(filesArray)
	resArr = nil
	resMap = sync.Map{}

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
	mapLen, mappedRes := mapSyncMap(&resMap)

	log.Println("--------- Benchmark ----------")
	log.Printf("Number of threads %d", threads)
	log.Printf("Process took %s", elapsedTime)
	log.Println("Total regular array records:", len(resArr))
	log.Println("Total sync.Map records:", mapLen)
	log.Println("------------------------------")

	return mappedRes
}

func buildInvertedIndex(arr []string, search string) {
	for _, element := range arr {
		substrCounter := fileSubstrCount(element, " "+search+" ")

		if substrCounter > 0 {
			resMap.Store(element, strconv.Itoa(substrCounter))
			resArr = append(resArr, element+" - "+strconv.Itoa(substrCounter))
		}
	}
}
