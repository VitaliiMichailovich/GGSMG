package xmlgen

import (
	"testing"
	"os"
	"io"
	"github.com/VitaliiMichailovich/GGSMG/parser"
)

var testCases = []struct {
	doamin string
	mapa string
	expect bool
}{
	{"lib.ua", "text", true},
	{"liba.ua", "libatext", true},
	{"library.ua", "libtextua", true},
	{"library.dp.ua", "lib.ua", true},
	{"lib.ua", "lib.ua", true},
}

var testCase_StyleCreator = []struct{
	sitemap map[string]parser.Link
	domain string
	expect_toFile string
	expect_domain string
}{
	{map[string]parser.Link{
		"test" : parser.Link{Count:1,Checked:true,Routine:true},
	}, "http://library.dp.ua/", "<?xml version=\"1.0\" encoding=\"UTF-8\"?>" +
		"\r\n<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:schemaLocation=\"http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd\">" +
		"\r\n<url><loc>test</loc><changefreq>daily</changefreq><priority>1</priority></url>" +
		"\r\n</urlset>", "http://library.dp.ua/"},
}
func TestStyleCreator(t *testing.T) {
	for _, tc := range testCase_StyleCreator {
		//case1["www"] = parser.Link{Count:1,Checked:true,Routine:true}
		toFile, domain := StyleCreator(tc.sitemap, tc.domain)
		if toFile != tc.expect_toFile {
			t.Fatalf("TC(%v, %v) got error \nGot %v, but want \n%v",
				tc.sitemap, tc.domain, toFile, tc.expect_toFile)
		}
		if domain != tc.expect_domain {
			t.Fatalf("TC(%v, %v) got error \nGot %v, but want \n%v",
				tc.sitemap, tc.domain, domain, tc.expect_domain)
		}
	}
}

func TestFileWriter (t *testing.T) {
	for _, tc := range testCases {
		err := FileWriter(tc.doamin, tc.mapa)
		if err != nil {
			t.Fatalf("TC(%v, %v) got error \nGot %v, but want \n%v",
				tc.doamin, tc.mapa, err.Error(), tc.expect)
		}
		file, err := os.OpenFile("client/"+tc.doamin+"/sitemap.xml", os.O_RDWR, 0644)
		if err != nil {
			t.Fatalf("TC(%v, %v) got error \nGot %v, but want \n%v",
				tc.doamin, tc.mapa, err.Error(), tc.expect)
		}
		stat, _ := file.Stat()
		var text = make([]byte, stat.Size())
		for {
			_, err = file.Read(text)
			if err == io.EOF {
				break
			}
			if err != nil && err != io.EOF {
				if err != nil {
					t.Fatalf("TC(%v, %v) got error \nGot %v, but want \n%v",
						tc.doamin, tc.mapa, err.Error(), tc.expect)
				}
			}
		}
		if string(text) != tc.mapa {
			t.Fatalf("TC(%v, %v) got error \nGot %v, but want \n%v",
				tc.doamin, tc.mapa, err.Error(), tc.expect)
		}
		file.Close()
	}
	err := os.RemoveAll("client")
	if err != nil {
		t.Fatalf("Cleaning error. %v", err.Error())
	}
}