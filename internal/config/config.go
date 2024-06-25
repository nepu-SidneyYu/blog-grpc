package config

// Config结构体，包含MySqlConfig字段
type Config struct {
	// MySql字段，指向MySqlConfig结构体
	MySql *MySqlConfig
	
}

type MySqlConfig struct {
    // 数据库连接信息
    Host     string
    Port     int
    Username string
    Password string
    Database  string
}
    

func 