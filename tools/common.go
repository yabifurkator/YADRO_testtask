package mycsv

import (
	"strconv"
	"strings"
)

const Space = " "
const EmptyString = ""
const ColumnNameRegexpr = "[a-zA-Z]+"
const RowNumberRegexpr = "[0-9]+"
const CellPosRegexpr = "(" + ColumnNameRegexpr + RowNumberRegexpr + ")"
const IntegerRegexpr = "([-+]?[0-9]+)"
const OperandRegexpr = "(" + CellPosRegexpr + "|" + IntegerRegexpr + ")"
const OperationRegexpr = "[\\+\\-\\*\\/]"

func RemoveSpaces(str *string) (string) {
	return strings.ReplaceAll(*str, Space, EmptyString)
}

func IsInt(str *string) bool {
	_, err := strconv.Atoi(*str);
	return err == nil;
}
