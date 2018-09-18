package flatstorage

import (
	"os"
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFlatStorage_Delete(t *testing.T) {
	os.MkdirAll("/tmp/flatstorage/test", 0777)
	os.Create("/tmp/flatstorage/test/deleteme.json")
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
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test deletion of inexistent file",
			fields: fields{
				logger:  logrus.New(),
				mutex:   &sync.Mutex{},
				mutexes: make(map[string]*sync.Mutex, 0),
				path:    "/tmp/flatstorage",
			},
			args: args{
				collection: "test",
				resource:   "nonexistent",
			},
			wantErr: false,
		},
		{
			name: "Delete existing file",
			fields: fields{
				logger:  logrus.New(),
				mutex:   &sync.Mutex{},
				mutexes: make(map[string]*sync.Mutex, 0),
				path:    "/tmp/flatstorage",
			},
			args: args{
				collection: "test",
				resource:   "deleteme",
			},
			wantErr: false,
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
			if err := fs.Delete(tt.args.collection, tt.args.resource); (err != nil) != tt.wantErr {
				t.Errorf("FlatStorage.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFlatStorage_DeleteAll(t *testing.T) {
	os.MkdirAll("/tmp/flatstorage/test2", 0777)
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
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test deletion of inexistent collection",
			fields: fields{
				logger:  logrus.New(),
				mutex:   &sync.Mutex{},
				mutexes: make(map[string]*sync.Mutex, 0),
				path:    "/tmp/flatstorage",
			},
			args: args{
				collection: "test3",
			},
			wantErr: false,
		},
		{
			name: "Delete existing collection",
			fields: fields{
				logger:  logrus.New(),
				mutex:   &sync.Mutex{},
				mutexes: make(map[string]*sync.Mutex, 0),
				path:    "/tmp/flatstorage",
			},
			args: args{
				collection: "test2",
			},
			wantErr: false,
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
			if err := fs.DeleteAll(tt.args.collection); (err != nil) != tt.wantErr {
				t.Errorf("FlatStorage.DeleteAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFlatStorage_Exists(t *testing.T) {
	os.MkdirAll("/tmp/flatstorage/exist/", 0777)
	os.Create("/tmp/flatstorage/exist/exist.json")
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
		want   bool
	}{
		{
			name: "Test existing resource",
			fields: fields{
				logger:  logrus.New(),
				mutex:   &sync.Mutex{},
				mutexes: make(map[string]*sync.Mutex, 0),
				path:    "/tmp/flatstorage",
			},
			args: args{
				collection: "exist",
				resource:   "exist",
			},
			want: true,
		},
		{
			name: "Test non-existing resource",
			fields: fields{
				logger:  logrus.New(),
				mutex:   &sync.Mutex{},
				mutexes: make(map[string]*sync.Mutex, 0),
				path:    "/tmp/flatstorage",
			},
			args: args{
				collection: "exist",
				resource:   "nonexist",
			},
			want: false,
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
			if got := fs.Exists(tt.args.collection, tt.args.resource); got != tt.want {
				t.Errorf("FlatStorage.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatStorage_CollectionExists(t *testing.T) {
	os.MkdirAll("/tmp/flatstorage/exist_collection/", 0777)
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
		want   bool
	}{
		{
			name: "Test existing collection",
			fields: fields{
				logger:  logrus.New(),
				mutex:   &sync.Mutex{},
				mutexes: make(map[string]*sync.Mutex, 0),
				path:    "/tmp/flatstorage",
			},
			args: args{
				collection: "exist_collection",
			},
			want: true,
		},
		{
			name: "Test non-existing collection",
			fields: fields{
				logger:  logrus.New(),
				mutex:   &sync.Mutex{},
				mutexes: make(map[string]*sync.Mutex, 0),
				path:    "/tmp/flatstorage",
			},
			args: args{
				collection: "non-exist",
			},
			want: false,
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
			if got := fs.CollectionExists(tt.args.collection); got != tt.want {
				t.Errorf("FlatStorage.CollectionExists() = %v, want %v", got, tt.want)
			}
		})
	}
	os.RemoveAll("/tmp/flatstorage/exist_collection/")
}

func ExampleFlatStorage_read() {
	fs, err := NewFlatStorage("/var/db")
	if err != nil {
		panic(err)
	}

	type Test struct {
		Name string
	}

	// File: /var/db/test/test.json
	//
	//	{
	//		"name": "Hello World"
	//  }
	//

	test := Test{}
	fs.Read("test", "test", &test)

	if err != nil {
		panic(err)
	}

	print(test.Name)
	// Output: Hello World
}

func ExampleFlatStorage_write() {
	fs, err := NewFlatStorage("/var/db")
	if err != nil {
		panic(err)
	}

	type Test struct {
		Name string
	}

	test := Test{
		Name: "Hello World",
	}

	fs.Write("test", "test", &test)

	if err != nil {
		panic(err)
	}
}

func ExampleFlatStorage_delete() {
	fs, err := NewFlatStorage("/var/db")
	if err != nil {
		panic(err)
	}

	fs.Delete("test", "test")

	if err != nil {
		panic(err)
	}
}

func ExampleFlatStorage_deleteAll() {
	fs, err := NewFlatStorage("/var/db")
	if err != nil {
		panic(err)
	}

	fs.DeleteAll("test")

	if err != nil {
		panic(err)
	}
}

func ExampleFlatStorage_readAll() {
	fs, err := NewFlatStorage("/var/db")
	if err != nil {
		panic(err)
	}

	type Test struct {
		Name string
	}

	// File: /var/db/test/test.json
	//
	//	{
	//		"name": "Hello World"
	//  }
	//

	test := Test{}
	resources, err := fs.ReadAll("test", &test)

	if err != nil {
		panic(err)
	}

	print(resources[0])
	// Output: { Name: "Hello World" }
}
