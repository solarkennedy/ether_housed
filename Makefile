test:
	TZ="UTC" go test -v -bench=.
	# Not on tip in travis
	#go vet

run:
	go build . 
	bash -c 'source secrets.sh && \
	./ether_housed'

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
