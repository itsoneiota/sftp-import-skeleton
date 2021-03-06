package importer

import "time"

// Poller polls an importer.
type Poller struct {
	i Importer
	w Worker
}

// NewPoller returns a Poller to poll the given importer.
func NewPoller(i Importer, w Worker) *Poller {
	return &Poller{i, w}
}

// Start starts the poller polling. Blocking call.
func (p *Poller) Start(interval time.Duration) {
	for {
		p.i.Poll(p.w)
		time.Sleep(interval)
	}
}
