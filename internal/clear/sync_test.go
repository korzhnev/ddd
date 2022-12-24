package clear

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenAFileExistsInTheSourceButNotTheDestination(t *testing.T) {
	source := "/source"
	dest := "/dest"
	filename := "my-file"
	filesystem := NewFakeFileSystem()

	reader := func(path string) (map[string]string, error) {
		state := map[string]map[string]string{
			source: {
				"hash1": filename,
			},
			dest: {},
		}
		return state[path], nil
	}

	sync(reader, filesystem, source, dest)

	assert.Equal(t, [][]string{{copyOperation, getPath(source, filename), getPath(dest, filename)}}, filesystem.Operations)
}

func TestWhenAFileHasBeenRenamedInTheSource(t *testing.T) {
	source := "/source"
	dest := "/dest"
	filesystem := NewFakeFileSystem()

	reader := func(path string) (map[string]string, error) {
		state := map[string]map[string]string{
			source: {
				"hash1": "renamed-file",
			},
			dest: {
				"hash1": "original-file",
			},
		}
		return state[path], nil
	}

	sync(reader, filesystem, source, dest)

	assert.Equal(
		t,
		[][]string{{moveOperation, getPath(dest, "original-file"), getPath(dest, "renamed-file")}},
		filesystem.Operations,
	)
}
