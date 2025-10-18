package dotenv_test

import (
	"errors"
	"testing"

	"github.com/soumayg9673/dotenv"
	"github.com/soumayg9673/dotenv/internal/data"
)

func TestLoadEnvFile(t *testing.T) {
	testCases := []struct {
		Name string
		File string
		Envs map[string]string
		Err  error
	}{
		{
			Name: "happy path with .env file",
			File: ".env",
			Envs: map[string]string{},

			Err: nil,
		},
		{
			Name: "happy path with .env.test file",
			File: ".env.test",
			Envs: map[string]string{},

			Err: nil,
		},
		{
			Name: "invalid file name with test.env.test",
			File: "test.env.test",
			Envs: map[string]string{},
			Err:  dotenv.ErrInvalidFileFormat,
		},
		{
			Name: "no files specified",
			File: "",
			Envs: map[string]string{},
			Err:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// create file
			if tc.File != "" {
				if err := data.CreateFile(tc.File, tc.Envs); err != nil {
					t.Errorf("cannot create %s file to create test data", tc.File)
				}
			}

			if err := dotenv.LoadEnvFile(tc.File); err != nil {
				if !errors.Is(err, dotenv.ErrInvalidFileFormat) {
					t.Errorf("expected error %v, but got %v", dotenv.ErrInvalidFileFormat, err)
				}
			}

			// clean-up file
			data.DeleteFile(tc.File)
		})
	}
}
