package app

import (
	"encoding/base64"
	"io/ioutil"
	"mime"
	"net/http"
	"strings"
)

var (
	// BasicRealm is header
	BasicRealm = "Authorization Required"
)

// BasicAuthFunc returns a Handler that authenticates via Basic Auth using the provided function.
// The function should return true for a valid username/password combination.
func BasicAuthFunc(authfn func(string, string, *http.Request) bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			auth := req.Header.Get("Authorization")
			if len(auth) < 6 || auth[:6] != "Basic " {
				unauthorized(res)
				return
			}
			b, err := base64.StdEncoding.DecodeString(auth[6:])
			if err != nil {
				unauthorized(res)
				return
			}
			tokens := strings.SplitN(string(b), ":", 2)
			if len(tokens) != 2 || !authfn(tokens[0], tokens[1], req) {
				unauthorized(res)
				return
			}
			next.ServeHTTP(res, req)
		})
	}
}

func unauthorized(res http.ResponseWriter) {
	res.Header().Set("WWW-Authenticate", "Basic realm=\""+BasicRealm+"\"")
	http.Error(res, "Not Authorized", http.StatusUnauthorized)
}

func init() {

	mime.AddExtensionType(".atom", "application/atom+xml")

	mime.AddExtensionType(".crx", "application/x-chrome-extension")
	mime.AddExtensionType(".eot", "application/vnd.ms-fontobject")
	mime.AddExtensionType(".ico", "image/x-icon")
	mime.AddExtensionType(".json", "application/json")

	mime.AddExtensionType(".m4v", "video/m4v")
	mime.AddExtensionType(".mp4", "video/mp4")
	mime.AddExtensionType(".ogg", "audio/ogg")
	mime.AddExtensionType(".oga", "audio/ogg")
	mime.AddExtensionType(".ogv", "video/ogg")
	mime.AddExtensionType(".otf", "font/opentype")
	mime.AddExtensionType(".rss", "application/rss+xml")

	mime.AddExtensionType(".svg", "images/svg+xml")
	mime.AddExtensionType(".swf", "application/x-shockwave-flash")
	mime.AddExtensionType(".ttf", "font/truetype")
	mime.AddExtensionType(".txt", "text/plain")
	mime.AddExtensionType(".unity3d", "application/vnd.unity")
	mime.AddExtensionType(".webm", "video/webm")
	mime.AddExtensionType(".webp", "image/webp")
	mime.AddExtensionType(".woff", "application/x-font-woff")
	mime.AddExtensionType(".xpi", "application/x-xpinstall")

	passBytes, _ := ioutil.ReadFile("password")
	pass := strings.TrimSpace(string(passBytes))
	www := http.FileServer(http.Dir("static"))
	auth := BasicAuthFunc(func(u, p string, r *http.Request) bool {
		return p == pass
	})
	www = auth(www)

	http.Handle("/", www)
}
