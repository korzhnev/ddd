package dirty

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestWhenAFileExistsInTheSourceButNotTheDestination(t *testing.T) {
	source, err := os.MkdirTemp("", "source")
	assert.NoError(t, err)

	dest, err := os.MkdirTemp("", "dest")
	assert.NoError(t, err)

	content := "I am really useful file"
	err = ioutil.WriteFile(getPath(source, "my-file"), []byte(content), filePerm)
	assert.NoError(t, err)

	Sync(source, dest)

	expectedPath := getPath(dest, "/my-file")
	assert.FileExists(t, expectedPath)
	fileContent, err := ioutil.ReadFile(expectedPath)
	assert.NoError(t, err)
	assert.Equal(t, content, string(fileContent))
}

func TestWhenAFileHasBeenRenamedInTheSource(t *testing.T) {
	source, err := os.MkdirTemp("", "source")
	assert.NoError(t, err)

	dest, err := os.MkdirTemp("", "dest")
	assert.NoError(t, err)

	content := "I am a file, which has been renamed"

	sourcePath := getPath(source, "source-filename")
	oldDestPath := getPath(dest, "dest-filename")

	expectedDestPath := getPath(dest, "source-filename")

	err = ioutil.WriteFile(sourcePath, []byte(content), filePerm)
	assert.NoError(t, err)

	err = ioutil.WriteFile(oldDestPath, []byte(content), filePerm)
	assert.NoError(t, err)

	Sync(source, dest)

	assert.NoFileExists(t, oldDestPath)

	fileContent, err := ioutil.ReadFile(expectedDestPath)
	assert.NoError(t, err)
	assert.Equal(t, content, string(fileContent))
}
