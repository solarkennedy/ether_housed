# ether_housed

[![Build Status](https://travis-ci.org/solarkennedy/ether_housed.svg)](https://travis-ci.org/solarkennedy/ether_housed)

[![GoDoc](https://godoc.org/github.com/solarkennedy/ether_housed?status.svg)](https://godoc.org/github.com/solarkennedy/ether_housed)


## About

This is the  API **server** for ether\_house project.
[Client software](https://github.com/solarkennedy/ether_house) runs on an
Arduino and interacts with this API.

In designing the software, I aimed for longevity. I want the software to
continue to run for many years without maintenance. I decided to use golang.

* Go binaries are statically compiled, which means the same binary I compile now
  will continue to run on new platforms for years to come.
* With godeps I can include all compatible libraries together with no external
  dependencies, regardless of their long term state.
* I use Heroku to deploy the code. Heroku is free for small installs and a
  stable platform. They can probably keep this server up better than I can.
* I use a DNS name I can control for service discovery. This gives me the
  flexibility to change platforms over time if necissary.

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

