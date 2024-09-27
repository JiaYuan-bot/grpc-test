package main

import (
	"time"
)

func genarateGlobalData() {
	const size = 128 * 1024 * 1024

	// Create a slice to hold the int64 elements
	GlobalData = make([]int64, size)

	// Optionally, initialize the array with values
	for i := int64(0); i < size; i++ {
		GlobalData[i] = i // Or any other logic for initialization
	}
}

func main() {
	var err error
	var hbc *client

	genarateGlobalData()

	hbs := NewHeartBeater()

	if hbc, err = NewClient(); err != nil {
		panic(err)
	}

	hbs.Start()

	hbc.StartHeartbeat()

	time.Sleep(time.Hour)

}
