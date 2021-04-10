package scraper

import "testing"

func TestSliceContains(t *testing.T) {
	type sliceContainsTest struct {
		inputSlice []string
		inputVal   string
		expected   bool
	}

	sliceContainsTests := []sliceContainsTest{
		{[]string{"test1", "test2", "test3"}, "test2", true},
		{[]string{}, "test2", false},
		{[]string{}, "", false},
		{[]string{"test1", "test2", "test3"}, "", false},
		{[]string{"test1", "test3"}, "test2", false},
		{[]string{"val"}, "val", true},
		{[]string{"val"}, "val1", false},
	}
	for _, test := range sliceContainsTests {
		if actual := sliceContains(test.inputSlice, test.inputVal); actual != test.expected {
			t.Errorf("expected `%t` for slice %q and value %q, got `%t`",
				test.expected, test.inputSlice, test.inputVal, actual)
		}
	}
}
