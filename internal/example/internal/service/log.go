package service

import "github.com/laonsx/tnine/internal/example/internal/dao"

func LogErr(printErr bool) error {

	return dao.LogErrFromDown(printErr)
}
