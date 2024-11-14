package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
    // fundFile := "fund.yaml"
    fundFile := "data.md"

    // Read the yaml file
    rawData, err := os.ReadFile(fundFile)
    if err != nil {
        log.Fatal(err)
    }

    data := string(rawData)
    fmt.Println(data)

    var reader DataReader
    fundData, err := reader.ReadFromString(data)
    if err != nil {
        fmt.Fprintln(os.Stderr ,err)
        os.Exit(1)
    }

    fundData.calcSums()
    fmt.Println(fundData)

    newStringData, err := reader.UpdateString(data, fundData)

    if err != nil {
        fmt.Fprintln(os.Stderr ,err)
        os.Exit(1)
    }

    fmt.Println(newStringData)
}
