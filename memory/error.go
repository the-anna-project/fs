package memory

import (
	"fmt"

	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

func maskAnyf(err error, f string, v ...interface{}) error {
	if err == nil {
		return nil
	}

	f = fmt.Sprintf("%s: %s", err.Error(), f)
	newErr := errgo.WithCausef(nil, errgo.Cause(err), f, v...)
	newErr.(*errgo.Err).SetLocation(1)

	return newErr
}

var noSuchFileOrDirectoryError = errgo.New("no such file or directory")

// IsNoSuchFileOrDirectoryError asserts noSuchFileOrDirectoryError.
func IsNoSuchFileOrDirectoryError(err error) bool {
	return errgo.Cause(err) == noSuchFileOrDirectoryError
}
