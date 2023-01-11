package bundler

func View(name string) []byte {
	bundle := Transform(name)

	return []byte("<div id='root'></div>" + "<script>" + bundle + "</script>")
}
