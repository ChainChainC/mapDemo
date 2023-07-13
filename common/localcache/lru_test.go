package localcache

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	glru "github.com/hashicorp/golang-lru"
)

func TestMain(m *testing.M) {
	testLRUCache3, _ = glru.NewWithEvict(3, func(key interface{}, value interface{}) {
		if tv, ok := value.(*ttlValue); ok {
			release(tv)
		}
	})
	testLRUCache5, _ = glru.NewWithEvict(5, func(key interface{}, value interface{}) {
		if tv, ok := value.(*ttlValue); ok {
			release(tv)
		}
	})
	m.Run()
	// time.Sleep(100 * time.Second)
	os.Exit(0)
}

func TestLRUCache_Set(t *testing.T) {
	type fields struct {
		cache   *glru.Cache
		timeout time.Duration
	}
	type args struct {
		k string
		v interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "new1",
			fields: fields{
				cache:   testLRUCache3,
				timeout: 3 * time.Second,
			},
			args: args{
				k: "k1",
				v: "v1",
			},
			wantErr: false,
		},
		{
			name: "new2",
			fields: fields{
				cache:   testLRUCache3,
				timeout: 3 * time.Second,
			},
			args: args{
				k: "k2",
				v: "v2",
			},
			wantErr: false,
		},
		{
			name: "new3",
			fields: fields{
				cache:   testLRUCache3,
				timeout: 3, // 3纳秒
			},
			args: args{
				k: "k3",
				v: "v3",
			},
			wantErr: false,
		},
		{
			name: "new4",
			fields: fields{
				cache:   testLRUCache3,
				timeout: 3 * time.Second,
			},
			args: args{
				k: "k4",
				v: "v4",
			},
			wantErr: false,
		},
		{
			name: "new5",
			fields: fields{
				cache:   testLRUCache5,
				timeout: 3 * time.Second,
			},
			args: args{
				k: "555",
				v: []string{"aa", "bb", "cc"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LRUCache{
				cache:   tt.fields.cache,
				timeout: tt.fields.timeout,
			}
			c.Set(tt.args.k, tt.args.v)
		})
	}
}

func TestLRUCache_Len(t *testing.T) {
	type fields struct {
		cache   *glru.Cache
		timeout time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "len",
			fields: fields{
				cache:   testLRUCache3,
				timeout: 3 * time.Second,
			},
			want: 3,
		},
		{
			name: "len",
			fields: fields{
				cache:   testLRUCache5,
				timeout: 3 * time.Second,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LRUCache{
				cache:   tt.fields.cache,
				timeout: tt.fields.timeout,
			}
			if got := c.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUCache_Get(t *testing.T) {
	type fields struct {
		cache   *glru.Cache
		timeout time.Duration
	}
	type args struct {
		k string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}{
		{
			name: "已驱逐",
			fields: fields{
				cache:   testLRUCache3,
				timeout: 3,
			},
			args: args{
				"k1",
			},
			want:  nil,
			want1: false,
		},
		{
			name: "get succ",
			fields: fields{
				cache:   testLRUCache3,
				timeout: 3,
			},
			args: args{
				"k2",
			},
			want:  "v2",
			want1: true,
		},
		{
			name: "get expired",
			fields: fields{
				cache:   testLRUCache3,
				timeout: 3,
			},
			args: args{
				"k3",
			},
			want:  nil,
			want1: false,
		},
		{
			name: "get kkk",
			fields: fields{
				cache:   testLRUCache5,
				timeout: 3,
			},
			args: args{
				"555",
			},
			want:  []string{"aa", "bb", "cc"},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LRUCache{
				cache:   tt.fields.cache,
				timeout: tt.fields.timeout,
			}
			got, got1 := c.Get(tt.args.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLRUCache_Delete(t *testing.T) {
	type fields struct {
		cache   *glru.Cache
		timeout time.Duration
	}
	type args struct {
		k string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delete expired",
			fields: fields{
				cache:   testLRUCache3,
				timeout: 0,
			},
			args:    args{"k3"},
			wantErr: false,
		},
		{
			name: "delete not exist",
			fields: fields{
				cache:   testLRUCache3,
				timeout: 0,
			},
			args:    args{"-k1"},
			wantErr: false,
		},
		{
			name: "delete exist",
			fields: fields{
				cache:   testLRUCache3,
				timeout: 0,
			},
			args:    args{"k2"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LRUCache{
				cache:   tt.fields.cache,
				timeout: tt.fields.timeout,
			}
			if err := c.Delete(tt.args.k); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewLRUCache(t *testing.T) {
	type args struct {
		capacity int64
		d        time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    Cache
		wantErr bool
	}{
		{
			name: "new",
			args: args{
				capacity: 100,
				d:        10 * time.Second,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewLRUCache(tt.args.capacity, tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLRUCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_newValue(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want *ttlValue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newValue(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_release(t *testing.T) {
	type args struct {
		tv *ttlValue
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			release(tt.args.tv)
		})
	}
}

func Test_baseTest(t *testing.T) {
	// var RoomUserCache Cache
	// RoomUserCache, _ = NewLRUCache(20, 10*time.Second)
	// RoomUserCache.Set("a", "a")
	lru, err := NewLRUCache(20, 10*time.Second)
	if err != nil {
		fmt.Print("new lru cache failed.")
		t.Fail()
	}
	lru.Set("RoomId", []string{"userId1", "userId2"})
	res, exp := lru.Get("RoomId")
	if exp {
		fmt.Println(res)
		fmt.Println(reflect.TypeOf(res))
	}
	key := "Roomii"
	go func() {
		lru.Set(key, "ss")
	}()
	time.Sleep(5 * time.Second)
	res, exp = lru.Get(key)
	if exp {
		fmt.Print(res)
	}
	res, exp = lru.Get("RoomId")
	if exp {
		fmt.Print(res)
	}
	// fmt.Print(lru.Get("RoomId"))
}
