package mlmath

import (
	"errors"
	"math"
)

type Klabel struct {
	Values []float64
	Label  string
}

// Bubblesort for Klabel and repetitive sorting for small k
func KlabelBsort(a []Klabel) []Klabe