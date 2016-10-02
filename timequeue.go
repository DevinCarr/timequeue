package timequeue

import (
	"fmt"
	"net/http"
	//"net/url"
	"time"
)

type Store struct {
	expiry  time.Duration
	id      uint64
	address string
}

type TimeQueue struct {
	count uint64
}

// NewTimeQueue creates a new TimeQueue object
func NewTimeQueue() *TimeQueue {
	return &TimeQueue{
		count: 0,
	}
}

// wait will wait for the duration of the store and then call the
// provided url.
func wait(tq *TimeQueue, s *Store) {
	time.Sleep(s.expiry)
	resp, err := http.Get(s.address)
	//resp, err := http.PostForm(s.address,
	//url.Values{"value": {"123"}})

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	tq.Done()
}

// Push puts a new item in the queue with a time d in minutes
func (tq *TimeQueue) Push(d float64, id uint64, address string) {
	tq.count += 1
	fmt.Println("Put new item")
	go wait(tq, &Store{
		expiry:  time.Duration(d) * time.Minute,
		id:      id,
		address: address,
	})
}

// Done will remove the item in the queue and return the count of
// remaining items.
func (tq *TimeQueue) Done() uint64 {
	tq.count -= 1
	return tq.count
}

// Count returns the current amount of items in the queue
func (tq *TimeQueue) Count() uint64 {
	return tq.count
}
