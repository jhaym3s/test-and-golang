package sync

import "sync"

type Counter struct{
	mu    sync.Mutex
	number int
}


func (c *Counter) Inc()  {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.number++
}

func (c *Counter) Value()int{
  return c.number
}