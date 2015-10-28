package main

type Counter struct {
	ch chan uint64 //
}

func RunCounter(next uint64, counter chan Counter) {
	count := next
	for {
		select {
		case c := <-counter:
			count++
			c.ch <- count
		}
	}
}
