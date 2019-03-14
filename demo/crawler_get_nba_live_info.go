package demo

import (
	"fmt"
	"time"
	"net/http"
	"crypto/tls"
	"github.com/PuerkitoBio/goquery"
)

//爬虫小栗子
func Get_nba_live_info() {
	for {
		select {
		case <- time.After(time.Second * 1):
			tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}
			client := &http.Client{Transport: tr}
			//访问虎扑
			req, _ := http.NewRequest("GET", "https://nba.hupu.com/games/playbyplay/157018", nil)
			//req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")
			//req.Header.Add("Upgrade-Insecure-Requests", "2")
			result, _ := client.Do(req)
			defer result.Body.Close()
			doc, _ := goquery.NewDocumentFromReader(result.Body)
			s := doc.Find("div[class=\"table_list_live playbyplay_td table_overflow\"]").Find("table").Find("tr").Eq(0)
			fmt.Print(s.Text())
		}
	}
}
