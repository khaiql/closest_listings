package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchCSVSource(t *testing.T) {
	csvData := `"id","lat","lng"
382582,37.1768672,-3.608897
482365,52.36461880000000235,4.93169289999999982
28403,51.9245615,4.492032399999999
invalidRecord
20213,invalidLat,4.00123
20211,10.87113,invalidLng
`
	filename := "testData.csv"
	_ = ioutil.WriteFile(filename, []byte(csvData), os.ModePerm)
	defer os.Remove(filename)

	t.Run("ExistingFile", func(t *testing.T) {
		csvSource := NewCSVSource(filename)
		go csvSource.Fetch()

		listings := []Listing{}
		for listing := range csvSource.DataChan() {
			listings = append(listings, listing)
		}

		assert.NoError(t, csvSource.FetchError())
		assert.Len(t, listings, 3)
	})

	t.Run("FileNotFound", func(t *testing.T) {
		csvSource := NewCSVSource("notfound.csv")
		go csvSource.Fetch()

		_, ok := <-csvSource.DataChan()
		assert.False(t, ok)
		assert.Error(t, csvSource.FetchError())
	})
}
