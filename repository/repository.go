package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type repository[T any] struct {
	db  *gorm.DB
	ctx context.Context
}

func NewRepository[T any](db *gorm.DB, ctx context.Context) *repository[T] {
	return &repository[T]{
		db:  db,
		ctx: ctx,
	}
}

// Create insert the value into database
func (r *repository[T]) Create(values interface{}) error {
	err := r.db.WithContext(r.ctx).Create(values).Error
	if errors.Is(err, gorm.ErrEmptySlice) {
		return nil
	}
	if err != nil {
		return err
	}

	return nil
}

// FindAll find records that match given conditions
func (r *repository[T]) FindAll(conds ...interface{}) ([]T, error) {
	list := make([]T, 0)
	err := r.db.WithContext(r.ctx).Find(&list, conds...).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

// FindOne find first record that match given conditions, order by primary key
func (r *repository[T]) FindOne(conds ...interface{}) (*T, error) {
	item := new(T)
	err := r.db.WithContext(r.ctx).First(item, conds...).Error
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return nil, errors.New("not found")
	}

	if err != nil {
		return nil, err
	}

	return item, nil
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (r *repository[T]) Update(values interface{}) error {

	err := r.db.WithContext(r.ctx).Updates(values).Error
	if errors.Is(err, gorm.ErrEmptySlice) {
		return nil
	}
	if err != nil {
		return err
	}

	return nil
}

// Delete value match given conditions, if the value has primary key, then will including the primary key as condition
func (r *repository[T]) Delete(conds ...interface{}) error {
	item := new(T)
	err := r.db.WithContext(r.ctx).Delete(item, conds...).Error
	if err != nil {
		return err
	}

	return nil
}

// condition
func (r *repository[T]) Where(query interface{}, args ...interface{}) *repository[T] {
	r.db = r.db.Where(query, args...)
	return r
}

func (r *repository[T]) Order(value interface{}) *repository[T] {
	r.db = r.db.Order(value)
	return r
}
