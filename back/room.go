package main

import (
	"errors"
	"sync"
	"time"
)

type Room struct {
	Expire  time.Time
	QR      string
	Clients []*Client
	lock    sync.Mutex
	Alive   bool
}

type Client struct {
	Signal chan (int)
	Report chan (error)
	Alive  bool
}

func (r *Room) UpdateQRCode(qr string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if !r.Alive {
		return
	}
	if qr == r.QR {
		return
	}
	r.QR = qr
	for _, c := range r.Clients {
		if !c.Alive {
			continue
		}
		c.Signal <- 0
	}
	for _, c := range r.Clients {
		if !c.Alive {
			continue
		}
		err := <-c.Report
		if err == nil {
			continue
		}
		c.Alive = false
	}
}

func (r *Room) AddClient(c *Client) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if !r.Alive {
		return errors.New("room not alive")
	}
	r.Clients = append(r.Clients, c)
	return nil
}

func (r *Room) CloseRoom() {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.Alive = false
	for _, c := range r.Clients {
		if !c.Alive {
			continue
		}
		c.Signal <- 1
	}
}
