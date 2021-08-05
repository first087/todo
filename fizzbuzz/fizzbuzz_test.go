package fizzbuzz

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	cases := map[int]string{
		1:  "1",
		2:  "2",
		3:  "Fizz",
		6:  "Fizz",
		9:  "Fizz",
		5:  "Buzz",
		10: "Buzz",
		15: "FizzBuzz",
	}

	for given, want := range cases {
		t.Run(fmt.Sprintf("given %d want %q", given, want), func(t *testing.T) {
			get := Say(given)
			if want != get {
				t.Errorf("%q %q", want, get)
			}
		})
	}
}

type stubIntn struct {
	val int
}

func (s stubIntn) Intn(int) int {
	return s.val
}

func TestRandomFizzBuzz(t *testing.T) {
	want := "Fizz"
	get := RandomFizzBuzz(stubIntn{val: 2})

	if want != get {
		t.Error()
	}
}

type IntnFunc func(int) int

func (f IntnFunc) Intn(n int) int {
	return f(n)
}

func TestRandomFizzBuzz2(t *testing.T) {
	want := "Fizz"
	get := RandomFizzBuzz(IntnFunc(func(i int) int { return 2 }))

	if want != get {
		t.Error()
	}
}
