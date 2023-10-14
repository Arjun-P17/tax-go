package mongodb

import (
	"errors"
)

type Options struct {
	URI string
}

func (opts *Options) Validate() error {
	if opts == nil {
		return errors.New("mongodb options cannot be a nil pointer")
	}
	if opts.URI == "" {
		return errors.New("URI cannot be empty for mongodb")
	}

	return nil
}
