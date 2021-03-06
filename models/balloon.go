package models

import (
	_ "github.com/astaxie/beego/orm"
)


type Balloon struct {
	BalloonId			int32			`orm:"auto"`
	UserId				int32			`orm:"null"`
	Sid					int32			`orm:"null"`
	Cid					int32			`orm:"null"`
	Pid					int32			`orm:"null"`
	Status				int8			`orm:"default(0)"`
}
