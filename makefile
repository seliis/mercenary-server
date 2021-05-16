build:
	npm run build
	mkdir -p /mnt/d/mss/web/
	mv build/* /mnt/d/mss/web/
	GOOS=windows GOARCH=amd64 go build -o /mnt/d/mss/server.exe
	rm -rf build/
	cp -r userdata.csv /mnt/d/mss/
	