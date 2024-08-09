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

func Jump() {

	url := launcher.New().Headless(false).Proxy("http://127.0.0.1:10809").MustLaunch()
	// 通过URL连接到浏览器
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	

	page := browser.MustPage("https://www.tiktok.com/@tiktok")
	page.MustWaitLoad()

	// 解析并设置Cookies
	cookieStr := "tt_csrf_token=R7cZvLyO-Y5Zeh7nIPGbz2JrJFKnJxARwxlQ; passport_csrf_token=8f5d7731d6bc4c5921e7cc56721fb1d1; passport_csrf_token_default=8f5d7731d6bc4c5921e7cc56721fb1d1; msToken=I3_XWzs-c1CdCYzIqdqQ5oDm8kZcndQtQ1pVGEOY0w0G-TW5eIgO_76U233--vqrqYZUfzs_dH6_B7s1i8uxlu4T6E2452nJty5-Ixl9XvU-ELY1Kj6u3DdrgrHQnM48c3Kl_w==; s_v_web_id=verify_lueoywwc_urqi4eA9_kLMe_4LNO_9Qlc_Jzk4b2JbnanO; odin_tt=78ebeda6b50f6514c016cc1a6aee656d3c7c338e0dab46fc58f1e5f115a04263ea7b0e95ad512f2a6ce88065e4d1cb59d97221d05fc8740adc7d359db36bf0fe; multi_sids=7352294063289713707%3A8ac09a124680d217e1a1c102e9417549; cmpl_token=AgQQAPNSF-RO0rY2dtieqh0S_GkBa7Ib_4fZYNDPFA; ttwid=1%7COzWIeYRsRlcfE8WZLm9kLGDD9DMghfS94ebY1Jk2dHs%7C1711839417%7Cfbe6225ba16c49e2bc8d9fa241a5767d80226d12281d3ff84d27d24fed084f35; sid_guard=8ac09a124680d217e1a1c102e9417549%7C1711839417%7C15552000%7CThu%2C+26-Sep-2024+22%3A56%3A57+GMT; uid_tt=35b53e0d1922761541624da4a8db2c6bc5efa1cca46e85d81bd96bafba7a8bda; uid_tt_ss=35b53e0d1922761541624da4a8db2c6bc5efa1cca46e85d81bd96bafba7a8bda; sid_tt=8ac09a124680d217e1a1c102e9417549; sessionid=8ac09a124680d217e1a1c102e9417549; sessionid_ss=8ac09a124680d217e1a1c102e9417549; sid_ucp_v1=1.0.0-KGYwY2Y5ZmRiNGVjZTZhNWU2Y2U4ZDRjMjY1ZjYzMDVjZmQ0ZGUxZmYKGAiriOO08Y-mhGYQubGisAYYsws4AUDrBxAEGgd1c2Vhc3Q1IiA4YWMwOWExMjQ2ODBkMjE3ZTFhMWMxMDJlOTQxNzU0OQ; ssid_ucp_v1=1.0.0-KGYwY2Y5ZmRiNGVjZTZhNWU2Y2U4ZDRjMjY1ZjYzMDVjZmQ0ZGUxZmYKGAiriOO08Y-mhGYQubGisAYYsws4AUDrBxAEGgd1c2Vhc3Q1IiA4YWMwOWExMjQ2ODBkMjE3ZTFhMWMxMDJlOTQxNzU0OQ; store-idc=useast5; store-country-code=us; store-country-code-src=uid; tt-target-idc=useast8; tt-target-idc-sign=1b695kncwTOLxLJiblZQYEP_cQTgqPY9MU3AD-pjOZdtR6dcqppf_s-VHlBqIT64y5oKWBZtJGO1RZod9k5aI6pzSM4mCyBEsr3iVhj0KjaT2i2aJb0VeP3EL_6tDItSvMdlKWvPPGRopJDYW45_1WtKBFZDfrdfvDKlXd9-TPgs67MrOpLwg7dcmbBYprSmEofwP5clVkvrXrMxbkZz5sc9IqdXpW0l93BfENGrpwdiTJA57xcCHji1nj125a2puwtmB-rnaV3KOqoWwwlvfns_uy1KfjtQFwSHdPb2erXEzEA_9-wg_UzlaJhHtVI9KG1rr6aETUoCChV-jAhfiZoz8s1NzvcRUO-WFHkj8UlVd_7w642wnjIUWxcC8dIFLdVsvVaxz4H69jTODFohsOuf5Y7ieYeG2PjjILLdCvDBzA1FjXYaFUUobCvD46CwbF6vjK99kOajpGe4OZYR8ZPs8ANZfiELTWBmt_3pc5u3JrdKG33lHKamO_vrOrW5; msToken=4XgA_IoGwiQTMp7Ue6m4RyIaXbFdeqt_08ZO4T5_6TvC5VfM7IRYXZjylue6gdIQC7913FHmA41ua1bmBzzwSB6JaSUe1UXq6m_uHdKKN7moMo23IroVNYge0WfyvOyxg4tmdA==; last_login_method=email"
	cookies := parseCookies(cookieStr)

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
	page.MustNavigate("https://www.tiktok.com/@tiktok")
	page.MustWaitLoad()

	time.Sleep(500 * time.Second)
}

// parseCookies 是一个帮助函数，用来将Cookies字符串解析为Rod的Cookie对象
func parseCookies(cookieStr string) []*proto.NetworkCookie {
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