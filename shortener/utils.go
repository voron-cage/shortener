package shortener

import (
	"fmt"
	"net/url"
)

const (
	alphabet    = "ynAJfoSgdXHB5VasEMtcbPCr1uNZ4LG723ehWkvwYR6KpxjTm8iQUFqz9D"
	alphabetLen = uint64(len(alphabet))
)

func getShortURLFromInt(id uint64) string {
	var shortStr string
	for id > 0 {
		shortStr = fmt.Sprintf("%s%s", shortStr, string(alphabet[id%alphabetLen]))
		id /= alphabetLen
	}
	return shortStr
}

func buildAbsoluteURI(isHttps bool, host, urlPath string) string {
	scheme := "https"
	if !isHttps {
		scheme = "http"
	}
	u := url.URL{Scheme: scheme, Host: host, Path: urlPath}
	return u.String()
}
