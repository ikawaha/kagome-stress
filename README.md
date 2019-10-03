# kagome stress test

## Server

```shellsession
$ go build ./cmd/stress
$ ./stress
[stressapi] 01:50:15 HTTP "Start" mounted on POST /start
[stressapi] 01:50:15 HTTP "Stop" mounted on POST /stop
[stressapi] 01:50:15 HTTP "Tokenize" mounted on POST /tokenize
[stressapi] 01:50:15 HTTP "./gen/http/openapi.json" mounted on GET /openapi.json
[stressapi] 01:50:15 HTTP server listening on "localhost:8000"
```

## Client

```shellsession
$ go build ./cmd/stress-cli
$ ./stress-cli --help
./stress-cli is a command line client for the stress API.

Usage:
    ./stress-cli [-host HOST][-url URL][-timeout SECONDS][-verbose|-v] SERVICE ENDPOINT [flags]

    -host HOST:  server host (localhost). valid values: localhost
    -url URL:    specify service URL overriding host URL (http://localhost:8080)
    -timeout:    maximum number of seconds to wait for response (30)
    -verbose|-v: print request and response details (false)

Commands:
    stress (start|stop|tokenize)

Additional help:
    ./stress-cli SERVICE [ENDPOINT] --help

Example:
    ./stress-cli stress start
```

### start

```shellsession
$ ./stress-cli stress start
```
### stop

```shellsession
$ ./stress-cli stress stop
```

### tokenize

```shellsession
$ ./stress-cli stress tokenize -body '{"sentence": "すもももももももものうち"}'
```

