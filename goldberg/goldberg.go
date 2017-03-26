package goldberg

import (
	"github.com/MichaelMauderer/polyhedra"
	"errors"
)

func New(n int, m int) (polyhedra.Interface, error) {
	if n != m {
		return nil, errors.New("Class III not supported.")
	}
	if m != 0 {
		return nil, errors.New("Class II not supported.")
	}
}
