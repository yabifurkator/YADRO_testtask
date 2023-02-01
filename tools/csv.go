package mycsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
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
	columnNames, err := csvReader.Read();
    if err != nil {
		return nil, err;
    }
	
	firstColumnName := RemoveSpaces(&columnNames[0])
	if firstColumnName != EmptyString {
		return nil, fmt.Errorf("first column name must be empty")
	}	
	columnNames = columnNames[1:]	
	if len(columnNames) == 0 {
		return nil, fmt.Errorf(
			"besides the column with the empty name, " +
			"there must be other columns",
		)
	}

    table := newTable();
    for _, columnName := range columnNames {
		columnName = RemoveSpaces(&columnName)
		_, isKeyAlreadyExists := table.columns[columnName]
		if isKeyAlreadyExists {
			return nil, fmt.Errorf("duplicate column name \"" + columnName + "\"")
		}

		table.columns[columnName] = make(tableColumn)
		table.colKeys = append(table.colKeys, columnName);
    }

    for {
       	record, err := csvReader.Read();
        if err != nil {
			if err == io.EOF {
				if len(table.rowKeys) == 0 {
					return nil, fmt.Errorf(
						"table must contain at least one row",
					)
				}
				break;
			}
			return nil, err;
        }
        
        rowNumber := RemoveSpaces(&record[0]);
		for _, rowKey := range table.rowKeys {
			if rowNumber == rowKey {
				return nil, fmt.Errorf("duplicate row number \"" + rowNumber + "\"")
			}
		}

		table.rowKeys = append(table.rowKeys, rowNumber);
        record = record[1:];
        for index, cell := range record {
			cell = RemoveSpaces(&cell)
			table.columns[table.colKeys[index]][rowNumber] = cell;
        }
    }
	return table, nil;
}
