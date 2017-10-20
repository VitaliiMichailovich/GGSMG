//package parser
package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Link struct {
	Count   int
	Checked bool
	Routine bool
}

var Links = make(map[string]Link)

var mutex = &sync.Mutex{}

var wg = sync.WaitGroup{}

//var ch = make(chan struct{}, 1000)

func Parser(uri string) (links map[string]Link, err error) {
	mutex.Lock()
	Links[uri] = Link{Count: 1, Checked: false, Routine: false}
	mutex.Unlock()
	for {
		mutex.Lock()
		li := make(map[string]Link, len(Links))
		for key, value := range Links {
			li[key] = value
		}
		mutex.Unlock()
		var checked, runned, run int
		for k, v := range li {
			if v.Checked == false {
				if v.Routine == false {
					mutex.Lock()
					Links[k] = Link{Count: Links[k].Count, Checked: Links[k].Checked, Routine: true}
					mutex.Unlock()
					run++
					wg.Add(1)
					//ch <- struct{}{}
					go linksWriter(k, uri)
				} else {
					runned++
				}
			} else {
				checked++
			}
		}
		time.Sleep(time.Second * 1)
		mutex.Lock()
		summ := len(Links)
		mutex.Unlock()
		fmt.Printf("Total links:\t%v\t\tChecked:\t%v\t\tRunning:\t%v\t\tRuned:\t%v\n", summ, checked, runned, run)
		if summ == 3992 && checked == 3590 && runned == 402 && run == 0 {
			break
		}
		if summ == checked-1 {
			mutex.Lock()
			fmt.Println("Links len 0 : " + strconv.Itoa(len(Links)))
			mutex.Unlock()
			time.Sleep(time.Second * 2)
		}
		if summ == checked {
			mutex.Lock()
			fmt.Println("Links len 1 : " + strconv.Itoa(len(Links)))
			mutex.Unlock()
			break
		}
	}
	wg.Wait()
	mutex.Lock()
	fmt.Println("Links len 2 : " + strconv.Itoa(len(Links)))
	for k, v := range Links {
		fmt.Println(k, v)
	}
	links = make(map[string]Link)
	err = errors.New("AAA")
	mutex.Unlock()
	return
}

func linksWriter(link, uri string) {
	defer wg.Done()
	response, err := http.Get(link)
	defer recover()
	if err != nil {
		fmt.Println(err.Error())
		mutex.Lock()
		Links[link] = Link{Checked: false, Count: Links[link].Count, Routine: false}
		mutex.Unlock()

		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 501 {
		mutex.Lock()
		Links[link] = Link{Checked: false, Count: Links[link].Count, Routine: false}
		mutex.Unlock()
		response.Body.Close()
		return
	}
	if response.StatusCode != http.StatusOK {
		mutex.Lock()
		Links[link] = Link{Checked: true, Count: Links[link].Count, Routine: false}
		mutex.Unlock()
		fmt.Println("Not Status OK -------------------------- ", response.StatusCode, link)
		response.Body.Close()
		return
	}
	if response.StatusCode == 200 {
		if response.Header["Content-Type"][0] == "text/html" {
			tokenizer := html.NewTokenizer(response.Body)
		top:
			for {
				linkInfo := tokenizer.Next()
				switch {
				case linkInfo == html.ErrorToken:
					break top
				case linkInfo == html.StartTagToken:
					tagInfo := tokenizer.Token()
					isAnchor := tagInfo.Data == "a"
					if !isAnchor {
						continue top
					}
					ok, url := getHref(tagInfo, uri)
					if !ok {
						continue top
					}
					mutex.Lock()
					Links[url] = Link{Count: Links[url].Count + 1, Checked: Links[url].Checked, Routine: Links[url].Routine}
					mutex.Unlock()
				}
			}
			mutex.Lock()
			Links[link] = Link{Checked: true, Count: Links[link].Count, Routine: false}
			mutex.Unlock()
		}
	}
}

func getHref(t html.Token, uri string) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			if strings.Contains(a.Val, "@") {
				return false, ""
			}
			if a.Val == "#" {
				return false, ""
			}
			href = a.Val
			href = strings.Replace(href, "./", uri, 1)
			href = strings.Replace(href, "www.", "", 1)
			if strings.Contains(href, "#") {
				href = href[:strings.Index(href, "#")]
			}
			if len(href) < 4 {
				return false, ""
			}
			if href[0:4] != "http" {
				href = uri + href
			}
			if !strings.Contains(href, uri) {
				return false, ""
			}
			last4 := href[len(href)-4:]
			if last4 == ".pdf" ||
				last4 == ".doc" ||
				last4 == ".xls" ||
				last4 == ".txt" ||
				last4 == ".rtf" ||
				last4 == ".jpg" ||
				last4 == ".jpeg" ||
				last4 == ".docx" ||
				last4 == ".xlsx" {
				return false, ""
			}
			ok = true
		}
	}
	return ok, href
}

func main() {
	Parser("http://library.dp.ua/")
}

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
