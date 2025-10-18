package dotenv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Read environment files and load environment variables in project
Note:
- Environment file shall end with .env only.
- If there exists a key with no value, the value for that is set to empty string.
- Required variables are deleted map 'rqdEnvList'.
*/
func LoadEnvFile(files ...string) error {
	for _, file := range files {
		checkFile := strings.Split(file, ".")
		if len(checkFile) == 2 && checkFile[1] == "env" {
			f, err := os.Open(file)
			if err != nil {
				return err
			}
			defer f.Close()

			scanner := bufio.NewScanner(f)

			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if len(line) > 0 && strings.Contains(line, "=") {
					// Split the key=value
					arr := strings.SplitN(line, "=", 2)

					// Check for empty key
					if arr[0] == "" {
						return ErrEmptyKey
					}

					if err := os.Setenv(arr[0], arr[1]); err != nil {
						return err
					}
					deleteFromRqd(arr[0], arr[1])
				}
			}
		} else {
			return fmt.Errorf("%s file format is not supported", file)
		}
	}
	return nil
}
