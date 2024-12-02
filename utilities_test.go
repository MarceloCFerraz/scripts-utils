package utils

import (
	"testing"
)

func TestBatching(t *testing.T) {
	arr := []int{0, 1, 2, 3}
	
	batches, err := SplitInBatches(&arr, 3)

	expected := [][]int{
		{0, 1, 2},
		{3},
	}
	
	if err != nil {
		t.Error(err)
		return
	}

	if len(*batches) != len(expected) {
		t.Errorf("expected %d batches, but got %d", len(expected), len(*batches))
		return
	}

	for i := range *batches {
		for j := range (*batches)[i] {
			if (*batches)[i][j] != expected[i][j] {
				t.Errorf("expected %d, but got %d", expected, *batches)
				return
			}
		}
	}
}
