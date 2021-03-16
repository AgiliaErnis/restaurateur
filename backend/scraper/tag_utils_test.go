package scraper

import "testing"

func TestGetCuisines(t *testing.T) {
	type getCuisinesTest struct {
		inputSlice []string
		expected   []string
	}
	getCuisinesTests := []getCuisinesTest{
		getCuisinesTest{
			[]string{"Americká", "Italská", "test", "Thajská"},
			[]string{"American", "Italian", "Thai"},
		},
		getCuisinesTest{
			[]string{"", "", "test", ""},
			[]string{},
		},
		getCuisinesTest{
			[]string{},
			[]string{},
		},
		getCuisinesTest{
			[]string{"Indická"},
			[]string{"Indian"},
		},
	}
	for _, test := range getCuisinesTests {
		actual := getCuisines(test.inputSlice)
		for i := range actual {
			if actual[i] != test.expected[i] {
				t.Errorf("expected %v for input slice %v, got %v",
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
		getPriceRangeTest{
			[]string{"tag1", "tag2", " ", "Do 300 Kč", "tag3"},
			"0 - 300 Kč",
		},
		getPriceRangeTest{
			[]string{},
			"Not available",
		},
		getPriceRangeTest{
			[]string{"300 - 600 Kč", ""},
			"300 - 600 Kč",
		},
		getPriceRangeTest{
			[]string{"tag", "Nad 600 Kč"},
			"600+ Kč",
		},
		getPriceRangeTest{
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
