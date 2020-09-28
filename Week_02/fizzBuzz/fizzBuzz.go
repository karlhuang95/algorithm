package fizzBuzz

import "strconv"

func fizzBuzz(n int) []string {
	res := make([]string, n+1)
	for i := 1; i <= n; i++ {
		if res[i] == "" {
			res[i] = strconv.Itoa(i)
		}
		a := i * 3
		if a <= n && res[a] == "" {
			res[a] = "Fizz" // 将*3的填好
		}
		a = i * 5
		if a <= n && res[a] == "" {
			res[a] = "Buzz" // 将*5的填好
		}
		a = i * 3 * 5
		if a <= n && res[a] == "" {
			res[a] = "FizzBuzz" // 将*3*5的填好
		}
	}
	return res[1:] // 舍弃0号元素

}

