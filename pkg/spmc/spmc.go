package spmc

import (
	"sync"
    "github.com/rs/zerolog/log"
)

//Single Producer Multi Consumer channel

type Spmc[T any] struct {
	chans map[int]chan T
	mutex sync.RWMutex
}

func NewSpmc[T any]() Spmc[T] {
	return Spmc[T]{
		chans: make(map[int]chan T),
		mutex: sync.RWMutex{},
	}
}

func (s *Spmc[T]) Subscribe(user_id int) <-chan T {
	ch := make(chan T, 10)
	s.mutex.Lock()
	s.chans[user_id] = ch
	s.mutex.Unlock()
    log.Info().Msgf("DEBUG:SPMC SUBSCRIBE %d\n", user_id)

	return ch
}

func (s *Spmc[T]) UnSubscribe(user_id int) {
	s.mutex.Lock()
	ch := s.chans[user_id]
	delete(s.chans, user_id)
	close(ch)
	s.mutex.Unlock()
	log.Info().Msgf("DEBUG:SPMC UNSUBSCRIBE %d\n", user_id)
}

/// Blocking call
func (s *Spmc[T]) Broadcast(item T) {
	s.mutex.RLock()
	for idx, ch := range s.chans {
		log.Debug().Msgf("DEBUG:SPMC BROADCAST BEFORE SENDING %d %+v %d\n", idx, item, len(ch))
		select { //select the savior....
		case ch <- item: //Getting blocked here !!!
		default:
		}

		log.Debug().Msgf("DEBUG:SPMC BROADCAST AFTER SENDING %d %+v %d\n", idx, item, len(ch))
	}
	s.mutex.RUnlock()
}
