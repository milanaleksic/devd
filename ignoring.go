package devd

import "path"

var ignoredFiles map[string]bool

func init() {
	ignoredFiles = make(map[string]bool)
	ignoredFiles[".css"] = true
	ignoredFiles[".js"] = true
	ignoredFiles[".woff2"] = true
	ignoredFiles[".ttf"] = true
	ignoredFiles[".png"] = true
	ignoredFiles[".ico"] = true
}

func IsIgnored(dpath string) bool {
	_, ok := ignoredFiles[path.Ext(dpath)]
	return ok
}
