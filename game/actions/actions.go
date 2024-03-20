package actions

type Action string

const (
	Retry Action = "r"
	Down  Action = "j"
	Left  Action = "h"
	Right Action = "l"
)
