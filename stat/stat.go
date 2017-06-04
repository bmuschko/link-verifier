package stat

func SumSuccesses(aggregateSummary []Summary) int {
    successes := collect(aggregateSummary, func(summary Summary) int {
        return summary.Successful
    })

    return sum(successes)
}

func SumFailures(aggregateSummary []Summary) int {
    failures := collect(aggregateSummary, func(summary Summary) int {
        return summary.Failed
    })

    return sum(failures)
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

type Summary struct {
    Successful int
    Failed int
}
