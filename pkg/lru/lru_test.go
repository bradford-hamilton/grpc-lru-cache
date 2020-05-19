package lru

import (
	"container/list"
	"reflect"
	"sync"
	"testing"
)

func TestCache_GetAndSet(t *testing.T) {
	type fields struct {
		size  int
		list  *list.List
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
				size:  10,
				list:  list.New(),
				items: make(map[interface{}]*list.Element),
			},
			args: args{Key: "someKey"},
			want: "someValue",
			ok:   true,
		},
		{
			name: "",
			fields: fields{
				size:  10,
				list:  list.New(),
				items: make(map[interface{}]*list.Element),
			},
			args: args{Key: 255},
			want: 10598,
			ok:   true,
		},
		{
			name: "",
			fields: fields{
				size:  10,
				list:  list.New(),
				items: make(map[interface{}]*list.Element),
			},
			args: args{Key: struct{ name string }{"daaaavid"}},
			want: struct{ name string }{"daaaavid"},
			ok:   true,
		},
		{
			name: "",
			fields: fields{
				size:  10,
				list:  list.New(),
				items: make(map[interface{}]*list.Element),
			},
			args: args{Key: 0.45},
			want: 0.45,
			ok:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				size:  tt.fields.size,
				list:  tt.fields.list,
				items: tt.fields.items,
				mu:    tt.fields.mu,
			}
			if ok := c.Set(tt.args.Key, tt.want); !ok {
				t.Errorf("failed to Set() cache with key: %s and value: %s", tt.args.Key, tt.want)
			}
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
