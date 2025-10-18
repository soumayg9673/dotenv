package data

import "os"

func DeleteFile(file string) error {
	return os.Remove(file)
}
