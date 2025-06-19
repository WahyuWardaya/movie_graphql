package models

type Broadcast struct {
	ID     uint     `gorm:"column:id_broadcast;primary_key" json:"id"`
	Name   string   `gorm:"column:name_Broadcast" json:"name_Broadcast"`
	Link   string   `gorm:"column:link_broadcasts" json:"link"`
	Movies []Movies `gorm:"many2many:movie_broadcast;foreignKey:ID;joinForeignKey:id_broadcast;References:ID;joinReferences:id_movies" json:"-"`
}

func (Broadcast) TableName() string {
	return "Broadcast"
}