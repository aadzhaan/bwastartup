package campaign

import (
	"bwastartup/user"
	"time"
)

type Campaign struct {
	Id               int
	UserId           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
	User             user.User
}

type CampaignImage struct {
	Id         int
	CampaignId int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
