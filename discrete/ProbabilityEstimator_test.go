package discrete_test

import (
	"math"
	"testing"

	"github.com/kzahedi/goent/discrete"
)

func TestEmpirical1D(t *testing.T) {
	t.Log("Testing Empirical1D")
	d := []int{0, 0, 1, 1, 2, 2, 3, 3}
	p := discrete.Empirical1D(d)

	if len(p) != 4 {
		t.Errorf("Empirical1D should return a slice of length 4 and not %d", len(p))
	}

	for i := range p {
		if math.Abs(p[i]-1.0/4.0) > 0.0000001 {
			t.Errorf("p[%d] should be 1/4 and not %f", i, p[i])
		}
	}
}

func TestEmpirical2D(t *testing.T) {
	t.Log("Testing Empirical2D")

	d := [][]int{
		{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4},
		{1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4},
		{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4},
		{3, 0}, {3, 1}, {3, 2}, {3, 3}, {3, 4}}

	p := discrete.Empirical2D(d)

	if len(p) != 4 {
		t.Errorf("Empirical2D number of rows should be 4 but it is %d", len(p))
	}
	if len(p[0]) != 5 {
		t.Errorf("Empirical2D number of columns should be 5 but it is %d", len(p[0]))
	}

	for r := 0; r < 4; r++ {
		for c := 0; c < 5; c++ {
			if math.Abs(p[r][c]-1.0/20.0) > 0.0000001 {
				t.Errorf("p[%d][%d] should be 1/20 and not %f", r, c, p[r][c])
			}
		}
	}
}

func TestEmpirical3D(t *testing.T) {
	t.Log("Testing Empirical2D")

	d := [][]int{
		{0, 0, 0}, {0, 0, 1}, {0, 0, 2}, {0, 0, 3},
		{0, 1, 0}, {0, 1, 1}, {0, 1, 2}, {0, 1, 3},
		{0, 2, 0}, {0, 2, 1}, {0, 2, 2}, {0, 2, 3},
		{1, 0, 0}, {1, 0, 1}, {1, 0, 2}, {1, 0, 3},
		{1, 1, 0}, {1, 1, 1}, {1, 1, 2}, {1, 1, 3},
		{1, 2, 0}, {1, 2, 1}, {1, 2, 2}, {1, 2, 3}}

	p := discrete.Empirical3D(d)

	if len(p) != 2 {
		t.Errorf("Empirical3D 1st dimension should be 2 but it is %d", len(p))
	}
	if len(p[0]) != 3 {
		t.Errorf("Empirical3D 2nd dimension should be 3 but it is %d", len(p[0]))
	}
	if len(p[0][0]) != 4 {
		t.Errorf("Empirical3D 3rd dimension should be 4 but it is %d", len(p[0][0]))
	}

	for a := 0; a < 2; a++ {
		for b := 0; b < 3; b++ {
			for c := 0; c < 4; c++ {
				if math.Abs(p[a][b][c]-1.0/24.0) > 0.0000001 {
					t.Errorf("p[%d][%d][%d] should be 1/24 and not %f", a, b, c, p[a][b][c])
				}
			}
		}
	}
}

func TestEmpirical4D(t *testing.T) {
	t.Log("Testing Empirical4D")

	d := make([][]int, 2*3*4*2, 2*3*4*2)
	for i := 0; i < 2*3*4*2; i++ {
		d[i] = make([]int, 4, 4)
	}

	index := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 4; k++ {
				for l := 0; l < 2; l++ {
					d[index][0] = i
					d[index][1] = j
					d[index][2] = k
					d[index][3] = l
					index++
				}
			}
		}
	}

	p := discrete.Empirical4D(d)

	if len(p) != 2 {
		t.Errorf("Empirical4D 1st dimension should be 2 but it is %d", len(p))
	}
	if len(p[0]) != 3 {
		t.Errorf("Empirical4D 2nd dimension should be 3 but it is %d", len(p[0]))
	}
	if len(p[0][0]) != 4 {
		t.Errorf("Empirical4D 3rd dimension should be 4 but it is %d", len(p[0][0]))
	}
	if len(p[0][0][0]) != 2 {
		t.Errorf("Empirical4D 4th dimension should be 4 but it is %d", len(p[0][0][0]))
	}

	for a := 0; a < 2; a++ {
		for b := 0; b < 3; b++ {
			for c := 0; c < 4; c++ {
				for d := 0; d < 2; d++ {
					if math.Abs(p[a][b][c][d]-1.0/(2.0*3.0*4.0*2.0)) > 0.0000001 {
						t.Errorf("p[%d][%d][%d][%d] should be 1/%d and not %f", a, b, c, d, 2*3*4*2, p[a][b][c][d])
					}
				}
			}
		}
	}
}
