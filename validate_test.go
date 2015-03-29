package main

import "testing"

func containsCode(warnings []int, code int) bool {
	for _, warning := range warnings {
		if warning == code {
			return true
		}
	}
	return false
}

func testValidateLine(t *testing.T, tests []string, code int, hasToContainCode bool) {
	for _, test := range tests {
		warnings := validateLine(test)
		if containsCode(warnings, code) != hasToContainCode {
			t.Errorf("%s (expected %t): \"%s\" (warnings: %v)\n", warningMessages[code], hasToContainCode, test, warnings)
		}
	}
}

func runValidateLineTests(t *testing.T, negativeTests []string, positiveTests []string, code int) {
	testValidateLine(t, negativeTests, code, true)
	testValidateLine(t, positiveTests, code, false)
}

func TestSingleBrace(t *testing.T) {
	code := 1
	negativeTests := []string{"{", "{ ", " {", " { "}
	positiveTests := []string{"public string get_text () {"}
	runValidateLineTests(t, negativeTests, positiveTests, code)
}

func TestEndingSpaces(t *testing.T) {
	code := 2
	negativeTests := []string{"int a = 2; ", "{ ", " } ", " ();    "}
	positiveTests := []string{"int a = 2;", "{", " }", " ();"}
	runValidateLineTests(t, negativeTests, positiveTests, code)
}

func Test4SpacesIndenting(t *testing.T) {
	code := 3
	negativeTests := []string{"      int a = 2;", " {", "\t}", "    \t();"}
	positiveTests := []string{"    int a = 2;", "        }", "();"}
	runValidateLineTests(t, negativeTests, positiveTests, code)
}

func TestOpeningParenthese(t *testing.T) {
	code := 4
	negativeTests := []string{"=()", "+()", " ()()", "=((()))", "= (3 +(4))"}
	positiveTests := []string{" ()", "+ ()", " () () ", "= (((3 + ())))"}
	runValidateLineTests(t, negativeTests, positiveTests, code)
}

func TestEqualWithSpaces(t *testing.T) {
	code := 5
	negativeTests := []string{"b=", "=3", "a=3", " =(3 + 4)"}
	positiveTests := []string{"b = 2", "a!= 2", " =>", "<= ", ">= ", " == "}
	runValidateLineTests(t, negativeTests, positiveTests, code)
}

func TestLineLengthGreaterThan120(t *testing.T) {
	code := 6
	string120 := "kkzvt84AlytfJRgugEo2mJs0K6utlfe5zrwDCC8lfXPt3GMFLAfDhxFN5nJX9tzkpyQMQ8lHQTuvC0ZEjmnrEmQ5rOcAHVXLGxK5Y28xmURvWn1EMjytRFcm"
	negativeTests := []string{string120 + "a", string120 + "abcdef", string120 + string120}
	positiveTests := []string{string120, "hello", string120[:119], ""}
	runValidateLineTests(t, negativeTests, positiveTests, code)
}

func TestGLibNotNecessary(t *testing.T) {
	code := 7
	negativeTests := []string{"    GLib.print (\"Hi\");", "GLib.println()"}
	positiveTests := []string{"    print (\"Hi\");", "println()"}
	runValidateLineTests(t, negativeTests, positiveTests, code)
}
