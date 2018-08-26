package pipelines

import "testing"

func TestLaunchPipeline(t *testing.T) {
	// [list size][the result]
	tableTest := [][]int{
		{3, 14},
		{5, 55},
	}

	// ...
	var res int
	for _, test := range tableTest {
		// test[0] will retrieve a number n which is list size
		// then will create a list with [1,2,3.....n]
		res = LaunchPipeline(test[0])
		if res != test[1] {
			t.Fatal()
		}

		t.Logf("%d == %d\n", res, test[1])
	}
}
