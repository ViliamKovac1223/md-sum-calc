package main

import (
	"csvsumcalc/dataprocessor"
	"fmt"
	"log"
	"os"
)

func main() {
    fundFile := ""

    for i, arg := range os.Args {
        if arg == "-f" && len(os.Args) > i + 1 {
            fundFile = os.Args[i + 1]
        }
    }

    if fundFile == "" {
        fmt.Fprintln(os.Stderr, "You have to define which file to use with -f <file_name>")
        os.Exit(1)
    }

    // Read data file
    rawData, err := os.ReadFile(fundFile)
    if err != nil {
        log.Fatal(err)
    }
    data := string(rawData)

    // Get structured data from string
    var reader dataprocessor.DataReader
    fundData, err := reader.ReadFromString(data)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    // Recalculate all sums in structured data
    fundData.CalcSums()

    // Get new updated data in string format
    newStringData, err := reader.UpdateString(data, fundData)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    // Write new data to the original file
    os.WriteFile(fundFile, []byte(newStringData), 0664)
}
