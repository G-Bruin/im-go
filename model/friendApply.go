package model

type FriendApply struct {
	ID       int    `gorm:"primary_key" json:"id"`
	UserId   int    `gorm:"unique_index:idx_user_friend_id"`
	FriendId int    `gorm:"unique_index:idx_user_friend_id"`
	IsDown   string `gorm:"type:enum('agree','refuse','untreated');default:'untreated'"`
	BaseModel
	DeletedAt
}
