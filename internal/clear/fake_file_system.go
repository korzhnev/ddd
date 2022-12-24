package clear

const (
	copyOperation   = "COPY"
	moveOperation   = "MOVE"
	deleteOperation = "DELETE"
)

type FakeFileSystem struct {
	Operations [][]string
}

func NewFakeFileSystem() *FakeFileSystem {
	return &FakeFileSystem{}
}

func (f *FakeFileSystem) Copy(from string, to string) error {
	f.Operations = append(f.Operations, []string{copyOperation, from, to})
	return nil
}

func (f *FakeFileSystem) Move(from string, to string) error {
	f.Operations = append(f.Operations, []string{moveOperation, from, to})
	return nil
}

func (f *FakeFileSystem) Delete(path string) error {
	f.Operations = append(f.Operations, []string{deleteOperation, path})
	return nil
}
