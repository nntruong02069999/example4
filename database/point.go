package database

import (
	"errors"
	"log"
)

type Point struct {
	User_id    string
	Points     int64
	Max_points int64
}

func (db *Db) CreatePoint(point *Point) error {
	_, err := db.engine.Insert(point)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (db *Db) UpdatePoint(point *Point) (error) {
	_, err := db.engine.Update(point)
	if err != nil {
		log.Println("Update failed")
		return err
	}
	return nil
}

func (db *Db) GetPointById(id string) (*Point, error) {
	point := &Point{User_id: id}
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
