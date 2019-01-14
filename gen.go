package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/base64"
)

const (
	CONST_PREFIX = "angular/dist/angular/"
	CONST_FILE_TEMPLATE = `package static

const(
	%s
)
`
)

func main() {

	list := []string{
		"favicon.ico",
		"index.html",
		"main.js",
		"polyfills.js",
		"runtime.js",
		"styles.css",
	}

	files := []string{}

	for _, item := range list {

		b, err := ioutil.ReadFile(CONST_PREFIX + item)
		if err != nil {
			panic(err)
		}

		files = append(
			files,
			fmt.Sprintf(
				`	CONST_SRC_%s = "%s"`,
				strings.Replace(strings.ToUpper(item), ".", "_", -1),
				base64.StdEncoding.EncodeToString(b),
			),
		)

	}

	fileBytes := []byte(
		fmt.Sprintf(
			CONST_FILE_TEMPLATE,
			strings.Join(files, "\n"),
		),
	)

	err := ioutil.WriteFile("gen/const.go", fileBytes, 0777)
	if err != nil {
		panic(err)
	}
}
