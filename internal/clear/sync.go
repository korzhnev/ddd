package clear

type Filesystem interface {
	Copy(from string, to string) error
	Move(from string, to string) error
	Delete(path string) error
}

func sync(
	reader func(path string) (map[string]string, error),
	filesystem Filesystem,
	source string,
	dest string,
) error {
	sourceHashes, err := reader(source)
	if err != nil {
		return err
	}
	destHashes, err := reader(dest)
	if err != nil {
		return err
	}

	for sha, filename := range sourceHashes {
		if _, ok := destHashes[sha]; !ok {
			sourcePath := getPath(source, filename)
			destPath := getPath(dest, filename)
			if err := filesystem.Copy(sourcePath, destPath); err != nil {
				return err
			}
		} else if destHashes[sha] != filename {
			oldDestPath := getPath(dest, destHashes[sha])
			newDestPath := getPath(dest, filename)
			if err := filesystem.Move(oldDestPath, newDestPath); err != nil {
				return err
			}
		}
	}

	for sha, filename := range destHashes {
		if _, ok := sourceHashes[sha]; !ok {
			if err := filesystem.Delete(getPath(dest, filename)); err != nil {
				return err
			}
		}
	}

	return nil
}

func getPath(dir string, file string) string {
	return dir + "/" + file
}
