package user

import (
	"echo.osaxon/model"
)

type Store interface {
	GetAll() ([]*model.User, error)
}
