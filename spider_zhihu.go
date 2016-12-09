package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

/*************************************************************************************
	Global Variables
*********************************************************************************/
var (
	gCurCookies   []*http.Cookie
	gCurCookieJar *cookiejar.Jar
)

/****************************************************************************************
	Functions
****************************************************************************************/
func initAll() {
	gCurCookies = nil
	gCurCookieJar, _ = cookiejar.New(nil)
}

func getLogin() {
	zhihuURL := "https://www.zhihu.com"
	zhihuLoginURL := "https://www.zhihu.com/login/phone_num"
	zhihuProfile := "https://www.zhihu.com/settings/profile"

	httpReq, err := http.NewRequest("GET", zhihuURL, nil)
	if err != nil {
		log.Fatal("zhihu get error")
	}

	// POST表单及COOKIE
	formData := &url.Values{}
	formData.Set("_xsrf", "981e181c9cd21793b704013b50f619d1")
	formData.Set("password", "1qaz2wsx")
	formData.Set("captcha_type", "cn")
	formData.Set("phone_num", "13568801795")
	// 初始化请求
	loginReq, err := http.NewRequest("POST", zhihuLoginURL, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Fatal("Fatal error", err.Error())
		os.Exit(10)
	}

	// 存储GET返回的COOKIE
	getCookie := gCurCookieJar.Cookies(httpReq.URL)
	cookieLen := len(getCookie)
	for i := 0; i < cookieLen; i++ {
		loginReq.AddCookie(getCookie[i])
	}

	loginReq.Header.Add("Accept", "*/*")
	loginReq.Header.Add("Accept-Encoding", "gzip, deflate, br")
	loginReq.Header.Add("Accept-Language", "zh-CN,zh;q=0.8")
	loginReq.Header.Add("Connection", "keep-alive")
	loginReq.Header.Add("Content-Length", "231")
	loginReq.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	loginReq.Header.Add("Host", "www.zhihu.com")
	loginReq.Header.Add("Origin", "https://www.zhihu.com")
	loginReq.Header.Add("Referer", "https://www.zhihu.com/?next=%2Fsettings%2Fprofile")
	loginReq.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.19 Safari/537.36")
	loginReq.Header.Add("X-Requested-With", "XMLHttpRequest")
	loginReq.Header.Add("X-Xsrftoken", "ed764a74f91b864343569f2dec5651ae")
	// gCurCookieJar.SetCookies(httpReq.URL, getCookie)

	client := &http.Client{}
	loginResp, err := client.Do(loginReq)
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(11)
	}

	defer loginResp.Body.Close()
	if loginResp.StatusCode == 200 {
		fmt.Println("Login Success")
		loginCookie := loginResp.Cookies()

		httpProfile, err := http.NewRequest("GET", zhihuProfile, nil)
		if err != nil {
			log.Fatal("Fatal error", err.Error())
		}

		cookieLen = len(loginCookie)
		for i := 0; i < cookieLen; i++ {
			httpProfile.AddCookie(loginCookie[i])
		}

	}

	// resp, err := client.PostForm(zhihuLoginURL, *formData)
	// if err != nil {
	// 	log.Fatal("Login Error!")
	// }

	// 获取登录后COOKIE并打印
	// loginCookie := resp.Cookies()
}

func main() {
	fmt.Println("============ BEGIN ==============")
	initAll()
	getLogin()
	fmt.Println("============ END ================")
}

// func cookiesParsePrint(cookies []*http.Cookie) {
// 	cookieNum := len(cookies)
// 	fileName := "cookie.log"

// 	if _, err := os.Stat(fileName); os.IsNotExist(err) {
// 		os.Create(fileName)
// 	}

// 	fmt.Println("====  Cookie Parse  ====")
// 	handle, err := os.OpenFile(fileName, os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Fatal("File open error")
// 	}
// 	defer handle.Close()

// 	for i := 0; i < cookieNum; i++ {
// 		curV := cookies[i]

// 		io.WriteString(handle, "\r\n\r\n")
// 		io.WriteString(handle, "Name ===>\t"+curV.Name+"\r\n")
// 		io.WriteString(handle, "Value ===>\t"+curV.Value+"\r\n")
// 		io.WriteString(handle, "Path ===>\t"+curV.Path+"\r\n")
// 		io.WriteString(handle, "Domain ===>\t"+curV.Domain+"\r\n")

// 		time := curV.Expires.Format("2006-01-02 15:04:05")
// 		io.WriteString(handle, "Expires ===>\t"+time+"\r\n")

// 		io.WriteString(handle, "RawExpires ===>\t"+curV.RawExpires+"\r\n")
// 		io.WriteString(handle, "MaxAge ===>\t"+(string)(curV.MaxAge)+"\r\n")
// 		io.WriteString(handle, "Secure ===>\t"+strconv.FormatBool(curV.Secure)+"\r\n")
// 		io.WriteString(handle, "HttpOnly ===>\t"+strconv.FormatBool(curV.HttpOnly)+"\r\n")
// 		io.WriteString(handle, "Unparsed ===>\t"+strings.Join(curV.Unparsed, "-")+"\r\n")
// 		io.WriteString(handle, "<====================================================>\r\n")
// 	}
// }
