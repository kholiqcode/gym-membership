package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gym/cmd/domain/class/entity"
	"gym/pkg/database"
	"strconv"
)

type ClassRepositoryImpl struct {
	Db *gorm.DB
}

func (r *ClassRepositoryImpl) FindAll(ctx echo.Context, pagination *database.Pagination) (*entity.ClassList, error) {
	var classes entity.ClassList
	isOffline := ctx.QueryParam("isOffline")
	query := r.Db.Debug().Preload("ClassCategory").Scopes(database.Paginate(classes, pagination, r.Db)).Preload(clause.Associations)
	if isOffline != "" {
		_isOffline, _ := strconv.ParseBool(isOffline)
		query = query.Where("is_offline = ?", _isOffline)
	}

	if e := query.Find(&classes).Error; e != nil {
		return nil, e
	}

	return &classes, nil
}

func (r *ClassRepositoryImpl) FindByIds(ctx echo.Context, classessId []int) (*entity.ClassList, error) {
	var classes entity.ClassList

	if e := r.Db.Debug().Preload("ClassCategory").Preload(clause.Associations).Find(&classes, classessId).Error; e != nil {
		return nil, e
	}

	return &classes, nil
}

func (r *ClassRepositoryImpl) Find(ctx echo.Context, id uint) (*entity.Class, error) {
	var class entity.Class
	if e := r.Db.Debug().First(&class, id).Error; e != nil {
		return nil, e
	}

	return &class, nil
}

func (r *ClassRepositoryImpl) FindByTrainer(ctx echo.Context, trainerID uint) (*entity.Class, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ClassRepositoryImpl) Insert(ctx echo.Context, class *entity.Class) (*entity.Class, error) {
	if e := r.Db.Debug().Create(&class).Error; e != nil {
		return nil, e
	}
	return class, nil
}

func (r *ClassRepositoryImpl) Update(ctx echo.Context, entity *entity.Class) (*entity.Class, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ClassRepositoryImpl) Delete(ctx echo.Context, id uint) (*entity.Class, error) {
	//TODO implement me
	panic("implement me")
}
