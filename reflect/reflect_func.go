package reflect

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type MapFunc struct {
	db map[string]reflect.Value
	mu sync.Mutex
}

func New() *MapFunc {
	return &MapFunc{
		db: make(map[string]reflect.Value),
	}
}

func (m *MapFunc) Add(key string,fn interface{}) (err error) {
	defer func() {
		if erc := recover();erc != nil {
			err = fmt.Errorf("%v",erc)
			return
		}
	}()
	m.mu.Lock()
	defer m.mu.Unlock()
	of := reflect.ValueOf(fn)
	m.db[key] = of
	return
}

func (m *MapFunc) Call(key string,params ...interface{}) (result []reflect.Value,err error) {
	defer func() {
		if erc := recover();erc != nil {
			err = fmt.Errorf("%v",erc)
			return
		}
	}()
	m.mu.Lock()
	defer m.mu.Unlock()
	value,ex := m.db[key]
	if !ex {
		return nil,errors.New("NOT EX")
	}
	if len(params) != value.Type().NumIn() {
		err = fmt.Errorf("%d parameters are required, but %d parameters are entered",value.Type().NumIn(),len(params))
		return
	}
	req := make([]reflect.Value,len(params))
	for k,v := range params {
		req[k] = reflect.ValueOf(v)
	}

	return value.Call(req),nil
}