package clear

const (
	copyAction   = "COPY"
	moveAction   = "MOVE"
	deleteAction = "DELETE"
)

type Action struct {
	Type  string
	Paths []string
}

func NewCopyAction(from string, to string) Action {
	return Action{Type: copyAction, Paths: []string{from, to}}
}

func NewMoveAction(from string, to string) Action {
	return Action{Type: moveAction, Paths: []string{from, to}}
}

func NewDeleteAction(path string) Action {
	return Action{Type: deleteAction, Paths: []string{path}}
}
