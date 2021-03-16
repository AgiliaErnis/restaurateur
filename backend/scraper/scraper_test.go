package scraper

import "testing"

func TestSliceContains(t *testing.T) {
	type sliceContainsTest struct {
		inputSlice []string
		inputVal   string
		expected   bool
	}

	sliceContainsTests := []sliceContainsTest{
		sliceContainsTest{[]string{"test1", "test2", "test3"}, "test2", true},
		sliceContainsTest{[]string{}, "test2", false},
		sliceContainsTest{[]string{}, "", false},
		sliceContainsTest{[]string{"test1", "test2", "test3"}, "", false},
		sliceContainsTest{[]string{"test1", "test3"}, "test2", false},
		sliceContainsTest{[]string{"val"}, "val", true},
		sliceContainsTest{[]string{"val"}, "val1", false},
	}
	for _, test := range sliceContainsTests {
		if actual := sliceContains(test.inputSlice, test.inputVal); actual != test.expected {
			t.Errorf("expected %t for slice %v and value %s, got %t",
				test.expected, test.inputSlice, test.inputVal, actual)
		}
	}
}
