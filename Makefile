test:
	go test -v

run:
	bash -c 'source secrets.sh && \
	go run main.go'

deps:
# Ignoring for now
#	godep get
