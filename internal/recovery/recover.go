package recovery

import (
	"github.com/any/companies/internal/infr/logger"
	"github.com/pkg/errors"
)

func RecoverToError() error {
	if recoveryMessage := recover(); recoveryMessage != nil {
		return makeRecoveryError(recoveryMessage)
	}

	return nil
}

func RecoverToLog(logger logger.Logger) {
	if err := RecoverToError(); err != nil {
		logger.Error(err.Error())
	}
}

func makeRecoveryError(recoveryMessage interface{}) error {
	var err error
	switch x := recoveryMessage.(type) {
	case string:
		err = errors.New("recovered panic! " + x)
	case error:
		err = errors.Wrap(x, "recovered panic!")
	default:
		err = errors.New("recovered panic! unknown")
	}

	return err
}
