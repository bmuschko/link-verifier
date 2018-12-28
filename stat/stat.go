package stat

// SumSuccesses sums up all successful verifications.
// Returns the overall number of successful verifications.
func SumSuccesses(aggregateSummary []Summary) int {
	successes := collect(aggregateSummary, func(summary Summary) int {
		return summary.Successful
	})

	return sum(successes)
}

// SumFailures sums up all failed verifications.
// Returns the overall number of failed verifications.
func SumFailures(aggregateSummary []Summary) int {
	failures := collect(aggregateSummary, func(summary Summary) int {
		return summary.Failed
	})

	return sum(failures)
}

// SumErrors sums up all errors.
// Returns the overall number of errors.
func SumErrors(aggregateSummary []Summary) int {
	errors := collect(aggregateSummary, func(summary Summary) int {
		return summary.Errored
	})

	return sum(errors)
}

func collect(list []Summary, f convert) []int {
	result := make([]int, len(list))

	for i, item := range list {
		result[i] = f(item)
	}

	return result
}

func sum(input []int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}

	return sum
}

type convert func(Summary) int

// Summary represents summary of a verification.
type Summary struct {
	Successful int
	Failed     int
	Errored    int
}
