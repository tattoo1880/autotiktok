package myservice

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func Jump(flist []string) []string {

	url := launcher.New().Headless(false).Proxy("http://127.0.0.1:10809").MustLaunch()
	// 通过URL连接到浏览器
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://www.tiktok.com/")
	page.MustWaitLoad()

	// 解析并设置Cookies
	cookieStr := "tt_csrf_token=vJ7z1ieg-0M9DBhit8p-1QWDOQqlX9TKtwJI; passport_csrf_token=ac684a02587ec5f1adcc1c3eab7ebfbb; passport_csrf_token=161c8527d39fa5ffa39541c01aa2736a; passport_csrf_token=39fd91636443ae8fe05e6cf19d2c67d8; passport_csrf_token_default=39fd91636443ae8fe05e6cf19d2c67d8; msToken=y-kqP6IkjdojMtUdSj3oxmHvunpHAlDAlbl8t7Ji_KrAwC43J5SYDeo2HZRBCsxR1reu78A5JUio_xgYcowcMbGGjQy-fnQjyoMKO9xVugLTqgRfT8bi; msToken=Uk1zRW_2ZVxxpYN451GdotKzq23Z7OWiOIw3bH-ZQQbtfCbkcM6b8sM3e-gQqDm2lMvBHWxt8NLYyKnofOcDtCe-29OGOjvmA4tq0VtGVkqz2UDvSdGY; msToken=6MlqHKH2eqAEZxmjrR31Gker622hReaQlJ39T9UezVllU_iDswNZ_K6NNmwkb4Ea2Um82ERXRUMk8Mx0A7ya4KPYM1JcxI4IU8d2M2CB5PF1XQ30-QO6; s_v_web_id=verify_lu73qhg7_NamwXueY_TAmN_40uu_BmSo_7qO2wbSshad1; odin_tt=b9f5c6c15351e8594e5330bf816670696f9106e969b3ef1d8188390a8ee6237caa5afca4f9b3286e470d04b5624ec5c3dc4480f85b92725b4e48645109f3186d; multi_sids=7350323007930729515%3A5fb888bbb6cc2eb4ffbe8724a256f676; cmpl_token=AgQQAPNSF-RO0rY_eXKIJV0S_GAOekMTP6LZYNCi9A; ttwid=1%7Ck1EX0WEHCsqqQDXo6HxN5WFTgZQ2ldThtOzWNdl1EcM%7C1711380500%7C7792f920f9a65870fd2043832a2fc663d5ba95e38e5a398d4124a8eae54744f6; sid_guard=5fb888bbb6cc2eb4ffbe8724a256f676%7C1711380500%7C15552000%7CSat%2C+21-Sep-2024+15%3A28%3A20+GMT; uid_tt=aa18edc1345f99c7bd0131307d5a81d01d895857af9c6e6f92bddfdae86a081f; uid_tt_ss=aa18edc1345f99c7bd0131307d5a81d01d895857af9c6e6f92bddfdae86a081f; sid_tt=5fb888bbb6cc2eb4ffbe8724a256f676; sessionid=5fb888bbb6cc2eb4ffbe8724a256f676; sessionid_ss=5fb888bbb6cc2eb4ffbe8724a256f676; sid_ucp_v1=1.0.0-KDZkNDU0MzUzZmUwZDdhZmEyY2RkZDFkMjBhNzkxZTkyOWYzNDExNDgKGAiriN6A0PrlgGYQlLCGsAYYsws4AUDrBxAEGgd1c2Vhc3Q1IiA1ZmI4ODhiYmI2Y2MyZWI0ZmZiZTg3MjRhMjU2ZjY3Ng; ssid_ucp_v1=1.0.0-KDZkNDU0MzUzZmUwZDdhZmEyY2RkZDFkMjBhNzkxZTkyOWYzNDExNDgKGAiriN6A0PrlgGYQlLCGsAYYsws4AUDrBxAEGgd1c2Vhc3Q1IiA1ZmI4ODhiYmI2Y2MyZWI0ZmZiZTg3MjRhMjU2ZjY3Ng; store-idc=useast5; store-country-code=us; store-country-code-src=uid; tt-target-idc=useast5; tt-target-idc-sign=hIVlkjRPVnpxmhnS1Yf1BCuvzgHLjigQehSfxh8wF0dav1LayZmbsUOcFT3sw9Gr_HRzTWbB8u1QkaThrb685i_u_4YPsA93YLGDmGCboevF91vUfGbWbCCAgI4YUTlFcOWIh88EBmAboPbpmjnRKxmKiTdUh1LVu-gdxdAqVxjmJ9WmHH-3FPStY6GDSLRNeu3MzOnbYiCW8NxPn3Mef4gd4oHdLAFDoPhIGOWXXg9bBGhPzb1WjWnxIf9gTWJCYonvGPygWbNRU4sM5R40JXHxRDG9Zo9u4esvBW-nVnfW-JpcrlUULLf4wSkvVlvzEzhLwTkZsE7YumAteEHyEXvTMS0uo9IImPK3BdHRo5o6AAASWDcJ45cTO5q_DfUXOVKU66u4zhAjueig7WVOiuYemitS9lB6BvCnBWyjmOGw03h1EWEOqcST84aw8FG7cxKHfoeusH8Nu5KRDGSvpCGg3sPZOBnfgOa17XBAo_t4Op_rTScDn_gknSZVrFca; msToken=gQgPVgfmmrNI20kzHfUMXv9HhJCgz9rCP3KQ0DSDzSgFy4VTW15XIBX7QGgYHJMlw6FmiFHLGzSv2DpQm7VPsik6L6mYgaAIru5yKI7O-j43YPm2cq91LhOES0aT0I9G1E2dBA==; last_login_method=email"
	cookies := ParseCookies(cookieStr)

	for _, cookie := range cookies {
		err := page.SetCookies([]*proto.NetworkCookieParam{{
			Name:  cookie.Name,
			Value: cookie.Value,
			// 你可以根据需要设置其他字段
			Domain: ".tiktok.com",
			Path:   "/",
		}})
		if err != nil {
			log.Fatalf("could not set cookies: %v", err)
		}
	}

	// 导航到目标页面
	page.MustNavigate("https://www.tiktok.com/")
	page.MustWaitLoad()

	// todo 找到搜索框，输入关键字，点击搜索按钮
	//<input placeholder="Search" name="q" type="search" autocomplete="off" role="combobox" aria-controls="" aria-label="Search" aria-expanded="false" aria-autocomplete="list" data-e2e="search-user-input" class="css-1geqepl-InputElement e14ntknm3" value="@fstar0685">
	ele := page.MustElement("input[placeholder='Search']")
	ele.MustInput("@camglocks300")

	time.Sleep(5 * time.Second)
	// button data-e2e="search-box-button"
	button := page.MustElement("button[data-e2e='search-box-button']")
	button.MustClick()

	time.Sleep(5 * time.Second)
	//aria-label="Watch in full screen"找到第一个视频
	video := page.MustElement("div[aria-label='Watch in full screen']")
	video.MustClick()

	// 等待视频播放完成
	time.Sleep(10 * time.Second)
	//aria-label="Reply"找到评论按钮
	reply := page.MustElement("span[aria-label='Reply']")
	reply.MustClick()

	time.Sleep(3 * time.Second)

	// namelist := []string{"imnotkj2", "._r3dthoughtz", "chetoisbroke", "b2ok3n", "jerry2fresh", "imnotkj2", "._r3dthoughtz", "chetoisbroke", "b2ok3n", "jerry2fresh", "jerry2fresh", "imnotkj2", "._r3dthoughtz", "chetoisbroke", "b2ok3n", "jerry2fresh"}
	count_int := 0
	return_list := flist

	for _, name := range flist {

		

		if count_int > 120 {
			fmt.Println("评论已经超过150个")
			break
		}
		// ! return_list 去掉name元素
		return_list = removeElement(return_list, name)
		func() {
			at := page.MustElement("div[data-e2e='comment-at-icon']")
			at.MustClick()

			//class="public-DraftEditorPlaceholder-root"找到评论框
			comment := page.MustElement(".public-DraftEditorPlaceholder-root")
			//输入评论
			//焦点定位到评论框
			comment.MustInput(name)

			time.Sleep(3 * time.Second)
			divcom := page.MustElement("div[data-e2e='comment-at-list']")
			divcom.MustClick()

			timeout := 5 * time.Second

			// 尝试在指定时间内查找元素
			count, err := page.Timeout(timeout).Element(`[class*="DivTextCount"]`)
			if err != nil {
				fmt.Println("元素不存在")
			} else {
				// 如果找到元素，则输出其文本内容
				fmt.Println(count.MustText())
				new_count := count.MustText()
				//64/150以/分割，取/前面的数字
				new_count = strings.Split(new_count, "/")[0]
				fmt.Println(new_count)
				//! 将字符串new_count变成int
				fmt.Sscanf(new_count, "%d", &count_int)
				

			}
		}()
	}
	fmt.Println("评论完成")
	//! class="css-1fqw03m-DivTextCount e1d90qbv5" 找到 class 包含DivTextCount字符的元素
	// comment2 := page.MustElement(".css-1fqw03m-DivTextCount")

	// comment2 := page.MustElement(".public-DraftEditorPlaceholder-root")
	// comment2.MustInput("I like it")
	// comment.MustInput("I like it")

	// //aria-label="Post"
	post := page.MustElement("div[aria-label='Post']")
	post.MustClick()
	page.MustWaitLoad()
	time.Sleep(150000 * time.Second)
	return return_list

	// time.Sleep(500000 * time.Second)
}


func removeElement(slice []string, element string) []string {
	newSlice := []string{}
	for _, item := range slice {
		if item != element {
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}