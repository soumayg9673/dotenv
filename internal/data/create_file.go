package data

import (
	"fmt"
	"os"
	"strings"
)

func CreateFile(file string, envs map[string]string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	s := strings.Builder{}

	l := len(envs)
	tl := 0

	for k, v := range envs {
		tl++
		switch tl {
		case l:
			s.WriteString(fmt.Sprintf("%s=%s", k, v))
		default:
			s.WriteString(fmt.Sprintf("%s=%s\n", k, v))
		}
	}

	f.Write([]byte(s.String()))

	return nil
}
