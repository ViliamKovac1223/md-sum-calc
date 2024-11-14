package main

import (
	"csvsumcalc/dataprocessor"
	"fmt"
	"log"
	"os"
)

func main() {
    fundFile := "data.md"

    // Read data file
    rawData, err := os.ReadFile(fundFile)
    if err != nil {
        log.Fatal(err)
    }

    data := string(rawData)
    fmt.Println(data)

    var reader dataprocessor.DataReader
    fundData, err := reader.ReadFromString(data)
    if err != nil {
        fmt.Fprintln(os.Stderr ,err)
        os.Exit(1)
    }

    fundData.CalcSums()
    fmt.Println(fundData)

    newStringData, err := reader.UpdateString(data, fundData)

    if err != nil {
        fmt.Fprintln(os.Stderr ,err)
        os.Exit(1)
    }

    fmt.Println(newStringData)
}
