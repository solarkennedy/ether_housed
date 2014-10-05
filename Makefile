test:
	go test -v -bench=.

run:
	go build .
	bash -c 'source secrets.sh && \
	./ether_housed'

deps:
# Ignoring for now
#	godep get

push_config:
	./push-config.sh
