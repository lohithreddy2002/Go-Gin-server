package models

import "time"

type Hai struct {
	Name       string    `form:"hai" binding:"required" json:"hai"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano" json:"created"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix" binding:"required"`
}
