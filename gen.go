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

	node.SetHeaders(
		map[string]interface{}{
			"Access-Control-Allow-Headers": "Authorization,Content-Type",
			"Access-Control-Allow-Origin": "*",
		},
	)

%s

	for _, filename := range FileList() {

		node.Add(filename, "$filename").GET(
			func (req http.Request) *http.Status {

				var contentType string
				resource := req.Param("$filename").(string)
				ext := strings.Split(resource, ".")[1]
				switch ext {
					case "png":
						contentType = "image/png"
					case "ico":
						contentType = "image/png"
					case "html":
						contentType = "text/html"
					case "js":
						contentType = "text/javascript"
					case "json":
						contentType = "application/json"
					case "css":
						contentType = "text/css"
					default:
						panic("MIME TYPE NOT RECOGNISED: "+ext)
				}

				req.SetResponseHeader("Content-Type", contentType)

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
		"manifest.json",
		"ngsw.json",
		"ngsw-worker.js",
		"safety-worker.js",
		"worker-basic.min.js",
		//
		"icon-128x128.png",
		"icon-144x144.png",
		"icon-152x152.png",
		"icon-192x192.png",
		"icon-384x384.png",
		"icon-512x512.png",
		"icon-72x72.png",
		"icon-96x96.png",
	}
	files := []string{}
	decoders := []string{}

	for _, filename := range fileList {

		b, err := ioutil.ReadFile(CONST_PREFIX + filename)
		if err != nil {
			panic(err)
		}

		constFilename := strings.Replace(
			strings.Replace(strings.ToUpper(filename), ".", "_", -1),
			"-",
			"_",
			-1,
		)

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
				constFilename,
				filename,
			),
		)

		files = append(
			files,
			fmt.Sprintf(
				"CONST_SRC_%s = \"%s\"\n",
				constFilename,
				base64.StdEncoding.EncodeToString(b),
			),
		)

	}

	fileBytes := []byte(
		fmt.Sprintf(
			CONST_FILE_TEMPLATE,
			strings.Join(fileList, `", "`),
			strings.Join(decoders, ""),
			strings.Join(files, ""),
		),
	)

	err := ioutil.WriteFile("gen/const.go", fileBytes, 0777)
	if err != nil {
		panic(err)
	}
}
