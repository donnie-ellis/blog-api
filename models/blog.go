// models/blog.go

package models

import "time"

type Blog struct {
	ID uint `json:"id" gorm:"primary_key"`
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Title string `json:"title"`
	Text string `json:"text"`
}