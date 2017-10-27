package checkIn

import (
	"errors"
	"net"
	"regexp"
	"strings"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var domainRegexp = regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z ]{2,3})$`)

func EmailFixer(email string) (string, error) {
	if !emailRegexp.MatchString(email) {
		return email, errors.New("invalid format")
	}

	_, host := split(email)
	_, err := net.LookupMX(host)
	if err != nil {
		return email, errors.New("unresolvable host")
	}

	return email, nil
}

func split(email string) (account, host string) {
	i := strings.LastIndexByte(email, '@')
	account = email[:i]
	host = email[i+1:]
	return
}

func DomainFixer(domain string) (string, error) {
	prefix := "http://"
	if strings.Contains(domain, prefix) {
		domain = strings.Replace(domain, prefix, "", 1)
	}
	if strings.Contains(domain, "https://") {
		prefix = "https://"
		domain = strings.Replace(domain, "https://", "", 1)
	}
	if !domainRegexp.MatchString(domain) {
		return domain, errors.New("invalid format")
	}
	return prefix + domain, nil
}
