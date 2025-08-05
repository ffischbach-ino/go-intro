package greetings

import (
	"regexp"
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	name := "Fynn"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Greet(name)

	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Greet("%s") = %s, %s, want match for %s`, name, msg, err, want)
	}
}

func TestGreetEmpty(t *testing.T) {
	emptyName := ""
	expectedErrMessage := "empty name"
	msg, err := Greet(emptyName)

	if err == nil {
		t.Errorf(`GreetMultiple with empty name list should throw error. Message = %s`, msg)
	}

	if !errorContains(err, expectedErrMessage) {
		t.Errorf(`Unpected error occurred. Expected %s. Actual %s`, expectedErrMessage, err)
	}
}

func TestGreetMultiple(t *testing.T) {
	names := []string{"Thomas", "Hans", "Leon"}
	messages, err := GreetMultiple(names)

	if len(messages) != len(names) {
		t.Errorf(`GreetMultiple(%s) = %s, %s, should contain %d`, names, messages, err, len(names))
	}
}

func TestGreetMultipleEmpty(t *testing.T) {
	names := []string{}
	messages, err := GreetMultiple(names)

	expectedErrMessage := "empty name list"

	if err == nil {
		t.Errorf(`GreetMultiple with empty name list should throw error. Messages = %s`, messages)
	}

	if !errorContains(err, expectedErrMessage) {
		t.Errorf(`Unexpected error "%s" occurred. Expected error "%s"`, err, expectedErrMessage)
	}
}

func errorContains(err error, expectedErrStr string) bool {
	if err == nil {
		return expectedErrStr == ""
	}
	if expectedErrStr == "" {
		return false
	}
	return strings.Contains(err.Error(), expectedErrStr)
}
