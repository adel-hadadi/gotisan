package commands

type Command interface {
	Make([]string)
	Help([]string)
}
