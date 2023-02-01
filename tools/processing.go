package mycsv

import (
	"fmt"
	"regexp"
	"strconv"
)

func validateKeys(table *Table) (bool, error) {

	colRegexpr := "^" + ColumnNameRegexpr + "$"
	colKeysValidator := regexp.MustCompile(colRegexpr)
	for index, key := range table.colKeys {
		if key == EmptyString {
			return false, fmt.Errorf(
				"column name can't be empty, column index: %d",
				index + 1,
			)
		}
		if !colKeysValidator.MatchString(key) {
			return false, fmt.Errorf(
				"column name must contain only EN letters (lower of upper case), column index: %d",
				index + 1,
			)
		}
	}

	rowRegexpr := "^" + RowNumberRegexpr + "$"
	rowKeysValidator := regexp.MustCompile(rowRegexpr)
	for index, key := range table.rowKeys {
		if key == EmptyString {
			return false, fmt.Errorf(
				"row number can't be empty, row index: %d",
				index,
			)
		}
		if !rowKeysValidator.MatchString(key) {
			return false, fmt.Errorf(
				"row number must contain only digits, row index: %d",
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
		operandRegexpr := "(" + ColumnNameRegexpr + RowNumberRegexpr + ")"
		exprRegexpr := "^=" + operandRegexpr + "[\\+\\-\\*\\/]" + operandRegexpr + "$"
		fmt.Println(exprRegexpr)
		expressionValidator := regexp.MustCompile(exprRegexpr);
		
		return expressionValidator.MatchString(*str)
	}

	for _, colKey := range table.colKeys {
		for _, rowKey := range table.rowKeys {
			cell := table.columns[colKey][rowKey];
			if cell == EmptyString {
				return false, fmt.Errorf(
					("cell can't be empty" +
					"\ncolumn name: %s, row number: %s"),
					colKey,
					rowKey,
				)
			}
			if (!isInt(&cell) && !isExpression(&cell)) {
				return false, fmt.Errorf(
					("cell must be an integer or an expression " +
					"with correct syntax\ncolumn name: %s, row numer: %s"),
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
