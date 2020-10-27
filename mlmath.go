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
func KlabelBsort(a []Klabel) []Klabel {
	a_len := len(a)
	a_temp := make([]Klabel, a_len)

	copy(a_temp, a)
