package main

import "testing"

func TestSimpleExpression(t *testing.T){
	expected := 42
	actual := calculate(21, 2, "*")
	if actual != expected {
		t.Error("Test failed : 21 * 2 = 42")
	}

	actual = calculate(21, 21, "+")
	if actual != expected {
		t.Error("Test failed : 21 + 21 = 42")
	}

	actual = calculate(63, 21, "-")
	if actual != expected {
		t.Error("Test failed : 63 - 21 = 42")
	}

	actual = calculate(84, 2, "/")
	if actual != expected {
		t.Error("Test failed : 84 / 2 = 42")
	}
}

func TestExpressionNegativeResult(t *testing.T){
	expected := -42
	actual := calculate(-21, 2, "*")
	if actual != expected {
		t.Error("Test failed : -21 * 2 = -42")
	}

	actual = calculate(-21, -21, "+")
	if actual != expected {
		t.Error("Test failed : -21 + -21 = -42")
	}

	actual = calculate(-63, -21, "-")
	if actual != expected {
		t.Error("Test failed : -63 - -21 = -42")
	}

	actual = calculate(-84, 2, "/")
	if actual != expected {
		t.Error("Test failed : -84 / 2 = -42")
	}

	actual = calculate(21, -2, "*")
	if actual != expected {
		t.Error("Test failed : 21 * -2 = -42")
	}

	actual = calculate(84, -2, "/")
	if actual != expected {
		t.Error("Test failed : 84 / -2 = -42")
	}
}

func TestSimpleExpressionArray(t *testing.T){
	expected := "42"
	input := []string{"21", "+", "21"}
	actual := calculateParsedExpression(input)
	if actual != expected {
		t.Error("Test failed : 21 + 21 = 42")
	}

	input = []string{"21", "*", "2"}
	actual = calculateParsedExpression(input)
	if actual != expected {
		t.Error("Test failed : 21 * 2 = 42")
	}

	input = []string{"63", "-", "21"}
	actual = calculateParsedExpression(input)
	if actual != expected {
		t.Error("Test failed : 63 - 21 = 42")
	}

	input = []string{"84", "/", "2"}
	actual = calculateParsedExpression(input)
	if actual != expected {
		t.Error("Test failed : 84 / 2 = 42")
	}
}

func TestSimpleExpressionArrayNegativeResult(t *testing.T){
	expected := "-42"
	input := []string{"-21", "+", "-21"}
	actual := calculateParsedExpression(input)
	if actual != expected {
		t.Error("Test failed : -21 + -21 = -42")
	}

	input = []string{"-21", "*", "2"}
	actual = calculateParsedExpression(input)
	if actual != expected {
		t.Error("Test failed : -21 * 2 = -42")
	}

	input = []string{"21", "*", "-2"}
	actual = calculateParsedExpression(input)
	if actual != expected {
		t.Error("Test failed : 21 * -2 = -42")
	}

	input = []string{"-63", "-", "-21"}
	actual = calculateParsedExpression(input)
	if actual != expected {
		t.Error("Test failed : -63 - -21 = -42")
	}

	input = []string{"-84", "/", "2"}
	actual = calculateParsedExpression(input)
	if actual != expected {
		t.Error("Test failed : -84 / 2 = -42")
	}

	input = []string{"84", "/", "-2"}
	actual = calculateParsedExpression(input)
	if actual != expected {
		t.Error("Test failed : 84 / -2 = -42")
	}
}

func TestComplexExpressionArray(t *testing.T){
	expected := "42"
	input := []string{"-21", "*", "-3", "+", "15", "+", "6", "/", "2"}
	actual := calculateParsedExpression(input)
	if actual != expected {
		t.Error("Test failed : -21 * -3 + 15 + 6 / 2 = 42")
	}

	expected = "-42"
	input = []string{"14", "*", "2", "-", "-15", "*", "3", "+", "7", "/", "-2", "+", "26"}
	actual = calculateParsedExpression(input)
	if actual != expected {
		t.Error("Test failed : 14 * 2 - -15 * 3 + 7 / -2 + 26 = -42")
	}
}

func TestParseExpression(t *testing.T){
	expected := []string{"-21", "*", "-3", "+", "15", "+", "6", "/", "2"}
	input := "-21 * -3 + 15 + 6 / 2"
	actual := parseExpression(input)
	if Equal(actual,expected) != true {
		t.Error("Test failed : -21 * -3 + 15 + 6 / 2")
	}
}

func TestParseMissingOperatorExpression(t *testing.T){
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()

	input := "-21 -3 + 15 + 6 / 2"
	parseExpression(input)
}

func TestParseMissingNumberExpressionPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()

	input := "-21 * + -3 + 15 + 6 / 2"
	parseExpression(input)
}

func TestCalculatePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()

	calculate(42, 42, "%")
}

func TestCalculateDivisionByZeroPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()

	calculate(42, 0, "/")
}

func Equal(a, b []string) bool {
	
		if a == nil && b == nil { 
			return true; 
		}
	
		if a == nil || b == nil { 
			return false; 
		}
	
		if len(a) != len(b) {
			return false
		}
	
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
	
		return true
	}
