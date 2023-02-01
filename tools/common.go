package mycsv

import "strings"

const Space = " "
const EmptyString = ""

func RemoveSpaces(str *string) (string) {
	return strings.ReplaceAll(*str, Space, EmptyString)
}
