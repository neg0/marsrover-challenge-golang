package marsrover

import (
	"fmt"
	"strings"
	"sync"

	"marsrover/pkg/util"
)

type MarsRover struct {
	x         int
	y         int
	direction int
	wg        *sync.WaitGroup
	mx 	  	  *sync.Mutex
}

func NewMarsRover() MarsRover {
	return MarsRover{
		wg: &sync.WaitGroup{},
		mx: &sync.Mutex{},
	}
}

func (mr *MarsRover) SetPosition(x int, y int, direction int) {
	mr.x = x
	mr.y = y
	mr.direction = direction
}

func (mr *MarsRover) Process(commands string) error {
	var processErr error

	mr.wg.Add(1)
	go func() {
		commandsCollection := strings.Split(commands, "")
		for _, command := range commandsCollection {
			switch command {
			case Left:
				mr.turnLeft()
			case Right:
				mr.turnRight()
			case Move:
				err := mr.step()
				if err != nil {
					processErr = err
					return
				}
			default:
				processErr = InvalidCommand{}
			}
		}
		mr.wg.Done()
	}()
	mr.wg.Wait()
	return processErr
}

func (mr *MarsRover) turnLeft() {
	mr.mx.Lock()
	currentFacing := mr.direction
	mr.direction = util.TernaryInt(currentFacing-1 < N, W, currentFacing-1)
	mr.mx.Unlock()
}

func (mr *MarsRover) turnRight() {
	mr.mx.Lock()
	currentFacing := mr.direction
	mr.direction = util.TernaryInt(currentFacing+1 > W, N, currentFacing+1)
	mr.mx.Unlock()
}

func (mr *MarsRover) step() error {
	mr.mx.Lock()
	switch mr.direction {
	case N:
		mr.y++
		mr.mx.Unlock()
	case E:
		mr.x++
		mr.mx.Unlock()
	case S:
		mr.y--
		mr.mx.Unlock()
	case W:
		mr.x--
		mr.mx.Unlock()
	default:
		mr.mx.Unlock()
		return InvalidDirection{}
	}
	return nil
}

func (mr MarsRover) directionIndicator() string {
	var dir string
	switch mr.direction {
	case N:
		dir = "N"
	case E:
		dir = "E"
	case S:
		dir = "S"
	case W:
		dir = "W"
	default:
		dir = "N"
	}
	return dir
}

func (mr MarsRover) String() string {
	return fmt.Sprintf(
		"%d %d %s",
		mr.x,
		mr.y,
		mr.directionIndicator(),
	)
}
