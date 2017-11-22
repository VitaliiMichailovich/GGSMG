package checkIn

import "testing"

var testCaseEmailFixer = []struct {
	email string
	want string
	err error
}{
	{"vitaliy@online.ua", "vitaliy@online.ua", nil},
}

var testCaseDomainFixer = []struct {
	domain string
	want string
	err error
}{
	{"library.dp.ua", "http://library.dp.ua", nil},
	{"library.dp", "http://library.dp", nil},
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
