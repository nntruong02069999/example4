package database

import (
	"errors"
	_"log"
)

type Point struct {
	UserId   string `json:"user_id"`
	Points     int64	`json:"points"`
	MaxPoints int64 `json:"max_points"`
}

func (database *Db) CreatePoint(point *Point) error {
	_, err := database.engine.Insert(point)
	if err != nil {
		return err
	}
	return nil
}

func (database *Db) UpdatePoint(point, conditions *Point) (error) {
	_, err := database.engine.Update(point, conditions)
	if err != nil {
		return err
	}
	return nil
}

func (database *Db) GetPointById(id string) (*Point, error) {
	point := &Point{UserId: id}
	has, err := database.engine.Get(point)
	if err != nil {
		return nil , err
	}
	if !has {
		return nil, errors.New("not found")
	}
	return point, nil
}
