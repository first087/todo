package fizzbuzz

import (
	"strconv"
)

func Say(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(n)
}

// 1. Parameter
// 2. Member

type Intner interface {
	Intn(n int) int
}

func RandomFizzBuzz(random Intner) string {
	// s := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(s)

	n := random.Intn(100) + 1
	return Say(n)
}

type RandomFizzBuzzHandler struct {
	random Intner
}

func (r RandomFizzBuzzHandler) Handler() string {
	n := r.random.Intn(100) + 1

	return Say(n)
}
