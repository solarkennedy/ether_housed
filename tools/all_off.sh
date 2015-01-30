#!/bin/bash
set -eu
. `dirname $0`/../secrets.sh
if [ -z "${ETHERHOUSE_HOST}" ]; then
  ETHERHOUSE_HOST=etherhouse.xkyle.com
fi
curl "${ETHERHOUSE_HOST}/off?id=0&api_key=${APIKEY0}"
curl "${ETHERHOUSE_HOST}/off?id=1&api_key=${APIKEY1}"
curl "${ETHERHOUSE_HOST}/off?id=2&api_key=${APIKEY2}"
curl "${ETHERHOUSE_HOST}/off?id=3&api_key=${APIKEY3}"
curl "${ETHERHOUSE_HOST}/off?id=4&api_key=${APIKEY4}"
curl "${ETHERHOUSE_HOST}/off?id=5&api_key=${APIKEY5}"
curl "${ETHERHOUSE_HOST}/off?id=6&api_key=${APIKEY6}"
curl "${ETHERHOUSE_HOST}/off?id=7&api_key=${APIKEY7}"

