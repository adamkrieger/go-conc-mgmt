package balancer

type RoundRobin interface {
	NextHost() <-chan string
}

type roundRobin struct {
	nextChan  chan string
	endpoints []string
	curIdx    int
}

//StartRoundRobin - Constructor and Run in one!
func StartRoundRobin(endpoints []string) RoundRobin {
	retObj := &roundRobin{
		nextChan:  make(chan string),
		endpoints: endpoints,
		curIdx:    0,
	}

	go retObj.generateNextHosts()

	return retObj
}

//NextHost - Threadsafe way to get next host
func (rr *roundRobin) NextHost() <-chan string {
	return rr.nextChan
}

func (rr *roundRobin) generateNextHosts() {
	for {
		next := rr.endpoints[rr.curIdx]

		rr.curIdx++
		rr.curIdx = rr.curIdx % len(rr.endpoints)

		rr.nextChan <- next
	}
}
