package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

// NewCSVSource constructs a ListingsSource instance that can read data from csv file
func NewCSVSource(filepath string) ListingsSource {
	return &csvSource{
		filepath: filepath,
		outchan:  make(chan Listing),
	}
}

type csvSource struct {
	filepath string
	outchan  chan Listing
	fetchErr error
}

// DataChan returns channel for reading Listing obtained through Fetch()
func (s *csvSource) DataChan() chan Listing {
	return s.outchan
}

// FetchError returns error happened during Fetch
func (s *csvSource) FetchError() error {
	return s.fetchErr
}

// Fetch reads data from csv file, line by line and publish to data channel.
// The channel will be closed after finished reading the file or any error happened
func (s *csvSource) Fetch() {
	defer close(s.outchan)

	file, err := os.Open(s.filepath)
	if err != nil {
		s.fetchErr = err
		return
	}

	csvReader := csv.NewReader(file)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		// ignore malformed line
		if len(record) < 3 {
			continue
		}
		lat, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			continue
		}
		lng, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			continue
		}
		listing := Listing{
			ID:         record[0],
			Coordinate: Coordinate{Lat: lat, Lng: lng},
		}

		s.outchan <- listing
	}

	return
}
