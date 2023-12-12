package dao

import (
	"log"

	"github.com/laonsx/tnine/internal/pkg/errors"
)

func LogErrFromDown(printErr bool) error {

	err := errors.New("err form dao")
	if err != nil {
		err = errors.WithStack(err)
		if printErr {
			log.Printf("dao LogErrFromDown: %+v", err)
		}
		return err
	}

	return nil
}
