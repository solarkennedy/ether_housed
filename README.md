# ether_housed

[![Build Status](https://travis-ci.org/solarkennedy/ether_housed.svg)](https://travis-ci.org/solarkennedy/ether_housed)

[![GoDoc](https://godoc.org/github.com/solarkennedy/ether_housed?status.svg)](https://godoc.org/github.com/solarkennedy/ether_housed)

The API server for ether\_house.

## Usage

This is an API server. It isn't really designed to be useful to humans.
All it really does is store shared state and configuration for ether\_house 
clients as they report in.

## Requirements

* golang 1.3
* memcached (optional)
* heroku (optional)

## Install

Using heroku

* Install heroku, get an account, log in with the tool
* Clone this repo
* Make a copy of the secrets file and make your own secrets
* Create an app

    # Use the golang buildback
    heroku create -b https://github.com/kr/heroku-buildpack-go.git

* Enable Memcache (optional)

    heroku addons:add memcachedcloud:25

* Push

    git push origin heroku

* Push your secrets

    make push_config

## Security

The client uses HTTP in combination with an API key to authenticate and
authorize. A compromised API key means someone can change the stored state
associated with that key, as well as retrieve the mac address associated
with that API key.

This key is considered a simple shared secret. If you want new ones, change the
secrets.sh file and program the new key on the arduino.

