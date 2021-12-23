package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type DateTime struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

func main() {
	fmt.Println(time.Now().Format("2006-01-02T15:04:05.000 MST"))
	inputData := []DateTime{}
	getInput(&inputData)

	for _, f := range []func(a, b DateTime) bool{
		isTimeOverlap,
		includeSameDay,
		isShorter,
	} {
		printFunctionName(f)
		m := make([][]bool, len(inputData))
		for i, v := range inputData {
			m[i] = make([]bool, len(inputData))
			for j, w := range inputData {
				m[i][j] = f(v, w)
			}
		}
		printBoolMatrix(m)
	}
}

func getInput(inputData interface{}) {
	if len(os.Args) > 1 {
		data, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := json.NewDecoder(data).Decode(inputData); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		if err := json.NewDecoder(os.Stdin).Decode(inputData); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func isTimeOverlap(a, b DateTime) bool {
	if a.From.After(b.To) {
		return false
	} else if a.To.Before(b.From) {
		return false
	}
	return true
}

func deleteTime(d DateTime) DateTime {
	const layout = "2006-01-02"
	d.From, _ = time.Parse(layout, d.From.Format(layout))
	d.To, _ = time.Parse(layout, d.To.Format(layout))
	return d
}

func includeSameDay(a, b DateTime) bool {
	return isTimeOverlap(
		deleteTime(a),
		deleteTime(b),
	)
}

func isShorter(a, b DateTime) bool {
	return a.To.Sub(a.From) < b.To.Sub(b.From)
}

func printFunctionName(f interface{}) {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	arr := strings.Split(name, ".")
	fmt.Fprintln(os.Stdout, arr[len(arr)-1])
}

func printBoolMatrix(m [][]bool) {
	for _, v := range m {
		for _, w := range v {
			if w {
				fmt.Fprint(os.Stdout, "T ")
			} else {
				fmt.Fprint(os.Stdout, "F ")
			}
		}
		fmt.Fprintln(os.Stdout)
	}
}
