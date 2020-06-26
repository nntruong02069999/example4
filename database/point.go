package database

import (
	"errors"
	"log"
)

type Point struct {
	UserID   string `json:user_id`
	Points     int64
	MaxPoints int64 `json:max_points`
}

func (db *Db) CreatePoint(point *Point) error {
	_, err := db.engine.Insert(point)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (db *Db) UpdatePoint(point, conditions *Point) (error) {
	_, err := db.engine.Update(point, conditions)
	if err != nil {
		log.Println("Update failed")
		return err
	}
	return nil
}

func (db *Db) GetPointById(id string) (*Point, error) {
	point := &Point{UserID: id}
	has, err := db.engine.Get(point)
	if err != nil {
		log.Println("Failed")
		return nil , err
	}
	if !has {
		return nil, errors.New("not found")
	}
	return point, nil
}
