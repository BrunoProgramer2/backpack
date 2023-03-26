package bundler

func View(dir string, name string) []byte {
	// trasnform
	bundle := Transform(dir, name)

	return []byte("<div id='root'></div>" + "<script>" + bundle + "</script>")
}
