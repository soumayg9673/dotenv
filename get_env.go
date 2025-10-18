package dotenv

import (
	"os"
	"strconv"
)

// Note: Returns the value if Key is found or STD if not found.
func GetString(key, std string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return std
	}
	return val
}

// Convert environment key-value to base-10 integer value.
// Note: Returns the value if KEY is found or STD if not found.
func GetInt(key string, std int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return std
	}

	// Convert the string to base 10 int
	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return std
	}

	return valAsInt
}
