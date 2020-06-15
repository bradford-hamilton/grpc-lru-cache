package mem

import (
	"container/list"
	"fmt"
	"reflect"
	"strconv"
	"sync"
	"testing"
)

func TestCache_GetAndSet(t *testing.T) {
	type fields struct {
		cap   int
		ll    *list.List
		items map[interface{}]*list.Element
		mu    sync.Mutex
	}
	type args struct {
		Key interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		ok     bool
	}{
		{
			name: "",
			fields: fields{
				cap:   10,
				ll:    list.New(),
				items: make(map[interface{}]*list.Element),
			},
			args: args{Key: "someKey"},
			want: "someValue",
			ok:   true,
		},
		{
			name: "",
			fields: fields{
				cap:   10,
				ll:    list.New(),
				items: make(map[interface{}]*list.Element),
			},
			args: args{Key: 255},
			want: 10598,
			ok:   true,
		},
		{
			name: "",
			fields: fields{
				cap:   10,
				ll:    list.New(),
				items: make(map[interface{}]*list.Element),
			},
			args: args{Key: struct{ name string }{"daaaavid"}},
			want: struct{ name string }{"daaaavid"},
			ok:   true,
		},
		{
			name: "",
			fields: fields{
				cap:   10,
				ll:    list.New(),
				items: make(map[interface{}]*list.Element),
			},
			args: args{Key: 0.45},
			want: 0.45,
			ok:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LRUCache{
				cache: &cache{
					cap:   tt.fields.cap,
					ll:    tt.fields.ll,
					items: tt.fields.items,
				},
				mu: tt.fields.mu,
			}
			_, ok := c.Set(tt.args.Key, tt.want)
			got, ok := c.Get(tt.args.Key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cache.Get() got = %v, want %v", got, tt.want)
			}
			if ok != tt.ok {
				t.Errorf("Cache.Get() ok = %v, want %v", ok, tt.ok)
			}
		})
	}
}

func TestLRUCache_Grow(t *testing.T) {
	type fields struct {
		cache *cache
		mu    sync.Mutex
	}
	type args struct {
		additionalCap int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantSize int
		wantErr  bool
	}{
		{
			name: "when called with an additional capacity of < 1, it should error",
			fields: fields{
				cache: &cache{
					cap:   1,
					ll:    &list.List{},
					items: map[interface{}]*list.Element{nil: {Value: nil}},
				},
				mu: sync.Mutex{},
			},
			args:     args{additionalCap: 0},
			wantSize: 1,
			wantErr:  true,
		},
		{
			name: "when called with an additional capacity 1, the cache should grow by 1",
			fields: fields{
				cache: &cache{
					cap:   1,
					ll:    &list.List{},
					items: map[interface{}]*list.Element{nil: {Value: nil}},
				},
				mu: sync.Mutex{},
			},
			args:     args{additionalCap: 1},
			wantSize: 2,
			wantErr:  false,
		},
		{
			name: "when called with an additional capacity 100, the cache should grow by 100",
			fields: fields{
				cache: &cache{
					cap:   100,
					ll:    &list.List{},
					items: map[interface{}]*list.Element{nil: {Value: nil}},
				},
				mu: sync.Mutex{},
			},
			args:     args{additionalCap: 100},
			wantSize: 200,
			wantErr:  false,
		},
		{
			name: "when called with an additional capacity that grows beyond the maxCacheSize, it should error",
			fields: fields{
				cache: &cache{
					cap:   maxCacheSize,
					ll:    &list.List{},
					items: map[interface{}]*list.Element{nil: {Value: nil}},
				},
				mu: sync.Mutex{},
			},
			args:     args{additionalCap: 1},
			wantSize: 1000000,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lru := &LRUCache{cache: tt.fields.cache, mu: tt.fields.mu}
			if err := lru.Grow(tt.args.additionalCap); (err != nil) != tt.wantErr {
				t.Errorf("LRUCache.Grow() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantSize != lru.cache.cap {
				t.Errorf("Expected a cache capacity of %d after growing, got: %d", tt.wantSize, lru.cache.cap)
			}
		})
	}
}

var sink bool
var item interface{}

func BenchmarkSetItem(b *testing.B) {
	c, err := NewLRUCache(1000)
	if err != nil {
		fmt.Printf("failed to create client: %v\n", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, ok := c.Set(i, "value#"+strconv.Itoa(i))
		sink = ok
	}
}

func BenchmarkGetItem(b *testing.B) {
	c, err := NewLRUCache(1000)
	if err != nil {
		fmt.Printf("failed to create client: %v\n", err)
	}
	for i := 0; i < 1000; i++ {
		_, ok := c.Set(i, "value#"+strconv.Itoa(i))
		sink = ok
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		itm, ok := c.Get("value#" + strconv.Itoa(i))
		sink = ok
		item = itm
	}
}
