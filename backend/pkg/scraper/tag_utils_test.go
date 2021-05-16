package scraper

import "testing"
import "fmt"

func TestGetCuisines(t *testing.T) {
	type getCuisinesTest struct {
		inputSlice []string
		expected   []string
	}
	getCuisinesTests := []getCuisinesTest{
		{
			[]string{"Americká", "Italská", "test", "Thajská"},
			[]string{"American", "Italian", "Thai"},
		},
		{
			[]string{"", "", "test", ""},
			[]string{},
		},
		{
			[]string{},
			[]string{},
		},
		{
			[]string{"Indická"},
			[]string{"Indian"},
		},
	}
	for _, test := range getCuisinesTests {
		actual := getCuisines(test.inputSlice)
		fmt.Println(actual)
		if len(actual) != len(test.expected) {
			t.Errorf("Length of actual `%q` not equal to expected `%q`", actual, test.expected)
			t.FailNow()
		}
		for i := range actual {
			if actual[i] != test.expected[i] {
				t.Errorf("expected %q for input slice %q, got %q",
					test.expected, test.inputSlice, actual)
				t.FailNow()
			}
		}
	}
}

func TestGetPriceRange(t *testing.T) {
	type getPriceRangeTest struct {
		inputSlice []string
		expected   string
	}
	getPriceRangeTests := []getPriceRangeTest{
		{
			[]string{"tag1", "tag2", " ", "Do 300 Kč", "tag3"},
			"0 - 300 Kč",
		},
		{
			[]string{},
			"Not available",
		},
		{
			[]string{"300 - 600 Kč", ""},
			"300 - 600 Kč",
		},
		{
			[]string{"tag", "Nad 600 Kč"},
			"600+ Kč",
		},
		{
			[]string{"tag", "tag2 ", "300 - 600"},
			"Not available",
		},
	}
	for _, test := range getPriceRangeTests {
		if actual := getPriceRange(test.inputSlice); actual != test.expected {
			t.Errorf("expected %s for slice %v, got %s",
				test.expected, test.inputSlice, actual)
		}
	}
}
