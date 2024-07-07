package utils

import (
	"fmt"
	"testing"
)

func TestGetIpSourceSimpleIdle(t *testing.T) {
	region := GetIpSourceSimpleIdle("81.2.69.142")
	GeoIP("192.168.1.2")
	fmt.Println(region)
}
