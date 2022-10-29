package sort_test

import "fmt"

func check(label string, result []int, expect []int) error {
	for i, e := range expect {
		if result[i] != e {
			return fmt.Errorf("%s Error, index %d, result: %d / expect: %d", label, i, result[i], e)
		}
	}
	return nil
}
