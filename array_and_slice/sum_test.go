package arrayandslice

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		got := Sum([]int{1, 2, 3, 4, 5})
		want := 15
		checkSumInt(t, got, want)
	})
}

func TestSumAll(t *testing.T) {

	t.Run("Test two slices", func(t *testing.T) {
		got := SumAll([]int{1, 3, 4}, []int{4, 5, 6})
		want := []int{8, 15}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got is '%d' but want '%d'", got, want)
		}
	})

	t.Run("Test one slice", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3, 4, 5, 6})
		want := []int{21}
		checkSumSlice(t, got, want)
	})
}

func TestSumAll2(t *testing.T) {
	got := SumAll2([]int{0, 2, 4}, []int{0, 0, 0})
	want := []int{6, 0}
	checkSumSlice(t, got, want)
}

func TestSumTails(t *testing.T) {

	t.Run("two slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 2})
		want := 7
		checkSumInt(t, got, want)
	})

	t.Run("one slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1})
		want := 5
		checkSumInt(t, got, want)
	})

	t.Run("empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3, 4})
		want := 9
		checkSumInt(t, got, want)
	})
}

func checkSumInt(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got is '%d', but want is '%d'", got, want)
	}
}

func checkSumSlice(t *testing.T, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got is '%v', but want is '%v'", got, want)
	}
}
