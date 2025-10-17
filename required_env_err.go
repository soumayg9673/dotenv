package dotenv

import "errors"

var (
	ErrAddRequiredKey      = errors.New("key with no value required already added to list")
	ErrAddRequiredKeyValue = errors.New("key with value required already added to list")
)
