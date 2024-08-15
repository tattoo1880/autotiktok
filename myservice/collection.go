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

func Collection() []string {

	url := launcher.New().Headless(false).Proxy("http://127.0.0.1:10809").MustLaunch()
	// 通过URL连接到浏览器
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	

	page := browser.MustPage("https://www.tiktok.com/")
	page.MustWaitLoad()

	// 解析并设置Cookies
	cookieStr := "tt_csrf_token=qXrzSrUv-aWSDTqnzv5sUpw_hnybECrMOTUA; passport_csrf_token=239d80f4ec0de9bab0886659f2ef2faf; passport_csrf_token_default=239d80f4ec0de9bab0886659f2ef2faf; msToken=b_PPWgF0QBwKolkT9Kcm9HolrEk2YkQr7BmLIqyeNo5O8jB3YlLoZWR1IpCwSCI66p-iD02OYLxx42JpjjaV-oS13xK1NGgZ9idB9t5RlKbH-f1PNokDGo_j-EGy8En6n3iK3g==; s_v_web_id=verify_lu9xfchv_BqvHi8mM_4aqw_44HJ_Bv19_YsFTaGaRtZsD; odin_tt=aa9232e7cb0c892f72042d23743285fc20a62e4700bee51709c13225b51eba8b52e606b08155320a1d12555ea313d2d23308f29a4a0b0b676379bb573ed78c21; multi_sids=7351056577607304234%3A58abef3555829460f97818bcba4769cb; cmpl_token=AgQQAPNSF-RO0rY63aUIJ90T_GWr0rmaf4TZYNDd6g; ttwid=1%7CG8zuR9TOPT0xyYUA5Cxl_V29Z4RazkyiK3AptTR18bA%7C1711551293%7C1e3f3c475cba75459d027ad6287f99c1b526e5061bc65399cb359ebff1e65ae0; sid_guard=58abef3555829460f97818bcba4769cb%7C1711551293%7C15552000%7CMon%2C+23-Sep-2024+14%3A54%3A53+GMT; uid_tt=4dba0afab0e81f22227e43c84eaae97421af3e7c38db97e077803ff2587d2ebd; uid_tt_ss=4dba0afab0e81f22227e43c84eaae97421af3e7c38db97e077803ff2587d2ebd; sid_tt=58abef3555829460f97818bcba4769cb; sessionid=58abef3555829460f97818bcba4769cb; sessionid_ss=58abef3555829460f97818bcba4769cb; sid_ucp_v1=1.0.0-KDM2YTZhNTk2YzBmM2JlYmIyZmE4NTEwOWE2NTNmOTM2YWIyZjc0YWEKGAiqiNSAqOCMgmYQveaQsAYYsws4AUDrBxAEGgd1c2Vhc3Q1IiA1OGFiZWYzNTU1ODI5NDYwZjk3ODE4YmNiYTQ3NjljYg; ssid_ucp_v1=1.0.0-KDM2YTZhNTk2YzBmM2JlYmIyZmE4NTEwOWE2NTNmOTM2YWIyZjc0YWEKGAiqiNSAqOCMgmYQveaQsAYYsws4AUDrBxAEGgd1c2Vhc3Q1IiA1OGFiZWYzNTU1ODI5NDYwZjk3ODE4YmNiYTQ3NjljYg; store-idc=useast5; store-country-code=us; store-country-code-src=uid; tt-target-idc=useast5; tt-target-idc-sign=TafllyQkdEBRsAP5DzcUfGZXLeTKJ_gxjQldNiGZ_w9JHq_xdbWUiGl8Ii0Xl6IbS6_zWsFN4wybHZOAVJZQqAtI8w6Fb2SZFecrux34awDdpERK6eb2fjh1OpPnqqEEobP-uupWmdj_-IWfKLChRyXIchWd-iUEL4HhYSfEzSA_A_QhyBEWubP8fHKD0jc_zV1V3ge0y0S_eQlOcar6RrzQQ0hlMigsKu31V4LkSu9Boxrpt7kxO9BcmOAqgo3eruknwv3DjY2wmBR1lHkOcVmkNhEbiXh1JDJA0rb_M0ubTLiwZZD4jOrK4ZUlGRPbe9irS-ERnE6_ZmrmSs0dJwGeuJvrkzS_8H-cGdiW61XJMc4_6z3rQp8WX3PF5AC6RuutcwK6UEys05ttEiP-RSemmm5KzcjgIyin2yevLewo-zlP5jMEowKXm3ASy5EuhffjOkhlHPmGfArQkkGKBcvGPx025JYFqFiD_CFY4aEUA-33YZRM_1SSPkx8mydO; msToken=WRHd3EKSh3d0NhXFfehQPPntQVMOIQ83tFxfDYr5IxcRX4qqnuuHc203P3PFRNXszezwkV66bobYGe028QFqdp2WN6XEvUQlZ95j2tXeYmqyHe9wr24ru0D57t7bJXFWM3L8cw==; last_login_method=email"
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

	//! 从终端获取输入要采集用户的用户名
	// fmt.Println("请输入要采集的用户的用户名：")
	var username string
	// fmt.Scanln(&username)
	username = "camglocks300"
	fmt.Println("你输入的用户名是：", username)

	page.MustNavigate("https://www.tiktok.com/@" + username)


	//! 等待页面加载完成
	page.MustWaitLoad()
	//! 找到data-e2e="followers" 的span
	followers := page.MustElement("span[data-e2e='followers']")
	//!click
	followers.MustClick()
	//! 等待页面加载完成
	page.MustWaitLoad()


	menu := page.MustElement("div[data-e2e='follow-info-popup']")

	if menu == nil {
		fmt.Println("没有找到关注者列表")
		return nil
	}
	box := menu.MustShape().Box()

	if box == nil {
		fmt.Println("没有找到关注者列表")
		return nil
	}


	centerX := (box.X + box.Width) / 2
	centerY := (box.Y + box.Height) / 2

	// 将鼠标移动到元素的中间
	page.Mouse.MustMoveTo(centerX, centerY)
	// 等待移动完成
	time.Sleep(1 * time.Second)

	
	// 启用Fetch事件监听
	err := proto.FetchEnable{}.Call(page)
	if err != nil {
		log.Fatalf("could not enable fetch: %v", err)
	}
	// 模拟向下滚动
	for i := 0; i < 10; i++ {
		page.Mouse.Scroll(0, 900, 5)
		time.Sleep(3 * time.Second)
	}
	
	// targetSubstring := "/api/user/list" // 替换为你要查找的URL子字符串
	// page.EachEvent(func(e *proto.FetchRequestPaused) {
	// 	if strings.Contains(e.Request.URL, targetSubstring) {
	// 		response, err := proto.FetchGetResponseBody{
	// 			RequestID: e.RequestID,
	// 		}.Call(page)
	// 		if err != nil {
	// 			log.Printf("无法获取响应体: %v", err)
	// 			return
	// 		}
	// 		fmt.Printf("截获的JSON数据: %s\n", response.Body)
	// 		// 继续处理请求
	// 		proto.FetchContinueRequest{
	// 			RequestID: e.RequestID,
	// 		}.Call(page)
	// 	}
	// })()

	// 
	flist := []string{}
	followerElements := page.MustElement("div[data-e2e='follow-info-popup']").MustElements("li a")

	for _, follower := range followerElements {
		hrefPtr, err := follower.Attribute("href")
		if err != nil {
			log.Printf("无法获取 href 属性: %v", err)
			continue
		}
		if hrefPtr == nil {
			fmt.Println("没有找到 href 属性")
			continue
		}
		href := *hrefPtr // 解引用
		fmt.Printf("关注者链接: %s\n", href)
		new_href := strings.Split(href, "/@")[1]
		flist = append(flist, new_href)
	}
	fmt.Println(flist)
	time.Sleep(10 * time.Second)

	

	
	

	return flist
}

// parseCookies 是一个帮助函数，用来将Cookies字符串解析为Rod的Cookie对象
