package model

type Friend struct {
	ID       int    `gorm:"primary_key" json:"id"`
	UserId   int    `gorm:"unique_index:idx_user_friend_id"`
	FriendId int    `gorm:"unique_index:idx_user_friend_id"`
	IsBlock  string `gorm:"type:enum('Y','N');default:'N'"`
	BaseModel
	DeletedAt
}
