package tests

import (
	mycsv "mainmod/tools"
	"regexp"
	"testing"
)

func TestRemoveSpaces(t *testing.T) {
	correct_transformations := [][]string {
		{"Hello, what is your name?", "Hello,whatisyourname?"},
		{" SpacesBefore", "SpacesBefore"},
		{"SpacesAfter ", "SpacesAfter"},
		{"Spaces Between", "SpacesBetween"},
		{" Spaces E v e r y w h e r e ! ! ! ", "SpacesEverywhere!!!"},
		{" ", ""},
		{"", ""},
	}
	for _, testCase := range correct_transformations {
		removeSpacesFunctionResult := mycsv.RemoveSpaces(&testCase[0])
		if removeSpacesFunctionResult != testCase[1] {
			t.Errorf(
				"invalid \"RemoveSpaces\" function\n" +
				"init string: \"" + testCase[0] + "\"\n" +
				"result:   \"" + removeSpacesFunctionResult + "\"\n" +
				"expected: \"" + testCase[1] + "\"\n",
			)
		}
	}
}

func TestIsInt(t *testing.T) {
	correct_integer_strings := []string{
		"0",
		"1",
		"-1",
		"999",
		"+9",
	}
	incorrect_integer_strings := []string {
		"12 1",
		"1.99",
		"1.00",
		"aboba",
		"12aboba",
		" ",
	}
	for _, str := range correct_integer_strings {
		if !mycsv.IsInt(&str) {
			t.Errorf(
				"invalid \"IsInt\" function\n" +
				"test string: \"" + str + "\"\n" +
				"expected: true",
				"result:   false",
			)
		}
	}
	for _, str := range incorrect_integer_strings {
		if mycsv.IsInt(&str) {
			t.Errorf(
				"invalid \"IsInt\" function\n" +
				"test string: \"" + str + "\"\n" +
				"expected: false\n" +
				"result:   true\n",
			)
		}
		
	}
}

func testRegexpr(
	t *testing.T,
	correct_test_strings []string,
	incorrect_test_strings []string,
	regexpr string,
) {
	validator := regexp.MustCompile("^" + regexpr + "$")
	for _, str := range correct_test_strings {
		if !validator.MatchString(str) {
			t.Fatalf(
				"invaild regexpr \"" + validator.String() + "\", " +
				"didn't match string \"" + str + "\"",
			)
		}
	}
	for _, str := range incorrect_test_strings {
		if validator.MatchString(str) {
			t.Fatalf(
				"invaild regexpr \"" + validator.String() + "\", " +
				"match string \"" + str + "\"",
			)
		}
	}
}

func TestColumnNameRegexpr(t *testing.T) {
	correct_test_strings := []string {
		"A",
		"B",
		"Cell",
		"ABOBA",
		"hello",
	}
	incorrect_test_strings := []string {
		"1",
		"2",
		"99",
		"A1",
		"2B",
		"A_",
		"!@#",
		" ",
	}
	testRegexpr(
		t,
		correct_test_strings,
		incorrect_test_strings,
		mycsv.ColumnNameRegexpr,
	)
}

func TestRowNumberRegexpr(t *testing.T) {
	correct_test_strings := []string {
		"1",
		"123",
		"30",
		"5",
		"007",
	}
	incorrect_test_strings := []string {
		"1A",
		"A1",
		"-3",
		"hello",
		"&^%$",
		" ",
	}
	testRegexpr(
		t,
		correct_test_strings,
		incorrect_test_strings,
		mycsv.RowNumberRegexpr,
	)
}

func TestCellPosRegexpr(t *testing.T) {
	correct_test_strings := []string {
		"A1",
		"B3",
		"CELL59",
		"C8",
		"Aboba99",	
	}
	incorrect_test_strings := []string {
		"-B1",
		"+A5",
		"1A",
		"hello?",
		"&^%$",
		" ",
	}
	testRegexpr(
		t,
		correct_test_strings,
		incorrect_test_strings,
		mycsv.CellPosRegexpr,
	)
}

func TestInteger(t *testing.T) {
	correct_test_strings := []string {
		"1",
		"-2",
		"+3",
		"99",
		"05",
		"0",
	}
	incorrect_test_strings := []string {
		"+-1",
		"1-",
		"howareyou",
		"0 9",
		"12.03",
		" ",
	}
	testRegexpr(
		t,
		correct_test_strings,
		incorrect_test_strings,
		mycsv.IntegerRegexpr,
	)
}

func TestOperandRegexpr(t *testing.T) {
	correct_test_strings := []string {
		"A1",
		"B3",
		"CELL59",
		"C8",
		"Aboba99",
		"-2",
		"+3",
		"99",
		"05",
		"0",
	}
	incorrect_test_strings := []string {
		"+-1",
		"1-",
		"howareyou",
		"0 9",
		"12.03",
		" ",
	}
	testRegexpr(
		t,
		correct_test_strings,
		incorrect_test_strings,
		mycsv.OperandRegexpr,
	)
}

func TestOperationRegexpr(t *testing.T) {
	correct_test_strings := []string {
		"+",
		"-",
		"*",
		"/",
	}
	incorrect_test_strings := []string {
		"imfine",
		"//",
		"x",
		"++",
		"--",
		"*=",
		" ",
	}
	testRegexpr(
		t,
		correct_test_strings,
		incorrect_test_strings,
		mycsv.OperationRegexpr,
	)
}
