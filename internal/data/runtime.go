package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Error to return if it's imposible to parse or conver JSON
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

// Now Runtime type satisfies Marshaler interfase, so in this
// method I can customize how Runtime type should display in JSON
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	// Wrap jsonValue string in double quotes it is nessecery to be VALID JSON string
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

// Satisfy Unmarshaler interface.
// Conver JSON string value (ex. "102 mins") to Runtime type
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	// Conver i (int32) to Runtime type and assign this to the receiver.
	*r = Runtime(i)

	return nil
}
