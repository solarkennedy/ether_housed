#!/bin/bash
set -eu
. ../secrets.sh
curl "etherhouse.xkyle.com/on?id=0&api_key=${APIKEY0}"
curl "etherhouse.xkyle.com/on?id=1&api_key=${APIKEY1}"
curl "etherhouse.xkyle.com/on?id=2&api_key=${APIKEY2}"
curl "etherhouse.xkyle.com/on?id=3&api_key=${APIKEY3}"
curl "etherhouse.xkyle.com/on?id=4&api_key=${APIKEY4}"
curl "etherhouse.xkyle.com/on?id=5&api_key=${APIKEY5}"
curl "etherhouse.xkyle.com/on?id=6&api_key=${APIKEY6}"
curl "etherhouse.xkyle.com/on?id=7&api_key=${APIKEY7}"

