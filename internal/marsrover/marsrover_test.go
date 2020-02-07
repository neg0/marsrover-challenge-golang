package marsrover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarsRover(t *testing.T) {
	sut := NewMarsRover()
	var err error

	sut.SetPosition(5, 5, N)
	sut.SetPosition(1, 2, N)
	err = sut.Process("LMLMLMLMM")

	assert.NoError(t, err)
	assert.Equal(t, "1 3 N", sut.String())

	sut.SetPosition(3, 3, E)
	err = sut.Process("MMRMMRMRRM")

	assert.NoError(t, err)
	assert.Equal(t, "5 1 E", sut.String())

}

func TestInvalidCommand_Error(t *testing.T) {
	sut := NewMarsRover()

	sut.SetPosition(5, 5, N)
	err := sut.Process("ZXJKTY")

	assert.Error(t, err)
	assert.IsType(t, InvalidCommand{}, err)
}
