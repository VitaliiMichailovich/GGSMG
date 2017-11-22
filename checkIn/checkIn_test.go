package checkIn

import (
	"testing"
	"errors"
)

var testCaseEmailFixer = []struct {
	email string
	want string
	err error
}{
	{"vitaliy@online.ua", "vitaliy@online.ua", nil},
	{"vitaliy", "vitaliy", errors.New("Invalid email format. Please correct your email adress and send again.")},
	{"vitaliy@an.on", "vitaliy@an.on", errors.New("Unresolvable host of email. Please correct your email adress and send again.")},
}

var testCaseDomainFixer = []struct {
	domain string
	want string
	err error
}{
	{"library.dp.ua", "http://library.dp.ua", nil},
	{"library.dp", "http://library.dp", nil},
	{"http://library.dp", "http://library.dp", nil},
	{"https://library.dp", "https://library.dp", nil},
	{"http://library/", "library/", errors.New("Invalid domain format. Please correct domain and send it again.")},
}

func TestEmailFixer(t *testing.T) {
	for _, tc := range testCaseEmailFixer {
		want, err := EmailFixer(tc.email)
		if err != nil {
			if err.Error() != tc.err.Error() {
				t.Fatalf("TC(%v) got wrong error \nGot %v, but want \n%v",
					tc.email, err.Error(), tc.err.Error())
			}
		}
		if want != tc.want {
			t.Fatalf("TC(%v) failed \nGot %v, but want \n%v",
				tc.email, want, tc.want)
		}
	}
}

func TestDomainFixer(t *testing.T) {
	for _, tc := range testCaseDomainFixer {
		want, err := DomainFixer(tc.domain)
		if err != nil {
			if err.Error() != tc.err.Error() {
				t.Fatalf("TC(%v) got wrong error \nGot %v, but want \n%v",
					tc.domain, err.Error(), tc.err.Error())
			}
		}
		if want != tc.want {
			t.Fatalf("TC(%v) failed \nGot %v, but want \n%v",
				tc.domain, want, tc.want)
		}
	}
}
