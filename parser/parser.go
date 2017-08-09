package parser

import (
	"html"
	"strings"
	"time"
	"net/http"
	"fmt"
	"errors"
)

type Link struct {
	Count   int
	Checked bool
}

//var Links = make(map[string]Link)
//var trues int

func Parser (uri string) (links map[string]Link, err error) {
	linksChannel := make(chan string)
	linksChannel <- uri
	fmt.Println("I made a channel.")
	links = make(map[string]Link)
	fmt.Println("I made a variable with Links.")

	go func() {

	}()

	err = errors.New("AAA")

	return
}

//func getHref(t html.Token, link string) (ok bool, href string) {
//	for _, a := range t.Attr {
//		if a.Key == "href" {
//			if strings.Contains(a.Val, "@") {
//				return false, ""
//			}
//			if a.Val == "#" {
//				return false, ""
//			}
//			href = a.Val
//			href = strings.Replace(href, "./", link, 1)
//			href = strings.Replace(href, "www.", "", 1)
//			if strings.Contains(href, "#") {
//				href = href[0:strings.Index(href, "#")]
//			}
//			if len(href) < 4 {
//				return false, ""
//			}
//			if href[0:4] != "http" {
//				href = link + href
//			}
//			if !strings.Contains(href, link) {
//				return false, ""
//			}
//			last4 := href[len(href)-4:]
//			if last4 == ".pdf" ||
//				last4 == ".doc" ||
//				last4 == ".xls" ||
//				last4 == ".txt" ||
//				last4 == ".rtf" ||
//				last4 == ".jpg" ||
//				last4 == ".jpeg" ||
//				last4 == ".docx" ||
//				last4 == ".xlsx" {
//				return false, ""
//			}
//			ok = true
//		}
//	}
//	return ok, href
//}
//
//func ParserA(uri string) error {
//	var link string
//	for k, v := range Links {
//		if v.Checked == false {
//			link = k
//			break
//		}
//	}
//afterSleep:
//	time.Sleep(1 * time.Millisecond)
//	response, err := http.Get(link)
//	if response.StatusCode == 503 {
//		time.Sleep(10 * time.Second)
//		//fmt.Println("503 : ", link)
//		goto afterSleep
//	}
//	if err == nil && response.StatusCode == 200 {
//		defer response.Body.Close()
//		if response.Header["Content-Type"][0] == "text/html" {
//			tokenizer := html.NewTokenizer(response.Body)
//		top:
//			for {
//				linkInfo := tokenizer.Next()
//				switch {
//				case linkInfo == html.ErrorToken:
//					break top
//				case linkInfo == html.StartTagToken:
//					tagInfo := tokenizer.Token()
//					isAnchor := tagInfo.Data == "a"
//					if !isAnchor {
//						continue top
//					}
//					ok, url := getHref(tagInfo, uri)
//					if !ok {
//						fmt.Println(len(Links), "/", trues, " - ", response.StatusCode, " - ", link, " - ", tagInfo)
//						continue top
//					}
//					Links[url] = Link{Count: Links[url].Count + 1, Checked: Links[url].Checked}
//				}
//			}
//		}
//	}
//	Links[link] = Link{Checked: true, Count: Links[link].Count}
//	trues++
//	if trues%1000 == 0 {
//		fmt.Println(len(Links), "/", trues, " - ", response.StatusCode, " - ", link)
//	}
//
//	if len(Links) != trues {
//		ParserA(uri)
//	}
//	return nil
//}
