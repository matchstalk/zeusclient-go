package zeusclient

import "time"

type CheckPermRsp struct {
	Code int `json:"code"`
	Data struct {
		Result bool `json:"result"`
	} `json:"data"`
}

type UserPermsRsp struct {
	Code int `json:"code"`
	Data struct{
		Result []string `json:"result"`
	} `json:"data"`
}

type UserMenusRsp struct {
	Code int `json:"code"`
	Data struct{
		Result []Menus  `json:"result"`
	} `json:"data"`
}

type Menus struct {
	Id             int       `json:"id"`
	ParentId       int       `json:"parent_id"`
	DomainId       int       `json:"domain_id"`
	Name           string    `json:"name"`
	Url            string    `json:"url"`
	Perms          string    `json:"perms"`
	Alias          string    `json:"alias"`
	MenuType       int       `json:"menu_type"`
	Icon           string    `json:"icon"`
	OrderNum       int       `json:"order_num"`
	CreateTime     time.Time `gorm:"type:time;column:create_time;not null;default:CURRENT_TIMESTAMP" json:"created_time,omitempty" example:"2019-07-10 0:39"`
	LastUpdateTime time.Time `gorm:"type:time;column:last_update_time;not null;default:CURRENT_TIMESTAMP ON UPDATE" json:"last_update_time,omitempty" example:"2019-07-10 0:39"`
}