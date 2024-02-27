package primes

import (
	"fmt"
	"sync"
)

type Primes struct {
	upperLimit int
	lowerLimit int
	pchannel   chan int
	wg         *sync.WaitGroup
}

func GetPrimes(lowerLimit int, upperLimit int) Primes {
	return Primes{
		upperLimit: upperLimit,
		lowerLimit: lowerLimit,
		pchannel:   make(chan int),
		wg:         new(sync.WaitGroup),
	}
}

func (p *Primes) Run() {
	fmt.Println("Starting primes package")

	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		p.calculatePrimes()
	}()

	for i := 0; i < 1; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			p.printPrimes()
		}()
	}
}

func (p *Primes) Wait() {
	p.wg.Wait()
}

func (p *Primes) calculatePrimes() {
	defer close(p.pchannel)
	for i := p.lowerLimit; i <= p.upperLimit; i++ {
		p.pchannel <- i
	}

}

func (p *Primes) printPrimes() {

	for p := range p.pchannel {
		if isPrime(p) {
			fmt.Println(p)
		}
	}

}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
