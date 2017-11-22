package xmlgen

import (
	"testing"
	"os"
	"io"
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