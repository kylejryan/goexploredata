package dataframe

import (
	"bytes"
	"fmt"
	"math"
)

func (df *DataFrame) Describe() (string, error) {
	summary := new(bytes.Buffer)
	summary.WriteString("Summary Statistics:\n")
	summary.WriteString(fmt.Sprintf("%-15s %-10s %-10s %-10s %-10s\n", "Column", "Mean", "Median", "StdDev", "Count"))

	for _, name := range df.DF.Names() {
		// Convert column to float64 slice
		values := df.DF.Col(name).Float()

		if len(values) == 0 {
			continue
		}

		mean := calculateMean(values)
		median := calculateMedian(values)
		stdDev := calculateStdDev(values)

		summary.WriteString(fmt.Sprintf("%-15s %-10.4f %-10.4f %-10.4f %-10d\n",
			name, mean, median, stdDev, len(values)))
	}

	return summary.String(), nil
}

func calculateMean(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func calculateMedian(values []float64) float64 {
	sorted := make([]float64, len(values))
	copy(sorted, values)

	// Simple bubble sort
	for i := 0; i < len(sorted); i++ {
		for j := 0; j < len(sorted)-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2
	}
	return sorted[mid]
}

func calculateStdDev(values []float64) float64 {
	mean := calculateMean(values)
	variance := 0.0
	for _, v := range values {
		variance += math.Pow(v-mean, 2)
	}
	variance /= float64(len(values))
	return math.Sqrt(variance)
}
