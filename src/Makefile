coverage_html:
	go test -v -coverprofile=cover.out ./...
	go tool cover -html cover.out -o cover.html
	rm cover.out