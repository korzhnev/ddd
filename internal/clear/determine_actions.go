package clear

func determineActions(
	sourceHashes map[string]string,
	destHashes map[string]string,
	source string,
	dest string,
) (events []Action) {
	for sha, filename := range sourceHashes {
		if _, ok := destHashes[sha]; !ok {
			sourcePath := getPath(source, filename)
			destPath := getPath(dest, filename)
			events = append(events, NewCopyAction(sourcePath, destPath))
		} else if destHashes[sha] != filename {
			oldDestPath := getPath(dest, destHashes[sha])
			newDestPath := getPath(dest, filename)
			events = append(events, NewMoveAction(oldDestPath, newDestPath))
		}
	}

	for sha, filename := range destHashes {
		if _, ok := sourceHashes[sha]; !ok {
			events = append(events, NewDeleteAction(getPath(dest, filename)))
		}
	}

	return
}

func getPath(dir string, file string) string {
	return dir + "/" + file
}
