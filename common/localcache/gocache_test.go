package localcache

import (
	"reflect"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
)

func TestNewGoCache(t *testing.T) {
	type args struct {
		defaultExpiration time.Duration
		cleanupInterval   time.Duration
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
				defaultExpiration: 10,
				cleanupInterval:   10,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewGoCache(tt.args.defaultExpiration, tt.args.cleanupInterval)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGoCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_goCache_Set(t *testing.T) {
	type fields struct {
		cache *cache.Cache
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "set 1",
			fields: fields{
				cache: testGoCache,
			},
			args: args{
				key:   "k1",
				value: "v1",
			},
			wantErr: false,
		},
		{
			name: "reset 1",
			fields: fields{
				cache: testGoCache,
			},
			args: args{
				key:   "k1",
				value: "v1.1",
			},
			wantErr: false,
		},
		{
			name: "set 2",
			fields: fields{
				cache: testGoCache,
			},
			args: args{
				key:   "k2",
				value: "v2",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &goCache{
				cache: tt.fields.cache,
			}
			c.Set(tt.args.key, tt.args.value)
		})
	}
}

func Test_goCache_SetWithExpirationGet(t *testing.T) {
	type fields struct {
		cache *cache.Cache
	}
	type args struct {
		key         string
		value       interface{}
		expDuration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "set 1",
			fields: fields{
				cache: testGoCache,
			},
			args: args{
				key:         "k2",
				value:       "v1",
				expDuration: time.Second * time.Duration(2), // 过期时间设置为2s
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &goCache{
				cache: tt.fields.cache,
			}
			c.SetWithExpiration(tt.args.key, tt.args.value, tt.args.expDuration)
			got, ok := c.Get(tt.args.key)
			if !ok {
				t.Errorf("Get() error")
			}
			if !reflect.DeepEqual(got, tt.args.value) {
				t.Errorf("Get() got = %#v, wantErr %#v", got, tt.args.value)
			}
			// key还没有过期
			time.Sleep(time.Second * time.Duration(1))
			got, ok = c.Get(tt.args.key)
			if !ok {
				t.Errorf("Get() error")
			}
			if !reflect.DeepEqual(got, tt.args.value) {
				t.Errorf("Get() got = %#v, wantErr %#v", got, tt.args.value)
			}

			c.SetWithExpiration(tt.args.key, tt.args.value, tt.args.expDuration)
			got, ok = c.Get(tt.args.key)
			if !ok {
				t.Errorf("Get() error")
			}
			if !reflect.DeepEqual(got, tt.args.value) {
				t.Errorf("Get() got = %#v, wantErr %#v", got, tt.args.value)
			}
			// key过期
			time.Sleep(time.Second * time.Duration(3))
			got, ok = c.Get(tt.args.key)
			if ok {
				t.Errorf("Get() error")
			}
			if reflect.DeepEqual(got, tt.args.value) {
				t.Errorf("Get() got = %#v, wantErr %#v", got, tt.args.value)
			}
		})
	}
}

func Test_goCache_Get(t *testing.T) {
	type fields struct {
		cache *cache.Cache
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}{
		{
			name: "get 1",
			fields: fields{
				cache: testGoCache,
			},
			args: args{
				key: "k1",
			},
			want:  "v1.1",
			want1: true,
		},
		{
			name: "get not exist",
			fields: fields{
				cache: testGoCache,
			},
			args: args{
				key: "-k1",
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &goCache{
				cache: tt.fields.cache,
			}
			got, got1 := c.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_goCache_Delete(t *testing.T) {
	type fields struct {
		cache *cache.Cache
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delete not exist",
			fields: fields{
				cache: testGoCache,
			},
			args: args{
				key: "-k1",
			},
			wantErr: false,
		},
		{
			name: "delete  exist",
			fields: fields{
				cache: testGoCache,
			},
			args: args{
				key: "k1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &goCache{
				cache: tt.fields.cache,
			}
			if err := c.Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
