package dirty

import (
	"ddd/internal/pkg/crypto"
	"io/ioutil"
	"os"
)

const filePerm = 0644

func Sync(source string, dest string) error {
	// Обойти исходную папку и создать словарь имен и их хешей

	sourceHashes := map[string]string{}
	files, err := ioutil.ReadDir(source)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		hash, err := crypto.HashFile(getPath(source, file.Name()))
		if err != nil {
			return err
		}
		sourceHashes[hash] = file.Name()
	}

	// Отслеживать файлы, найденные в целевой папке
	seen := map[string]struct{}{}

	// Обойти целевую папку и получить имена файлов и хеши

	files, err = ioutil.ReadDir(dest)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		destPath := getPath(dest, file.Name())

		hash, err := crypto.HashFile(destPath)
		if err != nil {
			return err
		}
		seen[hash] = struct{}{}

		// Если в целевой папке есть файл, которого нет, в источнике, то удалить его
		if _, ok := sourceHashes[hash]; !ok {
			os.Remove(destPath)
		} else if sourceHashes[hash] != file.Name() {
			// Если в целевой папке есть файл, который имеет другое имя, переименовать его
			os.Rename(destPath, getPath(dest, sourceHashes[hash]))
		}
	}

	// Каждый файл, который появляется в источнике, но не в месте назначения, скопировать в целевую папку
	for hash, file := range sourceHashes {
		if _, ok := seen[hash]; ok {
			continue
		}

		input, err := ioutil.ReadFile(getPath(source, file))
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(getPath(dest, file), input, filePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func getPath(dir string, file string) string {
	return dir + "/" + file
}
