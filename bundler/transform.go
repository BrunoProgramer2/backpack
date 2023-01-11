package bundler

import (
	"io/ioutil"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func Transform(name string) string {
	root, _ := os.Getwd()

	bytes, _ := ioutil.ReadFile(root + "/pages/" + name + ".jsx")

	file := string(bytes)

	code := api.Transform(file, api.TransformOptions{
		JSX: api.JSXTransform,

		Loader: api.LoaderJSX,
	})

	return string(code.Code)
}
