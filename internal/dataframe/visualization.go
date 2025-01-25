package dataframe

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func (df *DataFrame) Histogram(column string) error {
	series := df.DF.Col(column)

	// Convert to float64 slice
	values := series.Float()

	// Filter out NaN values
	var filteredValues []float64
	for _, val := range values {
		if !isNaN(val) {
			filteredValues = append(filteredValues, val)
		}
	}

	if len(filteredValues) == 0 {
		return fmt.Errorf("no valid data in column '%s' to plot", column)
	}

	// Create a histogram
	p := plot.New()
	p.Title.Text = fmt.Sprintf("Histogram of %s", column)
	p.X.Label.Text = column
	p.Y.Label.Text = "Frequency"

	// Create a histogram using gonum/plot
	h, err := plotter.NewHist(plotter.Values(filteredValues), 16)
	if err != nil {
		return fmt.Errorf("failed to create histogram: %v", err)
	}
	h.FillColor = color.RGBA{R: 255, A: 255}

	p.Add(h)

	// Save the plot to a PNG file
	fileName := fmt.Sprintf("histogram_%s.png", column)
	if err := p.Save(6*vg.Inch, 4*vg.Inch, fileName); err != nil {
		return fmt.Errorf("failed to save histogram: %v", err)
	}

	fmt.Printf("Histogram saved as %s\n", fileName)
	return nil
}

// isNaN checks if a float64 value is NaN.
func isNaN(f float64) bool {
	return f != f
}
