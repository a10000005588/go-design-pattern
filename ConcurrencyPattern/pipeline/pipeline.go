package pipelines

// return a channel that only do output.
func generator(max int) <-chan int {
	outChInt := make(chan int, 100)

	// createa goroutine that generate numbers into the list via channel.
	go func() {
		for i := 1; i <= max; i++ {
			outChInt <- i
		}

		close(outChInt)
	}()

	return outChInt
}

// power will take the channel which generate from generator()
// then return a chennl which only do output.
func power(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()

	return out
}

// sum take a channel which accept input channel
// then return a output channel
func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		var sum int

		for v := range in {
			sum += v
		}

		out <- sum
		close(out)
	}()

	return out
}

// Finally, we can implement the LaunchPipeline function
// which is return the result (int)
func LaunchPipeline(amount int) int {
	firstCh := generator(amount)
	secondCh := power(firstCh)
	thirdCh := sum(secondCh)

	result := <-thirdCh

	return result
}
