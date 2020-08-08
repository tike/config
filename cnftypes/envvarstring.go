package cnftypes

import (
	"fmt"
	"os"
)

// EnvVarString expands contained environment variables.
type EnvVarString string

// String implements fmt.Stringer
func (e EnvVarString) String() string {
	return string(e)
}

// UnmarshalTOML implements toml.Unmarshaller.
func (e *EnvVarString) UnmarshalTOML(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("cnftypes.EnvVarString: can't unmarshal non string value (%T) into Duration", v)
	}

	*e = EnvVarString(os.ExpandEnv(s))
	return nil
}
