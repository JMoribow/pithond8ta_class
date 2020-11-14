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

	for i, _ := range a {
		if i+1 != a_len {
			if a[i+1].Values[0] < a[i].V