package colly

import (
	"slices"
	"testing"
)

func TestTagCollector(t *testing.T) {
	expectedTags := []string{}

	collyTagCollector := CollyTagCollector{}

	tags, err := collyTagCollector.All()

	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	if !slices.Equal(tags, expectedTags) {
		t.Errorf("got %v want %v", tags, expectedTags)
	}
}
