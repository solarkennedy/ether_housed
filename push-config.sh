#!/bin/bash
set -eu
IFS="
"
for CONFIG in `grep '=' secrets.sh | sed 's/export //g'`; do
  heroku config:set $CONFIG
done
