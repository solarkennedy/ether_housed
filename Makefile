test:
	TZ="UTC" go test -v -bench=.

run:
	go build . 
	bash -c 'source secrets.sh && \
	./ether_housed'

deps:
	godep get
	godep save

push_config:
	./push-config.sh

clean:
	rm ether_housed

logs:
	heroku logs  -t -n 0

fmt:
	go fmt .

deploy:
	git push heroku master
