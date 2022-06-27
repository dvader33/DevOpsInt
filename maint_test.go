package main

import "testing"

func TestConcat(t *testing.T) {
	if concat("nombre") != "Hello nombre your message will be send" {
		t.Error("Not passing")
	}
}

func TestTableConcat(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"Daniel", "Hello Daniel your message will be send"},
		{"Jose", "Hello Jose your message will be send"},
		{"Ada", "Hello Ada your message will be send"},
		{"Fernando", "Hello Fernando your message will be send"},
		{"Lucia", "Hello Lucia your message will be send"},
	}

	for _, test := range tests {
		if output := concat(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
