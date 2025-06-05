package types

type Command struct {
	Name        string
	Description string
	Execute     func(args []string) error
}
