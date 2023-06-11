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

## TODO
- Share volume between docker <-> machine (maybe update cache location)
- Update save to disk to save the items in the correct order
