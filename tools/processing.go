package mycsv

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

func processExpression(expression *string, table *Table) (*string, error) {
	expression_to_drop := *expression

	operandExpr := regexp.MustCompile(OperandRegexpr)
	
	operand1Str := operandExpr.FindString(expression_to_drop)
	expression_to_drop = strings.Replace(
		expression_to_drop,
		"=" + operand1Str,
		"", 
		1,
	)

	operationStr := string(expression_to_drop[0])
	expression_to_drop = expression_to_drop[1:]
	
	operand2Str := operandExpr.FindString(expression_to_drop)

	operandToInt := func (operand *string) (*int, error) {
		value, err := strconv.Atoi(*operand)
		if err == nil {
			return &value, nil
		}
		
		colKey := regexp.MustCompile(ColumnNameRegexpr).FindString(*operand)
		rowKey := regexp.MustCompile(RowNumberRegexpr).FindString(*operand)
		
		var ok bool
		_, ok = table.columns[colKey]; if !ok {
			return nil, fmt.Errorf(
				"invalid expression \"" + *expression + "\"\n" +
				"column \"" + colKey + "\" doesn't exist",
			)
		}
		_, ok = table.columns[colKey][rowKey]; if !ok {
			return nil, fmt.Errorf(	
				"invalid expression \"" + *expression + "\"\n" +
				"row \"" + rowKey + "\" doesn't exist",
			)
		}

		cell := table.columns[colKey][rowKey]
		value, err = strconv.Atoi(cell)
		if err != nil {
			return nil, fmt.Errorf(
				"invalid operand \"" + *operand +
				"\" in expression \"" + *expression + "\"\n" +
				"not integer",
			)
		}

		return &value, nil
	}

	operand1, err := operandToInt(&operand1Str)
	if err != nil {
		return nil, err
	}
	operand2, err := operandToInt(&operand2Str)
	if err != nil {
		return nil, err
	}

	var result int

	switch operationStr {
		case "+": result = *operand1 + *operand2
		case "-": result = *operand1 - *operand2
		case "*": result = *operand1 * *operand2
		case "/":
			if *operand2 == 0 {
				return nil, fmt.Errorf(
					"invalid second operand value in expression \"" +
					*expression + "\"\nzero division",
				)
			}
			result = *operand1 / *operand2
	}

	resultStr := strconv.Itoa(result)
	return &resultStr, nil
}

func ProcessTable(table Table) (error) {
	_, err := validateKeys(&table)
	if err != nil {
		return err
	}

	isExpression := func (str *string) bool {
		exprRegexpr := "^=" + OperandRegexpr + OperationRegexpr + OperandRegexpr + "$"
		expressionValidator := regexp.MustCompile(exprRegexpr);
		
		return expressionValidator.MatchString(*str)
	}

	for _, colKey := range table.colKeys {
		for _, rowKey := range table.rowKeys {
			cell := table.columns[colKey][rowKey];
			if cell == EmptyString {
				return fmt.Errorf(
					("cell can't be empty" +
					"\ncolumn name: %s, row number: %s"),
					colKey,
					rowKey,
				)
			}
			if !IsInt(&cell) {
				if !isExpression(&cell) {
					return fmt.Errorf(
						("cell must be an integer or an expression " +
						"with correct syntax\ncolumn name: %s, row numer: %s"),
						colKey,
						rowKey,
					);
				}
				newCell, err := processExpression(&cell, &table)
				if err != nil {
					return err
				}
				table.columns[colKey][rowKey] = *newCell
			}
		}
	}
	return nil
}
