package main

//go:generate go run main.go

import (
	"os"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func main() {
	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top:    50,
				Left:   25,
				Right:  25,
				Bottom: 10,
			},
			FillColor: drawing.ColorFromHex("efefef"),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: chart.Seq[float64]{Sequence: chart.NewLinearSequence().WithStart(1.0).WithEnd(100.0)}.Values(),
				YValues: chart.Seq[float64]{Sequence: chart.NewRandomSequence().WithLen(100).WithMin(100).WithMax(512)}.Values(),
			},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
