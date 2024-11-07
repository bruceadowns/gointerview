package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcQuad(t *testing.T) {
	tt := []struct {
		a, b, c float64
		x, y    float64
	}{
		{1.0, -5.0, -14.0, 7.0, -2.0},
	}

	for _, test := range tt {
		x, y := calcQuad(test.a, test.b, test.c)
		assert.Equal(t, test.x, x)
		assert.Equal(t, test.y, y)
	}
}
