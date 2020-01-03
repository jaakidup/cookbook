package log

import (
	"log"

	"github.com/pkg/errors"
)

// OriginalError returns the original error
func OriginalError() error {
	return errors.New("Error occurred")
}

// PassThroughError calls OriginalError
// and forwards the error after wrapping
func PassThroughError() error {
	err := OriginalError()
	// no need to check error since this works with nil

	return errors.Wrap(err, " In PassThroughError")

}

// FinalDestination deals with the error
// and doesn't forward it
func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		// we log because an unexpected error occurred!
		log.Printf("an error occurred: %s\n", err.Error())
		return
	}
}
