package dotenv

import (
	"fmt"
	"maps"
	"strings"
)

/*
List of required environment variables
Notes:
 1. Key is same in .env files.
 2. Value is set to false/true.
    2.1. If value is required set TRUE.
    2.2. If value can be null set FALSE.
*/
type requiredEnv map[string]bool

var (
	rqdEnvList = requiredEnv{}
)

// Add environment variable key as required. Set value required with true/false.
func AddRqdKey(key string, val bool) error {
	rqdEnvList[key] = val
	return nil
}

// Add multiple environment variable key as required. Set value required with true/false.
func AddMulRqdKeys(rqds map[string]bool) {
	maps.Copy(rqdEnvList, rqds)
}

// Delete environment variable key from required.
func DeleteRqdKey(keys ...string) {
	for _, k := range keys {
		delete(rqdEnvList, k)
	}
}

// Get list of required environment variables
func GetAllRqd() map[string]bool {
	return rqdEnvList
}

// Delete the key from required list.
// This functionality is intended to delete the each key from required list
// after successful setting key-value in the project.
func deleteFromRqd(key, value string) {
	if v, ok := rqdEnvList[key]; ok {
		if value == "" && v { // error: value required
			errMsg := formatErrorMsg(key, errNoValue)
			rqdEnvErrs = append(rqdEnvErrs, errMsg)
		}
		delete(rqdEnvList, key)
	}
}

func ValidateRqdEnv() error {
	if len(rqdEnvList) != 0 {
		for k, v := range rqdEnvList {
			var err string
			switch v {
			case true:
				err = fmt.Sprintf(errNoKeyValue, k)
			case false:
				err = fmt.Sprintf(errNoKey, k)
			}
			rqdEnvErrs = append(rqdEnvErrs, err)
		}
		return fmt.Errorf("add the following environment variables:\n%s",
			strings.Join(rqdEnvErrs, ",\n"))
	}
	return nil
}
