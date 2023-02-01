package mycsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type tableColumn map[string]string
type Table struct{
	colKeys []string
	rowKeys []string
	columns map[string]tableColumn
}

func newTable() *Table {
	table := new(Table)
	table.columns = make(map[string]tableColumn)
	return table
}

func GetTableFromCsvFile(file *os.File) (*Table, error) {
	csvReader := csv.NewReader(file);
	columns, err := csvReader.Read();
    if err != nil {
		return nil, err;
    }
	
	firstColumnName := RemoveSpaces(&columns[0])
	if len(firstColumnName) != 0 {
		return nil, fmt.Errorf("first column name must be empty")
	}	
	columns = columns[1:]	
	if len(columns) == 0 {
		return nil, fmt.Errorf(
			"besides the column with the empty name, " +
			"there must be other columns",
		)
	}

    columns = columns[1:];
    table := newTable();
    for _, str := range columns {
		str := strings.ReplaceAll(str, " ", "")
		table.columns[str] = make(tableColumn)
		table.colKeys = append(table.colKeys, str);
    }

    for {
        record, err := csvReader.Read();
        if err != nil {
			if err == io.EOF {
				if len(table.rowKeys) == 0 {
					return nil, fmt.Errorf("ERROR! %s", "no table rows")
				}
				break;
			}
			return nil, fmt.Errorf("ERROR! %s", err);
        }
        
        rowNumber := record[0];
		table.rowKeys = append(table.rowKeys, rowNumber);
        record = record[1:];
        for index, str := range record {
			table.columns[table.colKeys[index]][rowNumber] = str;
        }
    }
	return table, nil;
}
