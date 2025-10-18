package dotenv

import (
	"errors"
	"strings"
)

var (
	ErrEmptyKey = errors.New("key is empty")
	// error messages for required fields
	rqdEnvErrs    = []string{}
	errNoKey      = "missing %s key"
	errNoKeyValue = "missing key-value for %s"
	errNoValue    = "missing value for %s"
)

type reqdEnvErrType int

const (
	ERR_NO_KEY       reqdEnvErrType = iota // Key does not exists in .env file.
	ERR_NO_KEY_VALUE                       // Key-value does not exists in .env file.
	ERR_NO_VALUE                           // Empty value for a key in .env file.
)

// Set custom error message for different types
// {key not found}.
// {key-value not found},
// {value not found}.
// Custom message can contain %s to mention key.
func SetErrorMsg(err string, terr reqdEnvErrType) {
	switch terr {
	case ERR_NO_KEY:
		errNoKey = err
	case ERR_NO_KEY_VALUE:
		errNoKeyValue = err
	case ERR_NO_VALUE:
		errNoValue = err
	}
}

func formatErrorMsg(key, err string) string {
	return strings.ReplaceAll(err, "%s", key)
}
