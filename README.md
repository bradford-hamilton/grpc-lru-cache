#### LRU cache over grpc
---
#### Running with docker
Build:
```
make docker-build
```
Run:
```
make docker-run
```

## TODO
- Update volume sharing to be more specific so we're not sharing all of $HOME
  - Maybe to ~/.grpc-lru-cache dir, but would that then need to be initialized in some way? Come back to this.
- Host machine needs Docker & Go
- Update save to disk to save the items in the correct order
