EVENTS=`papertrail --min-time "24 hours ago" "program:app/web.1 AND (turn_off OR turn_on)"`

for HOUSE in {0..7} ; do
	echo "Usage report for etherhouse$HOUSE:"
	echo "$EVENTS" | egrep ": ${HOUSE} $"
	echo
done
