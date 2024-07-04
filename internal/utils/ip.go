package utils

import (
	"context"
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/nepu-SidneyYu/blog-grpc/internal/logs"
	"go.uber.org/zap"
)

var vIndex []byte

func GetIpSource(ipAddress string) string {
	var dbPath = "../region/ip2region.xdb" // IP 数据库文件
	// 完全基于文件查询, 每次都读取文件
	// searcher, err := xdb.NewWithFileOnly(dbPath)

	// 缓存 VectorIndex 索引, 减少一次固定的 IO 操作
	if vIndex == nil {
		var err error
		vIndex, err = xdb.LoadVectorIndexFromFile(dbPath)
		if err != nil {
			logs.Error(context.Background(), "failed toload vector index: %v", zap.String("error", err.Error()))
			return ""
		}
	}
	searcher, err := xdb.NewWithVectorIndex(dbPath, vIndex)

	if err != nil {
		logs.Error(context.Background(), "failed to create ip2region searcher", zap.String("error", err.Error()))
		return ""
	}
	defer searcher.Close()

	// 国家|区域|省份|城市|ISP
	// 只有中国的数据绝大部分精确到了城市, 其他国家部分数据只能定位到国家, 后面的选项全部是 0
	region, err := searcher.SearchByStr(ipAddress)
	if err != nil {
		logs.Error(context.Background(), "failed to search ip: %v", zap.String("error", err.Error()))
		return ""
	}
	return region
}

func GetIpSourceSimpleIdle(ipAddress string) string {
	region := GetIpSource(ipAddress) // 国家|区域|省份|城市|ISP
	// 检测到是内网, 直接返回 "内网IP"
	// 0|0|0|内网IP|内网IP
	if strings.Contains(region, "内网IP") {
		return "内网IP"
	}

	// 一般无法获取到区域
	// 中国|0|江苏省|苏州市|电信
	ipSource := strings.Split(region, "|")
	if ipSource[0] != "中国" && ipSource[0] != "0" {
		return ipSource[0]
	}
	if ipSource[2] == "0" {
		ipSource[2] = ""
	}
	if ipSource[3] == "0" {
		ipSource[3] = ""
	}
	if ipSource[4] == "0" {
		ipSource[4] = ""
	}
	if ipSource[2] == "" && ipSource[3] == "" && ipSource[4] == "" {
		return ipSource[0]
	}
	return ipSource[2] + ipSource[3] + " " + ipSource[4]
}
