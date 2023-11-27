package main

import (
	"github.com/panjf2000/ants/v2"
	"sync"
)

type Pool struct {
	mu          sync.Mutex
	ready       []func()
	gp          *ants.Pool
	doCh        chan struct{}
	concurrency int
}

func NewTaskPool(concurrency int) (p *Pool, err error) {
	p = &Pool{}
	p.gp, err = ants.NewPool(concurrency)
	p.ready = make([]func(), 0, concurrency)
	p.doCh = make(chan struct{})
	p.concurrency = concurrency

	go p.do()

	return p, err
}

func (p *Pool) Submit(task func()) {
	p.mu.Lock()
	p.ready = append(p.ready, task)
	p.mu.Unlock()
	p.wakeup()
}

func (p *Pool) Close() chan struct{} {
	ch := make(chan struct{})
	go func() {
		p.mu.Lock()
		defer p.mu.Unlock()
		for _, task := range p.ready {
			_ = p.gp.Submit(task)
		}
		p.gp.Release()
		close(ch)
	}()
	return ch
}

func (p *Pool) wakeup() {
	select {
	case p.doCh <- struct{}{}:
	default:
	}
}

func (p *Pool) do() {
	for {
		<-p.doCh
		p.mu.Lock()
		if len(p.ready) == 0 {
			p.mu.Unlock()
			continue
		}
		todo := p.ready
		p.ready = make([]func(), 0, p.concurrency)
		p.mu.Unlock()
		for _, fn := range todo {
			_ = p.gp.Submit(fn)
		}
	}
}
