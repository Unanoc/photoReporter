package valid

import "os"

// GetValidPath checks if dir exists and remove "/" at the end of path.
func GetValidPath(path string) (string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", err
	}

	if len(path) > 1 && path[len(path)-1:] == "/" { // last character of string
		return path[:len(path)-1], nil
	}
	return path, nil
}
