package main

import (
	"encoding/csv"
	// "fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	// generate("big.csv", 10)
	big := Read("big.csv")
	small := Read("small.csv")

	x, y := find(big, small)
	log.Println(x+1, y+1)

}

func Read(name string) [][]int {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal("couldnt open file", err)
	}
	defer file.Close()
	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	result := [][]int{}
	for i := range records {
		row := []int{}
		for j := range records[i] {
			// fmt.Print(records[i][j])
			temp, _ := strconv.Atoi(records[i][j])
			row = append(row, temp)
		}
		result = append(result, row)
	}
	return result
}
func generate(name string, dim int) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatal("couldnt write file", name)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for x := 0; x < dim; x++ {
		row := []string{}
		for y := 0; y < dim; y++ {
			row = append(row, strconv.Itoa(rand.Intn(100)))
		}
		err = writer.Write(row)
		if err != nil {
			log.Fatal("couldnt write row", err)
		}
	}

}
