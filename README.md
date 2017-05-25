# COAP Experiment

This repo contains a minimal HTTP server and a minimal COAP server, where the
COAP server attempts to provide equivalent functionality to the HTTP server.

So far I've only attempted to start these servers using:

```bash
$ go run http.go
```

or

```bash
$ go run coap.go
```

## Dependencies

I've used glide to install the dustin/go-coap library, but the dependency is
vendored so should already be present.

## Clients

I made a couple of minimal clients for the COAP server in node.js and ruby. See
the README in each of the node/ruby folders.
