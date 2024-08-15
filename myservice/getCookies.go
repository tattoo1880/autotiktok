package myservice

import (
	"fmt"
	"strings"
	"github.com/go-rod/rod/lib/proto"
)



func ParseCookies(cookieStr string) []*proto.NetworkCookie {
	cookieList := []*proto.NetworkCookie{}
	cookiePairs := strings.Split(cookieStr, "; ")
	for _, pair := range cookiePairs {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) == 2 {
			cookie := &proto.NetworkCookie{
				Name:  parts[0],
				Value: parts[1],
			}
			cookieList = append(cookieList, cookie)
		}
	}
	fmt.Println(cookieList)
	return cookieList
}