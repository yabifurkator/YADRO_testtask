package main

import (
	"fmt"
	mycsv "mainmod/tools"
	"os"
)

func main() {
    infilepath := "input.txt"
    file, err := os.Open(infilepath)
    if err != nil {
        fmt.Println("Error opening input file:", err)
        return
    }
    defer file.Close()

    table, err := mycsv.GetTableFromCsvFile(file)
    if err != nil {
        fmt.Println("Error getting table from file:", err)
        return
    }
    fmt.Println(table)

    _, err = mycsv.ValidateTable(table)
    if err != nil {
        fmt.Println("Table is invalid:", err)
        return
    }
}
