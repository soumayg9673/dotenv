package dotenv

import "errors"

type requiredEnvErrors []string

var (
	rqdEnvErrs            = []requiredEnvErrors{}
	ErrAddRqdKey          = errors.New("key with no value required already added to list")
	ErrAddRqdKeyValue     = errors.New("key with value required already added to list")
	errRqdKeyMissing      = "missing %s key"
	errRqdKeyValueMissing = "missing key-value for %s"
	errRqdValueMissing    = "missing value for %s"
)

type reqdEnvErrType int

const (
	ERRKEY      reqdEnvErrType = iota // errRqdKeyMissing
	ERRKEYVALUE                       // errRqdKeyValueMissing
	ERRVALUE                          // errRqdValueMissing
)

// Set custom error message for different types
// {key not found}.
// {key-value not found},
// {value not found}.
// Custom message can contain %s once to mention key.
func SetErrorMsg(err string, terr reqdEnvErrType) {
	switch terr {
	case ERRKEY:
		errRqdKeyMissing = err
	case ERRKEYVALUE:
		errRqdKeyValueMissing = err
	case ERRVALUE:
		errRqdValueMissing = err
	}
}
