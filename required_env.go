package dotenv

import (
	"errors"
)

/*
List of required environment variables
Notes:
 1. Key is same in .env files.
 2. Value is set to false/true.
    2.1. If value is required set TRUE.
    2.2. If value can be null set FALSE.
*/
type mandatoryEnv map[string]bool

var (
	mandEnvList            = mandatoryEnv{}
	ErrAddRequiredKey      = errors.New("key with no value required already added to list")
	ErrAddRequiredKeyValue = errors.New("key with value required already added to list")
)

// Add environment variable key as required. Set value required with true/false.
// Throws error if key found in required list.
func AddRqdKey(key string, val bool) error {
	if v, ok := mandEnvList[key]; ok {
		if v {
			return ErrAddRequiredKeyValue
		}
		return ErrAddRequiredKey

	}
	mandEnvList[key] = val
	return nil
}

// Delete environment variable key from required.
func DeleteRqdKey(key string) {
	delete(mandEnvList, key)
}
