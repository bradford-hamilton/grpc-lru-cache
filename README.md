### LRU cache over grpc

I built the cache initially to support any type for both keys and values.. This is really nice when using it only within the context of Go code, however when I decided to build a grpc API for it, I realized that would get messy quickly. 

This has brought on some unnecessary conversions from `interface{} -> string` and `[]interface{} -> []string` which I could avoid by changing the cache to work with strings only... However I don't want to do that because it would feel like I'm removing very useful functionality. Although maybe I will build it in as an optional mode... 

Current binary size: 4.5MB.

Main goals:
- make it really fast
- make a super small application binary

### Run with docker
Build:
```
make docker-build
```
Run:
```
make docker-run
```

## Deploy to GCP
[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)