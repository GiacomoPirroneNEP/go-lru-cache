# go-lru-cache

Simple Least Recently Used Cache data structure implementation in Go(lang).

### Initializing LRU Cache
```go
// Makes a new LRU Cache of size 100
newLruCache := createLRU(100)
```

### Getting Value From LRU Cache
```go
// Gets element from LRU Cache with key of "key"
elem, err := newLRUCache.GetElem("key)
if err != nil {
  log.Println(err.Error)
}
```
