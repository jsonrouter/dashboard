package generate

import (
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/base64"
)

const (
	CONST_PREFIX = "angular/dist/angular/"
	CONST_FILE_TEMPLATE = `package static

import (
	"strings"
	"encoding/base64"
	//
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/tree"
)

func FileList() []string {
	return []string{"%s"}
}

func (self *Static) Dashboard(node *tree.Node) {

%s

	for _, filename := range FileList() {

		node.Add(filename, "$filename").GET(
			func (req http.Request) *http.Status {

				var contentType string
				resource := req.Param("$filename").(string)
				ext := strings.Split(resource, ".")[1]
				switch ext {
					case "ico":
						contentType = "image/png"
					case "html":
						contentType = "text/html"
					case "js":
						contentType = "text/javascript"
					case "css":
						contentType = "text/css"
					default:
						panic("MIME TYPE NOT RECOGNISED: "+ext)
				}

				req.SetHeader("Content-Type", contentType)

%s

				return req.Respond(
					self.files[resource].Cache,
				)
			},
		).Description(
			"Serves the dashboard file: " + filename,
		)
	}
}

const(
%s
)

`
)

func Generate() {

	fileList := []string{
		"favicon.ico",
		"index.html",
		"main.js",
		"polyfills.js",
		"runtime.js",
		"styles.css",
	}
	files := []string{}
	decoders := []string{}

	for _, filename := range fileList {

		b, err := ioutil.ReadFile(CONST_PREFIX + filename)
		if err != nil {
			panic(err)
		}

		decoders = append(
			decoders,
			fmt.Sprintf(
`
	if newFile, err := base64.StdEncoding.DecodeString(CONST_SRC_%s); err != nil {
		panic(err)
	} else {
		self.files["%s"] = &File{newFile}
	}
`,
				strings.Replace(strings.ToUpper(filename), ".", "_", -1),
				filename,
			),
		)

		files = append(
			files,
			fmt.Sprintf(
				"CONST_SRC_%s = \"%s\"\n",
				strings.Replace(strings.ToUpper(filename), ".", "_", -1),
				base64.StdEncoding.EncodeToString(b),
			),
		)

	}

	fileBytes := []byte(
		fmt.Sprintf(
			CONST_FILE_TEMPLATE,
			strings.Join(fileList, `", "`),
			strings.Join(decoders, ""),
`req.Log().Debugf(
				"Serving Dashboard file: %s with MimeType: %s",
				resource,
				req.GetHeader("Content-Type"),
			)`,
			strings.Join(files, ""),
		),
	)

	err := ioutil.WriteFile("gen/const.go", fileBytes, 0777)
	if err != nil {
		panic(err)
	}
}