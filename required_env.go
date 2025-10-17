package dotenv

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
// Throws error if key found in required list.
func AddRqdKey(key string, val bool) error {
	if v, ok := rqdEnvList[key]; ok {
		if v {
			return ErrAddRqdKeyValue
		}
		return ErrAddRqdKey
	}
	rqdEnvList[key] = val
	return nil
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
		if value == "" && !v {
			delete(rqdEnvList, key)
		} else if value != "" && v {
			delete(rqdEnvList, key)
		}
	}
}
