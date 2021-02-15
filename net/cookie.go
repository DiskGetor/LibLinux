package net

import (
	"net/http/cookiejar"
	"net/url"
)

func (*ObjNet) IsCookieInJar(jar *cookiejar.Jar, cookieName, Host string) (ok bool) {
	URL, err := url.Parse(Host)
	if !W.CheckErr(err) {
		return
	}
	for _, v := range jar.Cookies(URL) {
		if v.Name == cookieName {
			W.LogSuccess("", v)
			return
		}
	}
	return false
}
