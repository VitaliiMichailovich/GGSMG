package xmlgen

import (
	"github.com/VitaliiMichailovich/GGSMG/parser"
	"strconv"
	"os"
)

func StyleCreator(sitemap map[string]parser.Link, domain string) (string, string) {
	toFole := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\r\n<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:schemaLocation=\"http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd\">\r\n"
	for link, info := range sitemap {
		toFole += "<url><loc>" + link + "</loc><changefreq>daily</changefreq><priority>" + strconv.Itoa(info.Count) + "</priority></url>\r\n"
	}
	toFole += "</urlset>"
	return toFole, domain
}

func FileWriter(domain, mapa string) (error) {
	path := "client/"+domain
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(path+"/sitemap.xml"); !os.IsNotExist(err) {
		err = os.Remove(path+"/sitemap.xml")
		if err != nil {
			return err
		}
	}
	file, err := os.Create(path+"/sitemap.xml")
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(mapa)
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}
	return nil
}
