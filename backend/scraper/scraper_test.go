package scraper

import "testing"

func TestSliceContains(t *testing.T) {
    inputSlice := []string{"test1", "test2", "test3"}
    inputVal := "test2"
    expected := true
    actual := sliceContains(inputSlice, inputVal)
    if actual != expected {
        t.Errorf("expected %t for slice %v and value %s, got %t",
        expected, inputSlice, inputVal, actual)
    }
}
