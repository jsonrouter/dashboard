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
	"encoding/base64"
	"github.com/jsonrouter/core/http"
	"github.com/jsonrouter/core/tree"
)

func FileList() []string {
	return []string{"%s"}
}

func (self *Static) Dashboard(node *tree.Node) {

	node = node.Add("dashboard")

	for _, filename := range FileList() {

		%s

		node.Add(filename, "$item").GET(
			func (req http.Request) *http.Status {

				return req.Respond(
					self.files[filename].Cache,
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
			strings.Join(files, ""),
		),
	)

	err := ioutil.WriteFile("gen/const.go", fileBytes, 0777)
	if err != nil {
		panic(err)
	}
}
