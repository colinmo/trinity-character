win:
	cd src;go build .

manual:
	cd src ; pandoc -o Manual.pdf manual.markdown