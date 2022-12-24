package clear

import (
	"ddd/internal/pkg/crypto"
	"io/ioutil"
	"os"
)

const filePerm = 0644

func Sync(source string, dest string) error {
	sourceHashes, err := readPathsAndHashes(source)
	if err != nil {
		return err
	}
	destHashes, err := readPathsAndHashes(dest)
	if err != nil {
		return err
	}

	actions := determineActions(sourceHashes, destHashes, source, dest)
	for _, action := range actions {
		switch action.Type {
		case copyAction:
			input, err := ioutil.ReadFile(action.Paths[0])
			if err != nil {
				return err
			}
			if err = ioutil.WriteFile(action.Paths[1], input, filePerm); err != nil {
				return err
			}
		case moveAction:
			if err = os.Rename(action.Paths[0], action.Paths[1]); err != nil {
				return err
			}
		case deleteAction:
			if err = os.Remove(action.Paths[0]); err != nil {
				return err
			}
		}
	}
	return nil
}

func readPathsAndHashes(path string) (result map[string]string, err error) {
	result = map[string]string{}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		var hash string
		hash, err = crypto.HashFile(getPath(path, file.Name()))
		if err != nil {
			return
		}
		result[hash] = file.Name()
	}

	return
}
