package cnftypes

import (
	"fmt"
	"net/url"
)

// URL allows to specify URLs in config files parsing directly to the native type.
type URL struct {
	*url.URL
}

// UnmarshalTOML implements toml.Unmarshaller.
func (u *URL) UnmarshalTOML(v interface{}) (err error) {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("cnftypes.URL: can't unmarshal non string value (%T) into URL", v)
	}

	u.URL, err = url.Parse(s)
	return
}
