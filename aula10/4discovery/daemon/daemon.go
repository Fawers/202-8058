package daemon

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	docker "github.com/docker/docker/client"

	"4discovery/store"
)

type Daemon struct {
	HTTPEndpoint string

	log *log.Logger

	dockerClient *docker.Client
	dockerCtx    context.Context

	store store.ServiceStore
}

func New(address string) (*Daemon, error) {
	var d Daemon

	err := d.initDocker()

	if err != nil {
		return nil, err
	}

	if !d.checkDockerConn() {
		return nil, fmt.Errorf("conexão com docker está MORTA")
	}

	d.log = log.New(os.Stdout, "daemon", log.Flags())
	d.store = store.New()
	d.HTTPEndpoint = address

	return &d, nil
}

func (d *Daemon) Start() error {
	comms := make(chan error)

	go d.processDockerEvents(comms)
	go d.processHttpReq(comms)

	d.log.Printf("Processos do 4discovery inicializados\n")

	return <-comms
}

func (d *Daemon) AddService(cid string) {
	c, err := d.dockerClient.ContainerInspect(d.dockerCtx, cid)
	if err != nil {
		return
	}

	enabled, found := c.Config.Labels[store.LABEL_ENABLED]
	if !found || enabled != "true" {
		return
	}

	name, found := c.Config.Labels[store.LABEL_NAME]
	if !found {
		return
	}

	svc := store.Service{
		Name:        name,
		ContainerId: cid,
		Image:       c.Config.Image,
		IPAddress:   c.NetworkSettings.IPAddress,
		Endpoint:    c.NetworkSettings.EndpointID,
		Hostname:    c.Config.Hostname,
		Labels:      c.Config.Labels,
	}

	for port := range c.NetworkSettings.Ports {
		svc.Ports = append(svc.Ports, string(port))
	}

	d.log.Printf("adicionando serviço %v\n", svc)
	d.store.Add(name, svc)
}

func (d *Daemon) RemoveService(cid string) {
	c, err := d.dockerClient.ContainerInspect(d.dockerCtx, cid)
	if err != nil {
		return
	}

	enabled, found := c.Config.Labels[store.LABEL_ENABLED]
	if !found || enabled != "true" {
		return
	}

	name, found := c.Config.Labels[store.LABEL_NAME]
	if !found {
		return
	}

	d.log.Printf("removendo serviço %s\n", name)
	d.store.Remove(name)
}

func (d *Daemon) initDocker() (err error) {
	d.dockerCtx = context.Background()

	d.dockerClient, err = docker.NewClientWithOpts(
		docker.FromEnv, docker.WithAPIVersionNegotiation())

	if err != nil {
		d.log.Println("não consegui inicializar o cliente do docker\nerro:", err)
		return
	}

	return nil
}

func (d *Daemon) checkDockerConn() bool {
	_, err := d.dockerClient.Ping(d.dockerCtx)
	return err == nil
}

func (d *Daemon) processDockerEvents(c chan<- error) {
	evs, errs := d.dockerClient.Events(
		d.dockerCtx, types.EventsOptions{})

	for {
		select {
		case ev := <-evs:
			if ev.Type == "container" {
				switch ev.Status {
				case "create":
					d.AddService(ev.ID)

				case "die", "stop":
					d.RemoveService(ev.ID)
				}
			}

		case err := <-errs:
			c <- err
		}
	}
}
