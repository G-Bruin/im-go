package model

type Massage struct {
	ID         int    `gorm:"primary_key" json:"id"`
	UserId     int    `gorm:"index" json:"user_id"`
	TargetId   int    `gorm:"index" json:"target_id"`
	Resource   string `gorm:"type:varchar(255)" json:"resource"`
	Type       string `gorm:"type:enum('text','voice','image','video');default:'text'" json:"type"`
	IsReceived string `gorm:"type:enum('Y','N');default:'N'" json:"is_received"`
	BaseModel
}
