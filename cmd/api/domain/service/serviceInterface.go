package service

import (
	"crave/hub/cmd/model"
)

type IService interface {
	Init(origin, destination *model.Target) error
}
