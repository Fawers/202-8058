package store

import (
	"sync"
)

const (
	LABEL_ENABLED string = "4DISCOVERY_ENABLED"
	LABEL_NAME    string = "4DISCOVERY_NAME"
)

type Service struct {
	Name        string
	ContainerId string
	Image       string
	IPAddress   string
	Endpoint    string
	Hostname    string
	Labels      map[string]string
	Ports       []string
}

// Labels que queremos nos nossos containers são
// 4DISCOVERY_ENABLED = true
// 4DISCOVERY_NAME = ""

type ServiceStore struct {
	services map[string]Service
	sync.RWMutex
}

func New() ServiceStore {
	return ServiceStore{
		services: make(map[string]Service)}
}

func (ss *ServiceStore) Add(name string, service Service) {
	ss.Lock()
	defer ss.Unlock()

	ss.services[name] = service
}

func (ss *ServiceStore) Remove(name string) {
	ss.Lock()
	defer ss.Unlock()

	delete(ss.services, name)
}

func (ss *ServiceStore) List() []Service {
	ss.RLock()
	defer ss.RUnlock()

	services := []Service{}

	for _, svc := range ss.services {
		services = append(services, svc)
	}

	return services
}
