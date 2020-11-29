package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {

	feeds, err := RetrieveFeeds() // get the feeds!
	if err != nil {
		log.Fatal(err)
	}

	// Create an unbuffered channel to receive match results
	resultsChannel := make(chan *Result)

	// Setup a wait group so we can process all the feeds
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for a while
	// they process the individual feeds
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results
	for _, feed := range feeds {
		// Retrieve a matcher for the search
		matcher, exists := matchers[feed.Type]

		if !exists {
			matcher = matchers["default"]
		}

		// Launch the goroutine to perform the search
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, resultsChannel)
			waitGroup.Done()
		} (matcher, feed)


	}

	// Launch a goroutine to monitor when all the work is done
	go func () {
		/// wait for everything to be processed
		waitGroup.Wait()

		// Close the channel to signal to the display function
		// that we can exit the program
		close(resultsChannel)
	}()

	// Start displaying results as they are available and
	// return after the final result is displayed
	Display(resultsChannel)

}