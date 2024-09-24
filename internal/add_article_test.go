package internal

import (
	"reflect"
	"testing"
)

func TestAddArticle(t *testing.T) {
	tests := []struct {
		name        string
		input       []Article
		expected    map[int]Article
		expectError bool
	}{
		{
			name: "Simple input",
			input: []Article{
				{
					Title:   "Advanced Golang",
					Content: "This article covers...",
					ID:      1,
				},
				{
					Title:   "Healthy lifestyle",
					Content: "This article covers...",
					ID:      2,
				},
			},
			expected: map[int]Article{
				1: {
					Title:   "Advanced Golang",
					Content: "This article covers...",
					ID:      1,
				},
				2: {
					Title:   "Healthy lifestyle",
					Content: "This article covers...",
					ID:      2,
				},
			},
			expectError: false,
		},
		{
			name: "Invalid input",
			input: []Article{
				{
					Title:   "Advanced Golang",
					Content: "This article covers...",
					ID:      1,
				},
				{},
			},
			expected: map[int]Article{
				1: {
					Title:   "Advanced Golang",
					Content: "This article covers...",
					ID:      1,
				},
			},
			expectError: true,
		},
		{
			name:        "Empty input",
			input:       []Article{},
			expected:    map[int]Article{},
			expectError: false,
		},
		{
			name: "Duplicate ID",
			input: []Article{
				{
					Title:   "Advanced Golang",
					Content: "This article covers...",
					ID:      1,
				},
				{
					Title:   "Healthy lifestyle",
					Content: "This article covers...",
					ID:      1,
				},
			},
			expected: map[int]Article{
				1: {
					Title:   "Healthy lifestyle",
					Content: "This article covers...",
					ID:      1,
				},
			},
			expectError: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			container := NewContainer()
			for i, article := range test.input {
				err := container.AddArticle(article)
				if i == len(test.input)-1 {
					if test.expectError && err == nil {
						t.Error("Expected error, got none")
					} else if !test.expectError && err != nil {
						t.Errorf("Did not expect error but got one: %v", err)
					}
					if len(test.expected) != len(container.articles) {
						t.Errorf("Expected %d articles, got %d", len(test.expected), len(container.articles))
					}
					if !reflect.DeepEqual(test.expected, container.GetArticles()) {
						t.Errorf("Expected %v, got %v", test.expected, container)
					}
				}
			}
		})
	}
}
