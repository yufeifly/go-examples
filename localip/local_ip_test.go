package localip

import (
	"fmt"
	"testing"
)

func TestGetLocalIP(t *testing.T) {
	IPs, err := GetLocalIP()
	if err != nil {
		t.Error(err)
	} else {
		for _, ip := range IPs {
			fmt.Println(ip)
		}
	}
}
