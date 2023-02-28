package main_utils

func JoinChannelEntry(channelEntry ...<-chan string) <-chan string {
	resultChannel := make(chan string)
	go func() {
		for {

			for _, value := range channelEntry {
				select {
				case message := <-value:
					resultChannel <- message
				}
			}

		}
	}()
	return resultChannel
}

func JoinChannels(channels []<-chan interface{}) <-chan interface{} {
	resultChannel := make(chan interface{})
	go func() {
		for {

			for _, value := range channels {
				select {
				case message := <-value:
					resultChannel <- message
				}
			}

		}
	}()
	return resultChannel
}
