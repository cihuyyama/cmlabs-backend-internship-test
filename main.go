package main

func main() {
	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			println("FizzBuzz")
			continue
		}

		if i%3 == 0 {
			println("Fizz")
			continue
		}

		if i%5 == 0 {
			println("Buzz")
			continue
		}

		println(i)
	}
}
