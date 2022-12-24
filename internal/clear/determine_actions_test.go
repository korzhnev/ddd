package clear

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenAFileExistsInTheSourceButNotTheDestination(t *testing.T) {
	source := "/source"
	dest := "/dest"
	filename := "my-file"

	sourceHashes := map[string]string{
		"hash1": filename,
	}
	destHashes := map[string]string{}

	actions := determineActions(sourceHashes, destHashes, source, dest)
	expectedActions := []Action{{
		Type:  copyAction,
		Paths: []string{getPath(source, filename), getPath(dest, filename)},
	}}
	assert.Equal(t, expectedActions, actions)
}

func TestWhenAFileHasBeenRenamedInTheSource(t *testing.T) {
	source := "/source"
	dest := "/dest"

	sourceHashes := map[string]string{
		"hash1": "renamed-file",
	}
	destHashes := map[string]string{
		"hash1": "original-file",
	}

	actions := determineActions(sourceHashes, destHashes, source, dest)
	expectedActions := []Action{{
		Type:  moveAction,
		Paths: []string{getPath(dest, "original-file"), getPath(dest, "renamed-file")},
	}}

	assert.Equal(t, expectedActions, actions)
}
