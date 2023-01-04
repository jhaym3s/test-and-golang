package sync

type Counter struct{
	number int
}


func (c *Counter) Inc()  {
	c.number++
}

func (c *Counter) Value()int{
  return c.number
}