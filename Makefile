test:
	go test -v -bench=.

run:
	bash -c 'source secrets.sh && \
	go run main.go'

deps:
# Ignoring for now
#	godep get
