package models

import "time"

type Movies struct {
	ID          uint        `gorm:"column:id_movies;primary_key" json:"id"`
	Name        string      `gorm:"column:name_movies" json:"name_movies"`
	Poster      string      `gorm:"column:poster" json:"poster"`
	Actors 		[]Actors 	`gorm:"many2many:movie_actors;foreignKey:ID;joinForeignKey:id_movies;References:ID;joinReferences:id_actors" json:"actors"`
	Genres 		[]Genres 	`gorm:"many2many:movie_genres;foreignKey:ID;joinForeignKey:id_movies;References:ID;joinReferences:id_genres" json:"genres"`
	Broadcast   []Broadcast `gorm:"many2many:movie_broadcast;foreignKey:ID;joinForeignKey:id_movies;References:ID;joinReferences:id_broadcast" json:"broadcast"`
	Statuses    []Statuses  `gorm:"many2many:movie_statuses;foreignKey:ID;joinForeignKey:id_movies;References:ID;joinReferences:id_statuses" json:"statuses"`
	Description string      `gorm:"column:description_movies" json:"description_movies"`
	Rating      float32     `gorm:"column:rating" json:"rating"`
	Views 	    int         `gorm:"column:views" json:"views"`
	Duration    int         `gorm:"column:duration" json:"duration"`
	CreatedAt   time.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time   `gorm:"column:updated_at" json:"updated_at"`
}