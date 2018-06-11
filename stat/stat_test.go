package stat_test

import (
	. "github.com/bmuschko/link-verifier/stat"
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestSumSuccessesForEmptySlice(t *testing.T) {
	aggregateSummary := []Summary{}
	sum := SumSuccesses(aggregateSummary)
	Equal(t, 0, sum)
}

func TestSumSuccessesForPopulatedSlice(t *testing.T) {
	aggregateSummary := summaries()
	sum := SumSuccesses(aggregateSummary)
	Equal(t, 29, sum)
}

func TestSumFailuresForEmptySlice(t *testing.T) {
	aggregateSummary := []Summary{}
	sum := SumFailures(aggregateSummary)
	Equal(t, 0, sum)
}

func TestSumFailuresForPopulatedSlice(t *testing.T) {
	aggregateSummary := summaries()
	sum := SumFailures(aggregateSummary)
	Equal(t, 70, sum)
}

func summaries() []Summary {
	summaries := []Summary{}
	summaries = append(summaries, Summary{Successful: 20, Failed: 3})
	summaries = append(summaries, Summary{Successful: 7, Failed: 67})
	summaries = append(summaries, Summary{Successful: 2, Failed: 0})
	return summaries
}
