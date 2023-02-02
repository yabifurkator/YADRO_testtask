package main

import (
	"fmt"
	mycsv "mainmod/tools"
	"os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("wrong command-line args")
        return
    }

    infilepath := os.Args[1]
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

    err = mycsv.ProcessTable(*table)
    if err != nil {
        fmt.Println("Table is invalid:", err)
        return
    }
    
    fmt.Println(table.ToCsvView())
}
