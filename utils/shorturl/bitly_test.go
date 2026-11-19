package shorturl

import "testing"

func TestShortURL(t *testing.T) {
	ret, err := ShortURL("token", "https://saas-site-backoffice-web.ogweb.org/dashboard")
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
