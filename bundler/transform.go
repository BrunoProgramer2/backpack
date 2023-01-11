package bundler

import (
	"io/ioutil"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func Transform(name string) string {
	root, _ := os.Getwd()

	api.Build(api.BuildOptions{
		EntryPoints: []string{root + "/pages/" + name + ".jsx"},
		Outfile:     root + "/tmp/" + name + ".js",
		Bundle:      true,
		Write:       true,
		Loader: map[string]api.Loader{
			".jsx": api.LoaderJSX,
		},
	})

	bytes, _ := ioutil.ReadFile(root + "/tmp/" + name + ".js")

	file := string(bytes)

	return file
}
