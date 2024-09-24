package internal

import "testing"

var container = NewContainer()
var article1 = Article{Title: "Advanced Golang", Content: "This article covers...", ID: 1}
var article2 = Article{Title: "Advanced C++", Content: "This article covers...", ID: 2}

func TestSearch(t *testing.T) {
	tests := []struct {
		name        string
		input       int
		expected    Article
		expectError bool
	}{
		{
			name:        "Simple input",
			input:       1,
			expected:    article1,
			expectError: false,
		},
		{
			name:        "Invalid input",
			input:       3,
			expected:    Article{},
			expectError: true,
		},
		{
			name:        "Empty container",
			input:       1,
			expected:    Article{},
			expectError: true,
		},
	}
	for _, test := range tests {
		container.AddArticle(article1)
		container.AddArticle(article2)
		t.Run(test.name, func(t *testing.T) {})
		if test.name == "Empty container" {
			emptyContainer := NewContainer()
			article, err := emptyContainer.Search(test.input)
			if article != test.expected {
				t.Errorf("Expected article %v, got %v", test.expected, article)
			}
			if test.expectError && err == nil {
				t.Errorf("Expected error, got none")
			}
			if !test.expectError && err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		} else {
			article, err := container.Search(test.input)
			if article != test.expected {
				t.Errorf("Expected article %v, got %v", test.expected, article)
			}
			if test.expectError && err == nil {
				t.Errorf("Expected error, got none")
			}
			if !test.expectError && err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		}
	}
}
