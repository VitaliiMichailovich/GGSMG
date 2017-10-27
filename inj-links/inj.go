//package parser
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"sync"
	"time"
)

type link struct {
	Count   int
	Checked bool
	Routine bool
}

var links = make(map[string]link)

var mutex = &sync.RWMutex{}

var wg = sync.WaitGroup{}

var okLinks = []string{
	"056.ua",
	"11channel.dp.ua",
	"34.ua",
	"9-channel.com",
	"ab-centr.com.ua",
	"adm.dniprorada.gov.ua",
	"adm.dp.gov.ua",
	"afdnipro-ua.org",
	"aiesec.dp.ua",
	"antigorod.com",
	"artvertep.com",
	"aucentredeurope.ucoz.ua",
	"austria.dp.ua",
	"babylonians.narod.ru",
	"bibliofilial.blogspot.com",
	"bibliomist.org",
	"blogs.tvi.ua",
	"bukvoid.com.ua",
	"catamaran.org.ua",
	"center-osvita.dp.ua",
	"cfscom.org",
	"chembaby.com",
	"cinemahall.org",
	"citex.info",
	"codeclubua.org",
	"conference.nbuv.gov.ua",
	"cow.com.ua",
	"day.kiev.ua",
	"day.kyiv.ua",
	"dcz.gov.ua",
	"dilovod.com.ua",
	"diss.rsl.ru",
	"dissercat.com",
	"dk.dp.ua",
	"dlitera.at.ua",
	"dnepr.info",
	"dneprfilial2.blogspot.com",
	"dnepropetrovsk.startnews.net",
	"dneprpost.com.ua",
	"dneprpost.info",
	"dnipro.legalaid.gov.ua",
	"dniprograd.org",
	"dniprolit.org.ua",
	"dniprorada.gov.ua",
	"dnpr.com.ua",
	"docs.google.com",
	"dp.ric.ua",
	"dpchas.com.ua",
	"drive.google.com",
	"duk.dp.ua",
	"dv-gazeta.info",
	"e-reading.link",
	"e-services.dp.gov.ua",
	"ecoleague.net",
	"ecomir.org.ua",
	"ex.ua",
	"exp21.com.ua",
	"facebook.com",
	"filial18n.blogspot.com",
	"filmweb.pl",
	"flot.dnepredu.com",
	"gazeta.lviv.ua",
	"goethe.de",
	"goo.gl",
	"gorod.dp.ua",
	"greenford.jimdo.com",
	"gurt.org.ua",
	"gymnasium33.dp.ua",
	"hromadske.dp.ua",
	"imfgoldmine.com",
	"innuleska.livejournal.com",
	"instagram.com",
	"internet-centr.dp.ua",
	"irbis.library.dp.ua",
	"irex.ua",
	"irsenas.livejournal.com",
	"justus.com.ua",
	"kaizenclub.com.ua",
	"klex.ru",
	"kmu.gov.ua",
	"knyhobachennia.com",
	"lenregion.dp.ua",
	"liberiliberati.org",
	"library.dp.ua",
	"library30topolya3.blogspot.com",
	"library4.dp.ua",
	"libraryfil15.blogspot.com",
	"mayapovolotskaya.blogspot.com",
	"mcenterdnepr.inf.ua",
	"md-eksperiment.org",
	"memorialmap.org",
	"misto.news",
	"most-dnepr.info",
	"naiu.org.ua",
	"news.ui.ua",
	"novocti.org.ua",
	"obljust.gov.ua",
	"olga-kai.livejournal.com",
	"osvita.mediasapiens.ua",
	"partyofregions.ua",
	"periscope.tv",
	"pobeda126.blogspot.com",
	"prof-nvf.org",
	"prostir.museum",
	"rename.dp.ua",
	"robertdawson.com",
	"royallib.ru",
	"rspp.dp.ua",
	"sam-sam.pp.ua",
	"schedule.nrcu.gov.ua",
	"sds.in.ua",
	"teleteatr.com.ua",
	"times.dp.ua",
	"ualit.org",
	"uglov.tvereza.info",
	"uk.wikipedia.org",
	"ukrainka.org.ua",
	"ukrposhta.ua",
	"ula.org.ua",
	"uodi.in.ua",
	"uondnepr.inf.ua",
	"varganshik.livejournal.com",
	"vesti-ukr.com",
	"youtu.be",
	"youtube.com",
	"zno.academia.in.ua",
	"zorya.org.ua"}

func Parser(domain string) (map[string]link, error) {
	mutex.Lock()
	links[domain] = link{Count: 1, Checked: false, Routine: false}
	mutex.Unlock()
	for {
		mutex.RLock()
		li := make(map[string]link, len(links))
		for key, value := range links {
			li[key] = value
		}
		mutex.RUnlock()
		var checked, runned, run int
		for k, v := range li {
			if v.Checked == false {
				if v.Routine == false {
					mutex.Lock()
					links[k] = link{Count: links[k].Count, Checked: links[k].Checked, Routine: true}
					mutex.Unlock()
					run++
					wg.Add(1)
					go linksWriter(domain, k)
				} else {
					runned++
				}
			} else {
				checked++
			}
		}
		time.Sleep(time.Second * 1)
		mutex.RLock()
		summ := len(links)
		mutex.RUnlock()
		//fmt.Printf("Total links:\t%v\t\tChecked:\t%v\t\tRunning:\t%v\t\tRuned:\t%v\n", summ, checked, runned, run)
		if summ == checked {
			break
		}
	}
	wg.Wait()
	return links, nil
}

func linksWriter(domain, linka string) {
	defer wg.Done()
	response, err := http.Get(linka)
	if err != nil {
		mutex.Lock()
		links[linka] = link{Checked: false, Count: links[linka].Count, Routine: false}
		mutex.Unlock()
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 501 {
		mutex.Lock()
		links[linka] = link{Checked: false, Count: links[linka].Count, Routine: false}
		mutex.Unlock()
		response.Body.Close()
		return
	}
	if response.StatusCode != http.StatusOK {
		mutex.Lock()
		links[linka] = link{Checked: true, Count: links[linka].Count, Routine: false}
		mutex.Unlock()
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
					ok, url := getHref(domain, linka, tagInfo)
					if !ok {
						continue top
					}
					mutex.Lock()
					links[url] = link{Count: links[url].Count + 1, Checked: links[url].Checked, Routine: links[url].Routine}
					mutex.Unlock()
				}
			}
			mutex.Lock()
			links[linka] = link{Checked: true, Count: links[linka].Count, Routine: false}
			mutex.Unlock()
		} else {

			mutex.Lock()
			delete(links, linka)
			mutex.Unlock()
		}
	}
}

func getHref(domain string, linka string, t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			if strings.Contains(a.Val, "@") {
				return false, ""
			}
			if a.Val == "#" {
				return false, ""
			}
			href = a.Val
			href = strings.Replace(href, "./", domain, 1)
			href = strings.Replace(href, "www.", "", 1)
			if strings.Contains(href, "#") {
				href = href[:strings.Index(href, "#")]
			}
			if len(href) < 4 {
				return false, ""
			}
			if href[0:4] != "http" {
				href = domain + href
			}
			if !strings.Contains(href, domain) {
				i := 0
				for _, v := range okLinks {
					if strings.Contains(href, v) {
						i++
					}
				}
				if i == 0 {
					fmt.Println(linka, href)
				}
				return false, ""
			}
			last4 := href[len(href)-4:]
			if last4 == ".pdf" ||
				last4 == ".doc" ||
				last4 == ".xls" ||
				last4 == ".txt" ||
				last4 == ".rtf" ||
				last4 == ".jpg" ||
				last4 == "jpeg" ||
				last4 == "docx" ||
				last4 == "xlsx" {
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
