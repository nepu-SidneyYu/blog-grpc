package model

type UserAuth struct {
	ID            int    `gorm:"primary_key;auto_increment" json:"id"`
	Username      string `gorm:"unique;type:varchar(50)" json:"userName"`
	Password      string `gorm:"type:varchar(100)" json:"-"`
	LoginType     int    `gorm:"column:login_type;type:tinyint(1);comment:登录类型" json:"login_ype"`
	IpAddress     string `gorm:"type:varchar(20);comment:登录IP地址" json:"ip_address"`
	IpSource      string `gorm:"type:varchar(50);comment:IP来源" json:"ip_source"`
	LastLoginTime int64  `gorm:"last_login_time;type:datetime" json:"last_login_time"`
	IsDisable     bool   `json:"is_disable"`
	IsSuper       bool   `json:"is_super"` // 超级管理员只能后台设置
	CreatedAt     int64  `gorm:"column:created_at;comment:创建时间" json:"created_at"`
	UpdatedAt     int64  `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`
	//UserInfoId    int64        `json:"user_info_id"`
	//UserInfo      *UserInfo  `json:"info"`
	//Roles      []*Role   `json:"roles" gorm:"many2many:user_auth_role"`
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
