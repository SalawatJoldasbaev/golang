package entity

import "salawat/api/common/base"

type Category struct {
	base.Model
	Title string `json:"title" gorm:"not null;unique"`
}
