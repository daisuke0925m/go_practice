package bisect_test

import (
	"testing"

	"github.com/DaisukeMatsumoto0925/go-test/bisect"
)

func TestStop_block(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	key := 4

	ans := bisect.BisectLeft(a, key)
	if ans != 3 {
		t.Errorf("get wrong ans: %d", ans)
	}
}

func TestOver_left(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	key := -1

	ans := bisect.BisectLeft(a, key)
	if ans != 0 {
		t.Errorf("get wrong ans: %d", ans)
	}
}

func TestOver_right(t *testing.T) {
	a := []int{1, 2, 3, 4, 4, 4, 5, 7, 8}
	key := 9

	ans := bisect.BisectLeft(a, key)
	if ans != 9 {
		t.Errorf("get wrong ans: %d", ans)
	}
}

func TestNot__exist(t *testing.T) {
	a := []int{1, 2, 3, 4, 4, 4, 5, 7, 8}
	key := 6

	ans := bisect.BisectLeft(a, key)
	if ans != 7 {
		t.Errorf("get wrong ans: %d", ans)
	}
}
