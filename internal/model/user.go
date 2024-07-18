package model

type UserAuth struct {
	ID            int    `gorm:"primary_key;auto_increment;column:id" json:"id"`
	Username      string `gorm:"uniqueIndex;type:varchar(50);column:username;not null" json:"userName"`
	Password      string `gorm:"uniqueIndex;type:varchar(100);column:password;not null" json:"-"`
	LoginType     int    `gorm:"uniqueIndex;column:login_type;type:tinyint(1);comment:登录类型" json:"login_ype"`
	IpAddress     string `gorm:"type:varchar(20);comment:登录IP地址;column:ip_address" json:"ip_address"`
	IpSource      string `gorm:"type:varchar(50);comment:IP来源;column:ip_source" json:"ip_source"`
	LastLoginTime int64  `gorm:"last_login_time;type:bigint(20);comment:上次登录时间" json:"last_login_time"`
	IsDisable     bool   `json:"is_disable" gorm:"column:is_disable;comment:是否禁用;type:tinyint(1);default:0"`
	IsSuper       bool   `json:"is_super" gorm:"column:is_super;comment:超级管理员;type:tinyint(1);default:0"` // 超级管理员只能后台设置
	CreatedAt     int64  `gorm:"column:created_at;comment:创建时间;type:bigint(20)" json:"created_at"`
	UpdatedAt     int64  `gorm:"column:updated_at;comment:更新时间;type:bigint(20)" json:"updated_at"`
}

func (u UserAuth) TableName() string {
	return "user_auth"
}

type UserInfo struct {
	Model
	Email    string `json:"email" gorm:"type:varchar(30)"`
	Nickname string `json:"nickname" gorm:"unique;type:varchar(30);not null"`
	Avatar   string `json:"avatar" gorm:"type:varchar(1024);not null"`
	Intro    string `json:"intro" gorm:"type:varchar(255)"`
	Website  string `json:"website" gorm:"type:varchar(255)"`
}
