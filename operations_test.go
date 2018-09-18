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
		{
			name: "Verify deletion",
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
			wantErr: pathExists("/tmp/flatstorage/test/deleteme"),
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
