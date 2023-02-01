package mycsv

import "strings"

const Space = " "
const EmptyString = ""
const ColumnNameRegexpr = "[a-zA-Z]+"
const RowNumberRegexpr = "[0-9]+"
const OperandRegexpr = "(" + ColumnNameRegexpr + RowNumberRegexpr + ")"

func RemoveSpaces(str *string) (string) {
	return strings.ReplaceAll(*str, Space, EmptyString)
}
