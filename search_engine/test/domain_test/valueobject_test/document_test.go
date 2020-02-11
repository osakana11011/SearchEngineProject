package valueobject_test

import (
	"testing"
	"strconv"

	"search_engine_project/search_engine/src/domain/valueobject"
)

func TestDocument(t *testing.T) {
	testPatterns := []struct {
		givenTitle string
		givenURL string
		expectedTitle string
		expectedURL string
	} {
		{"Google", "https://ja.wikipedia.org/wiki/Google", "Google", "https://ja.wikipedia.org/wiki/Google"},
		{"メタ構文変数", "https://ja.wikipedia.org/wiki/%E3%83%A1%E3%82%BF%E6%A7%8B%E6%96%87%E5%A4%89%E6%95%B0", "メタ構文変数", "https://ja.wikipedia.org/wiki/%E3%83%A1%E3%82%BF%E6%A7%8B%E6%96%87%E5%A4%89%E6%95%B0"},
	}

	for i, data := range testPatterns {
		t.Run("test"+strconv.Itoa(i), func(t *testing.T) {
			res := valueobject.NewDocument(data.givenTitle, data.givenURL)

			// Titleの一致検証
			if res.Title() != data.expectedTitle {
				t.Fatalf("[Title不一致] expected: %s, actual: %s", data.expectedTitle, res.Title())
			}

			// URLの一致検証
			if res.URL() != data.expectedURL {
				t.Fatalf("[URL不一致] expected: %s, actual: %s", data.expectedURL, res.URL())
			}
		})
	}
}
