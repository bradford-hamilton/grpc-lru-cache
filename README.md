### LRU cache over grpc

Current binary size: 4.5MB.

Main goals:
- make it fast
- make it small

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

## TODO
- Update save to disk to save the items in the correct order
