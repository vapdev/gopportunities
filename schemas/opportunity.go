package schemas

import (
	"gorm.io/gorm"
)

type Opportunity struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Link        string `json:"link"`
	Salary      int64  `json:"salary"`
	IsRemote    bool   `json:"isRemote"`
}
