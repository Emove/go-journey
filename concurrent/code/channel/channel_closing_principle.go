package channel

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const MaxRandomNumber = 100000
const NumOfReceivers = 100

var wgReceivers = sync.WaitGroup{}

func MultiReceiver() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	wgReceivers.Add(NumOfReceivers)

	dataCh := make(chan int, 100)

	// the sender
	go func() {
		for {
			if value := rand.Intn(MaxRandomNumber); value == 0 {
				// the only sender can close the channel safely
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	// receivers
	for i := 0; i < NumOfReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			// receive values until dataCh is closed and
			// the value buffer queue of dataCh is empty
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}

func MultiSenderTest() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	wgReceivers.Add(1)

	dataCh := make(chan int, NumOfReceivers)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel
	// Its sender is the receiver of channel dataCh
	// Its receiver are the senders of channel dataCh

	// senders
	for i := 0; i < NumOfReceivers; i++ {
		go func() {
			for {
				value := rand.Intn(MaxRandomNumber)
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}()

		// the receiver
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				if value == MaxRandomNumber-1 {
					// the receiver of the dataCh channel is
					// also the sender if the stopCh channel
					// It is safe to close the stop channel here
					close(stopCh)
					return
				}
				log.Println(value)
			}
		}()

		wgReceivers.Wait()
	}
}

func MultiSenderAndReceiverTest() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const NumOfReceiver = 10
	const NumOfSender = 1000

	wgReceivers.Add(NumOfReceiver)

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	toStop := make(chan string, 1)
	// the channel toStop is used to notify the moderator
	// to close the additional signal channel (stopCh)
	// Its senders are any senders and receivers of dataCh
	// Its receiver is the moderator goroutine shown below

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop
		// part of the trick used to notify the moderator
		// to close the additional signal channel
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumOfSender; i++ {
		go func(id string) {
			for {
				value := rand.Intn(MaxRandomNumber)
				if value == 0 {
					// here, a trick is used to notify the moderator
					// to close the additional signal channel
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// the first select here is to try to exit
				// the goroutine as early as possible
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumOfReceiver; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				// same as senders, the first select here is to
				// try to exit the goroutine as early as possible
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == MaxRandomNumber-1 {
						// the same trick is used to notify the moderator
						// to close the additional signal channel
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}
					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}
