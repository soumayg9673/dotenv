package dotenv_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/soumayg9673/dotenv"
)

func TestGetString(t *testing.T) {
	testCases := []struct {
		Name        string
		SetEnv      bool
		EnvKey      string
		EnvValue    string
		EnvDefValue string
		Err         error
	}{
		{
			Name:        "happy path test with set env key-value",
			SetEnv:      true,
			EnvKey:      "PACKAGE_TEST_1",
			EnvValue:    "dotenv_1",
			EnvDefValue: "1_dotenv",
			Err:         errors.New("recived default value"),
		},
		{
			Name:        "happy path test with set env key only",
			SetEnv:      true,
			EnvKey:      "PACKAGE_TEST_2",
			EnvValue:    "",
			EnvDefValue: "2_dotenv",
			Err:         errors.New("recived default value"),
		},
		{
			Name:        "un-happy path test with no env",
			SetEnv:      false,
			EnvKey:      "PACKAGE_TEST_3",
			EnvValue:    "",
			EnvDefValue: "3_dotenv",
			Err:         errors.New("recived value"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Set env variable
			if tc.SetEnv {
				os.Setenv(tc.EnvKey, tc.EnvValue)
			}

			check := dotenv.GetString(tc.EnvKey, tc.EnvDefValue)
			switch tc.SetEnv {
			case true:
				if check == tc.EnvDefValue {
					t.Error(tc.Err)
				}
			case false:
				if check != tc.EnvDefValue {
					t.Error(tc.Err)
				}
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	testCases := []struct {
		Name        string
		SetEnv      bool
		EnvKey      string
		EnvValue    int
		EnvDefValue int
		Err         error
	}{
		{
			Name:        "happy path test with set env key-value",
			SetEnv:      true,
			EnvKey:      "PACKAGE_TEST_1",
			EnvValue:    1,
			EnvDefValue: 10,
			Err:         errors.New("recived default value"),
		},
		{
			Name:        "happy path test with set env key only",
			SetEnv:      true,
			EnvKey:      "PACKAGE_TEST_2",
			EnvValue:    2,
			EnvDefValue: 20,
			Err:         errors.New("recived default value"),
		},
		{
			Name:        "un-happy path test with no env",
			SetEnv:      false,
			EnvKey:      "PACKAGE_TEST_3",
			EnvValue:    3,
			EnvDefValue: 30,
			Err:         errors.New("recived value"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Set env variable
			if tc.SetEnv {
				os.Setenv(tc.EnvKey, fmt.Sprintf("%v", tc.EnvValue))
			}

			check := dotenv.GetInt(tc.EnvKey, tc.EnvDefValue)
			switch tc.SetEnv {
			case true:
				if check == tc.EnvDefValue {
					t.Error(tc.Err)
				}
			case false:
				if check != tc.EnvDefValue {
					t.Error(tc.Err)
				}
			}
		})
	}
}
