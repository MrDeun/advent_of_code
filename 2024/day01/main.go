package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

func readCSVFIle(filename string) (data []byte, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err = io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func parseCSV(data []byte) (*csv.Reader, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	return reader, nil
}

func get_csv_data(csv_file *csv.Reader) ([]int, []int, error) {
	r1 := []int{}
	r2 := []int{}

	for {
		record, err := csv_file.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, nil, err
		}

		val1, err := strconv.ParseInt(record[0], 10, 32)
		val2, err := strconv.ParseInt(record[1], 10, 32)
		r1 = append(r1, int(val1))
		r2 = append(r2, int(val2))
	}
	return r1, r2, nil
}

func main() {
	raw_data, _ := readCSVFIle("input.csv")
	csv_file, _ := parseCSV(raw_data)

	r1, r2, err := get_csv_data(csv_file)
	if err != nil {
		fmt.Printf("Damn it, something broke... %s", err.Error())
		os.Exit(1)
	}

	slices.Sort(r1)
	slices.Sort(r2)

	sum := 0
	for _, pattern := range r1 {
		count := 0
		for _, val := range r2 {
			if val == pattern {
				count++
			}
		}
		sum += pattern * count
	}
	fmt.Printf("%d\n", sum)
}
