package youtube

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	var cookies []*http.Cookie

	firstCookie := &http.Cookie{
		Name:   "PREF",
		Value:  "####", // replace with your value
		Path:   "/",
		Domain: ".youtube.com",
	}

	cookies = append(cookies, firstCookie)

	secondCookie := &http.Cookie{
		Name:   "VISITOR_INFO1_LIVE",
		Value:  "####", // replace with your value
		Path:   "/",
		Domain: ".youtube.com",
	}

	cookies = append(cookies, secondCookie)

	thirdCookie := &http.Cookie{
		Name:   "YSC",
		Value:  "#####", // replace with your value
		Path:   "/",
		Domain: ".youtube.com",
	}

	cookies = append(cookies, thirdCookie)

	fourthCookie := &http.Cookie{
		Name:   "LOGIN_INFO",
		Value:  "####", // replace with your value
		Path:   "/",
		Domain: ".youtube.com",
	}

	cookies = append(cookies, fourthCookie)

	// URL for cookies to remember
	cookieURL, err := url.Parse("https://www.youtube.com/results")

	if err != nil {
		log.Fatal(err)
	}

	jar.SetCookies(cookieURL, cookies)

	client := &http.Client{
		Jar: jar,
	}

	resp, err := client.Get("https://www.youtube.com/results?search_query=malvika")

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	err = ioutil.WriteFile("youtube.html", body, 0644)
	if err != nil {
		panic(err)
	}
}
