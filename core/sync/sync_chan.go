// Copyright The Jet authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sync

// Event Auto reset event.
type Event chan struct{}

func NewEvent() Event {
	return make(chan struct{}, 1)
}

func (e Event) Set() {
	select {
	case e <- struct{}{}:
	default:
	}
}

func (e Event) R() EventR {
	return EventR((chan struct{})(e))
}

// EventR You can determine whether event setted through EventR.
type EventR <-chan struct{}

// DoneChan You can notify something done through DoneChan.SetDone().
type DoneChan chan struct{}

func NewDoneChan() DoneChan {
	return make(chan struct{})
}

func (d DoneChan) SetDone() {
	defer func() { recover() }()
	select {
	case <-d:
	default:
		close(d)
	}
}

func (d DoneChan) R() DoneChanR {
	return (chan struct{})(d)
}

// DoneChanR You can determine whether something done through DoneChanR.Done().
type DoneChanR <-chan struct{}

func (d DoneChanR) Done() bool {
	select {
	case <-d:
		return true
	default:
		return false
	}
}

// Semaphore can be used to limit access to multiple resources.
type Semaphore chan struct{}

func NewSemaphore(n int) Semaphore {
	if n <= 0 {
		panic("invalid n")
	}
	return make(chan struct{}, n)
}

// Acquire n resources.
//
// s <- e
func (s Semaphore) Acquire(n int) {
	if n > cap(s) {
		panic("invalid n")
	}
	e := struct{}{}
	for i := 0; i < n; i++ {
		s <- e
	}
}

// Release n resources.
//
// e := <-s
func (s Semaphore) Release(n int) {
	if n > cap(s) {
		panic("invalid n")
	}
	for i := 0; i < n; i++ {
		<-s
	}
}
