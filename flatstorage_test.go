package flatstorage

import (
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
)

func ExampleNewFlatStorage() {
	fs, err := NewFlatStorage("/var/db")
	if err != nil {
		panic(err)
	}

	print(fs.CollectionExists("test"))

	// Output: false
}

func ExampleFlatStorage_read() {
	fs, err := NewFlatStorage("/var/db")
	if err != nil {
		panic(err)
	}

	type Test struct {
		Name string
	}

	test := Test{}
	fs.Read("test", "test", &test)

	if err != nil {
		panic(err)
	}

	print(test.Name)
	// Output: Hello World
}

func TestFlatStorage_resourcePath(t *testing.T) {
	type fields struct {
		mutex   *sync.Mutex
		mutexes map[string]*sync.Mutex
		path    string
		logger  *logrus.Logger
	}
	type args struct {
		collection string
		resource   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Test resource path generating #1",
			fields: fields{
				path: "/tmp",
			},
			args: args{
				collection: "test",
				resource:   "test1",
			},
			want: "/tmp/test/test1.json",
		},
		{
			name: "Test resource path generating #2",
			fields: fields{
				path: "/tmp",
			},
			args: args{
				collection: "test",
				resource:   "test1.json",
			},
			want: "/tmp/test/test1.json",
		},
		{
			name: "Test resource path generating #2",
			fields: fields{
				path: "/tmp",
			},
			args: args{
				collection: "test",
				resource:   "test1.xml",
			},
			want: "/tmp/test/test1.xml.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlatStorage{
				mutex:   tt.fields.mutex,
				mutexes: tt.fields.mutexes,
				path:    tt.fields.path,
				logger:  tt.fields.logger,
			}
			if got := fs.resourcePath(tt.args.collection, tt.args.resource); got != tt.want {
				t.Errorf("FlatStorage.resourcePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatStorage_collectionPath(t *testing.T) {
	type fields struct {
		mutex   *sync.Mutex
		mutexes map[string]*sync.Mutex
		path    string
		logger  *logrus.Logger
	}
	type args struct {
		collection string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Test collection path generating #1",
			fields: fields{
				path: "/tmp",
			},
			args: args{
				collection: "test",
			},
			want: "/tmp/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FlatStorage{
				mutex:   tt.fields.mutex,
				mutexes: tt.fields.mutexes,
				path:    tt.fields.path,
				logger:  tt.fields.logger,
			}
			if got := fs.collectionPath(tt.args.collection); got != tt.want {
				t.Errorf("FlatStorage.collectionPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
