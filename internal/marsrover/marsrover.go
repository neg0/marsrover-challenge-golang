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
	mx        *sync.Mutex
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
		mr.mx.Lock()
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
		mr.mx.Unlock()
		mr.wg.Done()
	}()
	mr.wg.Wait()
	return processErr
}

func (mr *MarsRover) turnLeft() {
	mr.direction = util.TernaryInt(mr.direction-1 < N, W, mr.direction-1)
}

func (mr *MarsRover) turnRight() {
	mr.direction = util.TernaryInt(mr.direction+1 > W, N, mr.direction+1)
}

func (mr *MarsRover) step() error {
	switch mr.direction {
	case N:
		mr.y++
	case E:
		mr.x++
	case S:
		mr.y--
	case W:
		mr.x--
	default:
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
