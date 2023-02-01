package mycsv

import (
	"fmt"
	"regexp"
	"strconv"
)

func validateKeys(table *Table) (bool, error) {
	colKeysValidator, _ := regexp.Compile("[0-9]")	
	for index, key := range table.colKeys {
		if len(key) == 0{
			return false, fmt.Errorf(
				"column name can't be empty, column index: %d",
				index,
			)
		}
		if colKeysValidator.MatchString(key) {
			return false, fmt.Errorf(
				"column name can't contain any numbers, column index: %d",
				index,
			)
		}
	}

	rowKeysValidator, _ := regexp.Compile("[^0-9]")
	for index, key := range table.rowKeys {
		if len(key) == 0 {
			return false, fmt.Errorf(
				"row number can't be empty, row index: %d",
				index,
			)
		}
		if rowKeysValidator.MatchString(key) {
			return false, fmt.Errorf(
				"row number can't contain anything but digits, row index: %d",
				index,
			)
		}
	}

	return true, nil
}

func validateCells(table *Table) (bool, error) {
	// TODO
	isInt := func (str *string) bool {
		_, err := strconv.Atoi(*str);
		return err == nil;
	}

	isExpression := func (str *string) bool {
		return (*str)[0] == '='
	}

	for _, colKey := range table.colKeys {
		for _, rowKey := range table.rowKeys {
			cell := table.columns[colKey][rowKey];
			if len(cell) == 0 {
				return false, fmt.Errorf(
					("cell can't be empty, cell position: " +
					"column name -> %s, row number -> %s"),
					colKey,
					rowKey,
				)
			}
			if (!isInt(&cell) && !isExpression(&cell)) {
				return false, fmt.Errorf(
					("cell must be an integer or an expression, " +
					"cell position: column name -> %s, row numer -> %s"),
					colKey,
					rowKey,
				);
			}
		}
	}
	return true, nil;
}

func validateExpressions(table *Table) (bool, error) {
	return true, nil
}

func ValidateTable(table *Table) (bool, error) {
	isKeysValid, err := validateKeys(table)
	if err != nil {
		return false, err
	}

	isCellsValid, err := validateCells(table)
	if err != nil {
		return false, err
	}

	isExpressionsValid, err := validateExpressions(table)
	if err != nil {
		return false, err
	}
	
	return (isKeysValid && isCellsValid && isExpressionsValid), nil
}

func ProcessCells() {
	// TODO
}
