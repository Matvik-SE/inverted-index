package main

import (
	"sync"
)

func mapSyncMap(syncMap *sync.Map) (int, map[string]interface{}) {
	counter := 0
	record := make(map[string]interface{})

	syncMap.Range(func(k interface{}, v interface{}) bool {
		record[k.(string)] = v
		counter++
		return true
	})

	return counter, record
}
