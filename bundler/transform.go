package bundler

import (
	"io/ioutil"

	"github.com/evanw/esbuild/pkg/api"
)

func Transform(dir string, name string) string {
	//trasform
	api.Build(api.BuildOptions{
		EntryPoints: []string{dir + "/pages/" + name + ".jsx"},
		Outfile:     dir + "/tmp/" + name + ".js",
		Bundle:      true,
		Write:       true,
		Loader: map[string]api.Loader{
			".jsx": api.LoaderJSX,
		},
	})

	bytes, _ := ioutil.ReadFile(dir + "/tmp/" + name + ".js")

	file := string(bytes)

	return file
}
