package models

type Statuses struct {
	ID     uint     `gorm:"column:id_statuses;primaryKey" json:"id"`
	Name   string   `gorm:"column:name_statuses" json:"name"`
	Movies []Movies `gorm:"many2many:movie_statuses;foreignKey:ID;joinForeignKey:id_statuses;References:ID;joinReferences:id_movies"`
}

func (Statuses) TableName() string {
	return "statuses"
}
