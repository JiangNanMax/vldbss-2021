package main

// URLTop10 .
func URLTop10(nWorkers int) RoundsArgs {
	// YOUR CODE HERE :)
	// And don't forget to document your idea.
	var args RoundsArgs
	// round 1: do url count
	args = append(args, RoundArgs{
		MapFunc:    ExampleURLCountMap,
		ReduceFunc: ExampleURLCountReduce,
		NReduce:    nWorkers,
	})
	// round 2: sort and get the 10 most frequent URLs
	args = append(args, RoundArgs{
		MapFunc:    ExampleURLTop10Map,
		ReduceFunc: ExampleURLTop10Reduce,
		NReduce:    1,
	})
	return args
}
