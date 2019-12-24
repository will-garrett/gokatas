package gokatas

import "testing"

func TestPlusMinus1(t *testing.T) {
	case1 := []int32{-4, 3, -9, 0, 4, 1}
	result := PlusMinus(case1)
	if result["positive"] != .5 {
		t.Errorf("Expected .5 got %v", result["positive"])
	}

	if result["negative"] != float64(1)/float64(3) {
		t.Errorf("Expected .33333 got %v", result["negative"])
	}

	if result["zero"] != float64(1)/float64(6) {
		t.Errorf("Expected .166667 got %v", result["zero"])
	}
}
func TestPlusMinus2(t *testing.T) {
	case2 := []int32{1, 2, 3, -1, -2, -3, 0, 0}
	result := PlusMinus(case2)
	if result["positive"] != float64(3)/float64(8) {
		t.Errorf("Expected %v got %v", float64(3)/float64(8), result["positive"])
	}
	if result["negative"] != float64(3)/float64(8) {
		t.Errorf("Expected %v got %v", float64(3)/float64(8), result["negative"])
	}
	if result["zero"] != float64(2)/float64(8) {
		t.Errorf("Expected %v got %v", float64(2)/float64(8), result["zero"])
	}

}
