package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// Now Runtime type satisfies Marshaler interfase, so in this
// method I can customize how Runtime type should display in JSON
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	// Wrap jsonValue string in double quotes it is nessecery to be VALID JSON string
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}
