package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var extra_life = false
var increase = false
var decrease = false

func reset_state() {
	extra_life = false
	increase = false
	decrease = false
}

func CheckDiff(val1 int, val2 int) bool {
	abs_diff := IntAbs(val1 - val2)
	if abs_diff > 3 || abs_diff < 1 {
		if extra_life {
			extra_life = false
		} else {
			return false
		}
	}
	return true
}

func CheckMono(val1 int, val2 int) bool {
	diff := val1 - val2
	if diff >= 0 {
		decrease = true
	} else {
		increase = true
	}

	if increase && decrease {
		if extra_life {
			extra_life = false
			if diff >= 0 {
				increase = false
			} else {
				decrease = false
			}
		} else {
			return false
		}
	}
	return true
}

func LoadData(filename string) ([][]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(f)
	result := [][]int{}
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		split := strings.Split(line, ",")
		report := []int{}
		for _, val := range split {
			int_val, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				report = append(report, int_val)
			}
		}
		result = append(result, report)
	}
	return result, nil
}

func main() {
	reports, err := LoadData("input.csv")
	if err != nil {
		fmt.Println(err.Error())
	}
	count := 0
	for _, report := range reports {
		safe := true
		reset_state()
		for i := 0; i < len(report)-1; i++ {
			safe = CheckDiff(report[i], report[i+1])
			safe = CheckMono(report[i], report[i+1])
		}
		if safe {
			count++
		}
	}
	fmt.Print(count)
}
