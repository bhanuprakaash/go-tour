package main

import (
	"net/url"
	"sync"
)

type ServerPool struct {
	backends []*url.URL
	current  uint64
	mu       sync.Mutex
}

func (s *ServerPool) AddToBackend(serverUrl *url.URL) {
	s.backends = append(s.backends, serverUrl)
}

func (s *ServerPool) GetNextIndex() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.current++

	return int(s.current % uint64(len(s.backends)))
}

func (s *ServerPool) GetNextPeer() *url.URL {
	next := s.GetNextIndex()
	return s.backends[next]
}
