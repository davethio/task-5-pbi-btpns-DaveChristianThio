package models

type Photo struct {
	ID       int64  `gorm:"primaryKey" json:"photoId"`
	Title    string `gorm:"size:255" json:"title"`
	Caption  string `gorm:"size:512" json:"caption"`
	PhotoUrl string `gorm:"size:512" json:"photo_url"`
	UserID   int64  `gorm:"not null" json:"user_id"`
	Users    User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"user"`
}
