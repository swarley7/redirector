# redirector

A simple HTTP redirector. Redirects requests (using HTTP 301 status code) to an arbitrary URL. Useful for testing SSRF, etc. where allow-listing may be in place for specific endpoints.

# Install / build

`go install github.com/swarley7/redirector`

# Usage:
Usage is fairly simple; there are three args that can be supplied on the commandline.

* `host`: the listening host ip / hostname (e.g. localhost, or 0.0.0.0)
* `port`: the listening port number (e.g. 80)
* `redirUrl`: the target URL to redirect requests to (e.g. http://www.example.com)


## Sample: 

```
redirector -host localhost -port 8080 -redirUrl http://example.com/
```

