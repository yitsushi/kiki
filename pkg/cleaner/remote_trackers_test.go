package cleaner_test

import (
	"testing"

	"gitea.code-infection.com/efertone/kiki/pkg/cleaner"
)

type TestCase struct {
	Input    string
	Expected string
}

func TestRemoveTrackers(t *testing.T) {
	suit := []TestCase{
		{"https://example.com/", "https://example.com/"},
		{"https://example.com/with/path", "https://example.com/with/path"},
		{"https://example.com/?query=asd", "https://example.com/?query=asd"},
		{"https://example.com/?utm_source=evil_source", "https://example.com/"},
		{"https://example.com/?utm_medium=evil_medium", "https://example.com/"},
		{"https://example.com/?utm_campaign=evil_campaign", "https://example.com/"},
		{
			"https://example.com/?query=asd&utm_campaign=evil_campaign",
			"https://example.com/?query=asd"},
		{
			"https://example.com/?utm_campaign=evil_campaign&query=asd",
			"https://example.com/?query=asd"},
		{
			"https://example.com/?query=asd&utm_campaign=evil_campaign&another=one",
			"https://example.com/?query=asd&another=one"},
	}

	for _, testCase := range suit {
		output := cleaner.RemoveTrackers(testCase.Input)
		if output != testCase.Expected {
			t.Errorf("RemoveTrackers(%s) = %s; want %s", testCase.Input, output, testCase.Expected)
		}
	}
}
