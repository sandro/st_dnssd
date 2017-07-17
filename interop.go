package st_dnssd

import "sync"

type CallbackState struct {
	store []interface{}
	mutex *sync.Mutex
}

func (o *CallbackState) Add(v interface{}) int {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.store = append(o.store, v)
	return len(o.store) - 1
}

func (o *CallbackState) Get(i int) interface{} {
	return o.store[i]
}

var callbackState CallbackState = CallbackState{}

func init() {
	callbackState.store = make([]interface{}, 0)
	callbackState.mutex = &sync.Mutex{}
}
