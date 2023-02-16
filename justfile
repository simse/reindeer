build:
	go build -o reindeer

pre-dev:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

dev:
	./bin/air -c .air.toml