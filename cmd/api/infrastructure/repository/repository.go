package repository

import (
	"crave/hub/cmd/model"
	"crave/shared/database"
)

type Repository struct {
	mysql database.MysqlWrapper
}

// FindByWorkIdAndPrevious implements IRepository.
func (r *Repository) FindByWorkIdAndPrevious(workId uint, previous uint) ([]model.Target, error) {
	var targets []model.Target
	result := r.mysql.Driver.Table("target").
		Where("work_id = ? AND previous = ?", workId, previous).Find(&targets)
	if result.Error != nil {
		return nil, result.Error
	}
	return targets, nil
}

func (r *Repository) GetLastIndex() (uint, error) {
	var lastIndex uint
	result := r.mysql.Driver.Table("work").Select("COALESCE(MAX(id), 0) as last_index").Scan(&lastIndex)
	if result.Error != nil {
		return 0, result.Error
	}
	return lastIndex, nil
}

func NewRepository(mysql *database.MysqlWrapper) *Repository {
	return &Repository{mysql: *mysql}
}

func (r *Repository) Create(target *model.Target) (*model.Target, error) {

	result := r.mysql.Driver.Create(target)
	if result.Error != nil {
		return nil, result.Error
	}

	return target, nil
}
