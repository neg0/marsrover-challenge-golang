package marsrover_test

import (
	"fmt"
	"marsrover/internal/marsrover"
)

func ExampleTestCase() {
	sut := marsrover.NewMarsRover()
	sut.SetPosition(5, 5, marsrover.N)
	sut.SetPosition(1, 2, marsrover.N)
	_ = sut.Process("LMLMLMLMM")

	fmt.Println(sut.String())

	sut.SetPosition(3, 3, marsrover.E)
	_ = sut.Process("MMRMMRMRRM")

	fmt.Println(sut.String())

	// Output:
	// 1 3 N
	// 5 1 E
}
