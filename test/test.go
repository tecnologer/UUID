package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/tecnologer/uuid"
)

const fileName = "results.txt"

var testresults = flag.Bool("t", false, "Check if exists some results are duplicated")
var count = flag.Int("c", 0, "Number of UUID to generate in the current test")

func main() {
	flag.Parse()

	var f *os.File
	var err error
	//fmt.Println(UUID.GetUUID())
	if _, err = os.Stat(fileName); os.IsNotExist(err) {
		f, err = os.Create(fileName)
	} else {
		f, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	}
	defer f.Close()

	max := *count
	if max == 0 {
		rand.Seed(time.Now().UnixNano())
		max = rand.Intn(100-10) + 10
	}

	for i := 0; i < max; i++ {
		f.WriteString(fmt.Sprintf("%s\n", UUID.GetUUID()))
	}

	if *testresults {
		fmt.Println("Testing results.")
		_ = testResults()
	}
}

func testResults() bool {
	inFile, _ := os.Open(fileName)
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	results := make(map[string]int)

	count := 0
	duplicated := 0
	for scanner.Scan() {
		uuid := scanner.Text()
		results[uuid]++

		if results[uuid] > 1 {
			duplicated++
			fmt.Println(uuid)
		}

		count++
	}
	fmt.Printf("Tested %d UUID and found %d UUID duplicated.\n", count, duplicated)
	return duplicated > 0
}
