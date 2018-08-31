package config

// PostProcessor can be implemented by config types to do some processing
// after parsing.
type PostProcessor interface {
	PostProcess() error
}

// Sanitizer can be implemented by config types to perform sanitization.
type Sanitizer interface {
	Sanitize() error
}

// Validator can be implied by config types to perm validation.
type Validator interface {
	Validate() error
}
