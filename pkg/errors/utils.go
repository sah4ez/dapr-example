package errors

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

var Is = errors.Is
var As = errors.As
var Wrapf = errors.Wrapf

func ExitOnError(log zerolog.Logger, err error, msg string) {
	if err != nil {
		log.Panic().Err(err).Msg(msg)
	}
}
