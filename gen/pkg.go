package static

import (
	"encoding/base64"
)

type Static struct {
	Favicon_ico *File
	Index_html *File
	Main_js *File
	Pollyfills_js *File
	Runtime_js *File
	Styles_css *File
}

type File struct {
	Cache []byte
}

func NewStatic(path string) *Static {

	f := func (encoded string) *File {

		c, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			panic(err)
		}

		return &File{
			Cache: c,
		}
	}

	return &Static{
		Favicon_ico: f(CONST_SRC_FAVICON_ICO),
		Index_html: f(CONST_SRC_INDEX_HTML),
		Main_js: f(CONST_SRC_MAIN_JS),
		Pollyfills_js: f(CONST_SRC_POLYFILLS_JS),
		Runtime_js: f(CONST_SRC_RUNTIME_JS),
		Styles_css: f(CONST_SRC_STYLES_CSS),
	}
}
