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

// func validateCells(table *Table) (bool, error) {
// 	isInt := func (str *string) bool {
// 		_, err := strconv.Atoi(*str);
// 		return err == nil;
// 	}

// 	isExpression := func (str *string) bool {
// 		operandRegexpr := "(" + ColumnNameRegexpr + RowNumberRegexpr + ")"
// 		exprRegexpr := "^=" + operandRegexpr + "[\\+\\-\\*\\/]" + operandRegexpr + "$"
// 		expressionValidator := regexp.MustCompile(exprRegexpr);
		
// 		return expressionValidator.MatchString(*str)
// 	}

// 	for _, colKey := range table.colKeys {
// 		for _, rowKey := range table.rowKeys {
// 			cell := table.columns[colKey][rowKey];
// 			if cell == EmptyString {
// 				return false, fmt.Errorf(
// 					("cell can't be empty" +
// 					"\ncolumn name: %s, row number: %s"),
// 					colKey,
// 					rowKey,
// 				)
// 			}
// 			if (!isInt(&cell) && !isExpression(&cell)) {
// 				return false, fmt.Errorf(
// 					("cell must be an integer or an expression " +
// 					"with correct syntax\ncolumn name: %s, row numer: %s"),
// 					colKey,
// 					rowKey,
// 				);
// 			}
// 		}
// 	}
// 	return true, nil;
// }

// func validateTable(table *Table) (bool, error) {
// 	isKeysValid, err := validateKeys(table)
// 	if err != nil {
// 		return false, err
// 	}

// 	isCellsValid, err := validateCells(table)
// 	if err != nil {
// 		return false, err
// 	}

// 	return (isKeysValid && isCellsValid), nil
// }

// func ProcessTable(table *Table) (*string, error) {
// 	_, err := validateTable(table);
// 	if err != nil {
// 		return nil, err
// 	}

// 	tbStr := "table_str"
// 	return &tbStr, nil
// }

func processExpression(expression *string, table *Table) (string, error) {
	operandExpr := regexp.MustCompile(OperandRegexpr)
	operands := operandExpr.FindAllString(*expression, 2)

	type cellPos struct {
		colKey string
		rowKey string
	}
	getRowColFromOperand := func (operand *string) (*cellPos) {
		pos := new(cellPos)
		pos.colKey = regexp.MustCompile(ColumnNameRegexpr).FindString(*operand)
		pos.rowKey = regexp.MustCompile(RowNumberRegexpr).FindString(*operand)
		return pos
	}
	operand1 := getRowColFromOperand(&operands[0])
	operand2 := getRowColFromOperand(&operands[1])

	var ok bool
	_, ok = table.columns[operand1.colKey]; if !ok {
		return 
	}
	_, ok = table.columns[operand1.colKey][operand1.rowKey]; if !ok {

	}

	fmt.Println("op1:", operand1, "op2:", operand2)
	// aboba := strconv.FormatFloat()
	// operand1 := 1.0
	// operand2 := 2.0
	// operation := "/"
	// var result float64
	// switch operation {
	// case "+": result = float64(operand1) + float64(operand2)
	// case "-": result = float64(operand1) - float64(operand2)
	// case "*": result = float64(operand1) * float64(operand2)
	// case "/": result = float64(operand1) / float64(operand2)
	// }
	return *expression, nil
}

func ProcessTable(table Table) (*Table, error) {
	isInt := func (str *string) bool {
		_, err := strconv.Atoi(*str);
		return err == nil;
	}

	isExpression := func (str *string) bool {
		exprRegexpr := "^=" + OperandRegexpr + "[\\+\\-\\*\\/]" + OperandRegexpr + "$"
		expressionValidator := regexp.MustCompile(exprRegexpr);
		
		return expressionValidator.MatchString(*str)
	}

	for _, colKey := range table.colKeys {
		for _, rowKey := range table.rowKeys {
			cell := table.columns[colKey][rowKey];
			if cell == EmptyString {
				return nil, fmt.Errorf(
					("cell can't be empty" +
					"\ncolumn name: %s, row number: %s"),
					colKey,
					rowKey,
				)
			}
			if !isInt(&cell) {
				if !isExpression(&cell) {
					return nil, fmt.Errorf(
						("cell must be an integer or an expression " +
						"with correct syntax\ncolumn name: %s, row numer: %s"),
						colKey,
						rowKey,
					);
				}
				newCell, err := processExpression(&cell, &table)
				if err != nil {
					return nil, err
				}
				table.columns[colKey][rowKey] = newCell
			}
		}
	}
	return &table, nil;
}
