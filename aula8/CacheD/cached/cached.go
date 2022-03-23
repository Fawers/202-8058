package cached

import (
	"errors"
	"sync"
)

var ErrKeyAlreadyExists = errors.New("chave já existe no banco")
var ErrKeyNotFound = errors.New("chave não existe no banco")

type CacheD struct {
	d map[string]string
	sync.RWMutex
}

func NewCacheD() *CacheD {
	return &CacheD{
		d: make(map[string]string),
	}
}

func (cd *CacheD) Add(key, value string) error {
	cd.Lock()
	defer cd.Unlock()

	if _, found := cd.d[key]; found {
		return ErrKeyAlreadyExists
	}

	cd.d[key] = value
	return nil
}

func (cd *CacheD) Del(key string) error {
	cd.Lock()
	defer cd.Unlock()

	if _, found := cd.d[key]; !found {
		return ErrKeyNotFound
	}

	delete(cd.d, key)
	return nil
}

func (cd *CacheD) Get(key string) (string, error) {
	cd.RLock()
	defer cd.RUnlock()

	if value, found := cd.d[key]; found {
		return value, nil
	} else {
		return "", ErrKeyNotFound
	}
}

func (cd *CacheD) Update(key, value string) bool {
	cd.Lock()
	defer cd.Unlock()

	_, updated := cd.d[key]
	cd.d[key] = value
	return updated
}

func (cd *CacheD) GetAll() [][2]string {
	cd.RLock()
	defer cd.RUnlock()

	r := make([][2]string, 0, len(cd.d))

	for key, val := range cd.d {
		r = append(r, [...]string{key, val})
	}

	return r
}

func (cd *CacheD) DelAll() {
	cd.Lock()
	defer cd.Unlock()

	for key := range cd.d {
		delete(cd.d, key)
	}
}

// ADD key:string value:string -> error
// DEL key:string -> error
// GET key:string -> value:string | error
// UPDATE key:string value:string -> bool
// GETALL -> [](string, string)
// DELALL
