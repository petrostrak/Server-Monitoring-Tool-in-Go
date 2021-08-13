!/bin/zsh

# This is the bare minimum to run in development. For full list of flags,
# run ./vigilate -help

go build -o Server-Monitoring-Tool-in-Go cmd/web/*.go && ./Server-Monitoring-Tool-in-Go \
-dbuser='petros' \
-pusherHost='localhost' \
-pusherKey='abc123' \
-pusherSecret='123abc' \
-pusherApp="1"
-pusherPort="4001"
-pusherSecure=false