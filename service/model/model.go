package model

import (
	"log"
	"time"
)

func initDB() {
	if err := engine.Sync2(new(User)); err != nil {
		log.Println(err)
	}
	if err := engine.Sync2(new(FriendRelation)); err != nil {
		log.Println(err)
	}
	if err := engine.Sync2(new(Event)); err != nil {
		log.Println(err)
	}
	if err := engine.Sync2(new(Participation)); err != nil {
		log.Println(err)
	}
	if err := engine.Sync2(new(Token)); err != nil {
		log.Println(err)
	}
}

// 用户数据结构
type User struct {
	Id      int64
	Name    string    `xorm:"notnull unique 'name'"`
	Email   string    `xorm:"notnull unique 'email'"`
	UserId  string    `xorm:"notnull unique 'user_id'"` //用于标识唯一性的id，通过第三方平台的用户唯一id进行加密得到
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type Token struct {
	Id      int64
	UserId  string    `xorm:"notnull unique 'user_id'"`
	Token   string    `xorm:"notnull unique 'token'"` //加密后的token数据
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type FriendRelation struct {
	Id         int64
	StarterId  string    `xorm:"notnull unique 'starter_id'"`  //发起者
	AccepterId string    `xorm:"notnull unique 'accepter_id'"` //接受者
	Confirmed  bool      `xorm:"notnull unique 'confirmed'"`   //是否得到对方确认
	Created    time.Time `xorm:"created"`
	Updated    time.Time `xorm:"updated"`
}

//事件
type Event struct {
	Id        int64
	Publisher string `xorm:"notnull 'user_id'"` //发布者id
	Name      string `xorm:"notnull 'name'"`    //事件名
	Address   string `xorm:"notnull 'address'"` //POI地址

	StartDate time.Time `xorm:"startDate"` //起始时间
	EndDate   time.Time `xorm:"endDate"`   //结束时间
	FinalDate time.Time `xorm:"finalDate"` //最终聚会时间

	Punishment string `xorm:"punishment"` //惩罚

	Lat float64 `xorm:"lat"` //纬度lat
	Lng float64 `xorm:"lng"` //经度lng

	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

//参与情况
type Participation struct {
	Id        int64
	EventId   int64     `xorm:"notnull 'event_id'"` //参与事件
	UserId    int64     `xorm:"notnull 'user_id'"`  //参与人
	StartDate time.Time `xorm:"startDate"`          //可参与起始时间
	EndDate   time.Time `xorm:"endDate"`            //可参与结束时间

	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
