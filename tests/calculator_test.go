package tests

import (
	"testing"
	"yandex_web_calc/pkg"
)

func TestCalculateValidExpression(t *testing.T) {
	result, err := pkg.Calc("3 + 5 * 2")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expected := 13.0
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestCalculateInvalidExpression(t *testing.T) {
	_, err := pkg.Calc("3 + * 5")
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
}

func TestCalculateDivisionByZero(t *testing.T) {
	_, err := pkg.Calc("10 / 0")
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
}

func TestCalculateParentheses(t *testing.T) {
	result, err := pkg.Calc("(3 + 5) * 2")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expected := 16.0
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestCalculateEmptyExpression(t *testing.T) {
	_, err := pkg.Calc("")
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
}

func TestCalculateSingleNumber(t *testing.T) {
	result, err := pkg.Calc("42")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expected := 42.0
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
