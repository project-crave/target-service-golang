package repository

import "crave/hub/cmd/model"

type IRepository interface {
	Create(*model.Target) (*model.Target, error)
	FindByWorkIdAndPrevious(workId, previous uint) ([]model.Target, error)
}
