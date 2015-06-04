# Taskcluster Proxy

This is the proxy server which is used in the docker-worker which allows
individual tasks to talk to various taskcluster services (auth, queue,
scheduler) without hardcoding credentials into the containers
themselves.

Credentials are expected to be passed via the `TASKCLUSTER_CLIENT_ID`
and `TASKCLUSTER_ACCESS_TOKEN` environment variables.


## Examples

TBD

## Deployment

The proxy server can be deployed directly by building `proxy/main.go`
but the prefered method is via the `./build.sh` script which will
compile the proxy server for linux/amd64 and deploy the server to a
docker image. [Godep](https://github.com/tools/godep) is required to run
this script.

```sh
./build.sh user/relengapi-proxy-server
```

## Download via `go get`

Set up your [GOPATH](https://golang.org/doc/code.html)

```sh
go get github.com/djmitche/relengapi-proxy
```

## Hacking

To build, just run

```sh
godep go build
```

## Tests

To run the full test suites you need a [RelengAPI](https://api.pub.build.mozilla.org/) token.
That token must have at least `base.tokens.tmp.issue`, as well as any permissions tasks may need.
The token is supplied with the --relengapi-token command-line argument.
Note that credentials must not be included in environment variables!
