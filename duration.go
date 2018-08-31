package config

import (
	"fmt"
	"time"
)

// Duration allows to specify durations in config files using the intuitive go
// syntax.
type Duration struct {
	time.Duration
}

// UnmarshalTOML implements toml.Unmarshaller.
func (d *Duration) UnmarshalTOML(v interface{}) (err error) {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("config.Duration: can't unmarshal non string value (%T) into Duration", v)
	}

	d.Duration, err = time.ParseDuration(s)
	return
}
