package marsrover

type InvalidDirection struct{}

func (id InvalidDirection) Error() string {
	return "direction is invalid"
}

type InvalidCommand struct{}

func (ic InvalidCommand) Error() string {
	return "invalid command"
}
