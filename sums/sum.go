package sums

func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

func SumAllTails(numbersToSum ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(numbersToSum, sumTail, []int{})
}

func Reduce[A any](collection []A, f func(A, A) A, initialValue A) A {
	result := initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}

func BalanceFor(transactions []Transaction, name string) float64 {
	var balance float64
	for _, t := range transactions {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}

	return balance
}

type Transaction struct {
	From, To string
	Sum      float64
}
