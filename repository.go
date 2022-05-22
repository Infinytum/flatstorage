package flatstorage

import (
	"io/ioutil"
	"reflect"
)

type GenericRepository[T any] struct {
	*FlatStorage
	collection string
}

func (r *GenericRepository[T]) Collection() string {
	return r.collection
}

func (r *GenericRepository[T]) Count() int {
	if !r.CollectionExists(r.Collection()) {
		return 0
	}

	resources, err := ioutil.ReadDir(r.FlatStorage.collectionPath(r.Collection()))
	if err != nil {
		return -1
	}
	return len(resources)
}

func (r *GenericRepository[T]) Delete(resourceId string) error {
	return r.FlatStorage.Delete(r.Collection(), resourceId)
}

func (r *GenericRepository[T]) DeleteAll() error {
	return r.FlatStorage.DeleteAll(r.Collection())
}

func (r *GenericRepository[T]) Exists(resourceId string) bool {
	return r.FlatStorage.Exists(r.Collection(), resourceId)
}

func (r *GenericRepository[T]) Get(resourceId string) (val *T, err error) {
	err = r.Read(r.Collection(), resourceId, &val)
	return
}

func (r *GenericRepository[T]) GetAll() (val []T, err error) {
	values, err2 := r.ReadAll(r.Collection(), reflect.New(r.Type()).Interface())
	err = err2
	if err2 == nil {
		for _, v := range values {
			val = append(val, *v.(*T))
		}
	}
	return
}

func (r *GenericRepository[T]) Type() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

func (r *GenericRepository[T]) Write(resourceId string, val T) error {
	return r.FlatStorage.Write(r.Collection(), resourceId, val)
}

func Repository[T any](fs *FlatStorage) *GenericRepository[T] {
	repo := &GenericRepository[T]{
		FlatStorage: fs,
	}
	repo.collection = reflectTypeKey(repo.Type())
	return repo
}

func NamedRepository[T any](fs *FlatStorage, collection string) *GenericRepository[T] {
	return &GenericRepository[T]{collection: collection, FlatStorage: fs}
}
