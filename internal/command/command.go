package command

type Command string

const (
	CMDGet    Command = "GET"
	CMDSet    Command = "SET"
	CMDDelete Command = "DELETE"
)

func (c Command) isValid() bool {
	validCommands := []Command{CMDGet, CMDSet, CMDDelete}
	for i := range validCommands {
		if c == validCommands[i] {
			return true
		}
	}
	return false
}
