package main

// ListingsSource defines interface for fetching data and push to DataChan
// Any error occurs during Fetch() will be returned via FetchError
type ListingsSource interface {
	DataChan() chan Listing
	Fetch()
	FetchError() error
}
