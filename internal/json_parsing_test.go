package internal

import (
	"testing"
)

func TestParseJSON(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		expected    Article
		expectError bool
	}{
		{
			name:  "Simple JSON-input",
			input: []byte(`{"title":"Advanced Golang","content":"This article covers...","id":1}`),
			expected: Article{
				Title:   "Advanced Golang",
				Content: "This article covers...",
				ID:      1,
			},
			expectError: false,
		},
		{
			name:        "Empty JSON-input",
			input:       []byte(`{}`),
			expected:    Article{},
			expectError: true,
		},
		{
			name:        "Empty input",
			input:       nil,
			expected:    Article{},
			expectError: true,
		},
		{
			name:        "Invalid JSON-title",
			input:       []byte(`{"title":"","content":"This article covers...","id":1}`),
			expected:    Article{},
			expectError: true,
		},
		{
			name:        "Invalid JSON-content",
			input:       []byte(`{"title":"Advanced Golang","content":"","id":1}`),
			expected:    Article{},
			expectError: true,
		},
		{
			name:        "Invalid JSON-id",
			input:       []byte(`{"title":"Advanced Golang","content":"This article covers...","id":null}`),
			expected:    Article{},
			expectError: true,
		},
		{
			name:        "Invalid JSON-syntax",
			input:       []byte(`{"title":"Advanced Golang","content":"This article covers...","id":null`),
			expected:    Article{},
			expectError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := test.input
			actual, err := ParseJSON(input)
			if err != nil && !test.expectError {
				t.Errorf("Did not expect error but got one: %v", err)
			} else if err == nil && test.expectError {
				t.Errorf("Expected error but got none")
			}
			if actual.Title != test.expected.Title {
				t.Errorf("Expected title %v but got %v", test.expected.Title, actual.Title)
			}
			if actual.Content != test.expected.Content {
				t.Errorf("Expected content %v but got %v", test.expected.Content, actual.Content)
			}
			if actual.ID != test.expected.ID {
				t.Errorf("Expected id %v but got %v", test.expected.ID, actual.ID)
			}
		})
	}
}
