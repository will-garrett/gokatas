package gokatas

// PlusMinus - https://www.hackerrank.com/challenges/plus-minus/problem
func PlusMinus(arr []int32) map[string]float64 {
	total := len(arr)
	results := map[string]int{
		"pos": 0,
		"neg": 0,
		"nil": 0,
	}
	for i := 0; i < total; i++ {
		if arr[i] < 0 {
			results["neg"]++
		} else if arr[i] > 0 {
			results["pos"]++
		} else {
			results["nil"]++
		}
	}
	return map[string]float64{
		"positive": float64(results["pos"]) / float64(total),
		"negative": float64(results["neg"]) / float64(total),
		"zero":     float64(results["nil"]) / float64(total),
	}
}
