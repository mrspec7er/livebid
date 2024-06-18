package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string         `json:"id" gorm:"primaryKey; index:priority:1; type:varchar(63)"`
	Picture       string         `json:"picture" gorm:"type:text"`
	Email         string         `json:"email" gorm:"type:varchar(63)"`
	VerifiedEmail bool           `json:"verified_email"`
	Password      string         `json:"password" gorm:"type:text"`
	Role          string         `json:"role" gorm:"type:varchar(63)"`
	Items         []*Item        `json:"items" gorm:"foreignKey:CreatorID;references:ID"`
	Owns          []*Item        `json:"owns" gorm:"foreignKey:OwnerID;references:ID"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Item struct {
	Number      string `json:"number" gorm:"primaryKey; index:priority:1; type:varchar(63)"`
	CreatorID   string `json:"creatorId" gorm:"type:varchar(63)"`
	Creator     *User  `json:"creator"`
	OwnerID     string `json:"ownerId" gorm:"type:varchar(63); default:null"`
	Owner       *User  `json:"owner"`
	Status      string `json:"status" gorm:"type:varchar(63)"`
	ApprovedBy  string `json:"approvedBy" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:text"`
	OpenDate    time.Time
	CloseDate   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type TradeMessage struct {
	ID          string `json:"id" gorm:"primaryKey; index:priority:1; type:varchar(63)"`
	CreatorID   string `json:"creatorId" gorm:"type:varchar(63)"`
	Creator     *User  `json:"creator"`
	ItemNumber  string `json:"itemNumber" gorm:"type:varchar(63)"`
	Item        *Item  `json:"item"`
	CreatorItem bool   `json:"creatorItem"`
	Message     string `json:"message" gorm:"type:text"`
}
