build:
	go build -o bin/terraform-app-init.exe main.go

clear:
	rm -r test-wd
	