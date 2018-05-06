package main

import (
	"flag"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

func TestGetDataSource(t *testing.T) {
	t.Run("valid csv source", func(t *testing.T) {
		flagset := flag.NewFlagSet("test", flag.ContinueOnError)
		flagset.String("filepath", "something.csv", "")
		flagset.String("datasource", "csv", "")

		ctx := cli.NewContext(nil, flagset, nil)

		source, err := getDataSource(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, source)
	})

	t.Run("csv source but missing filepath", func(t *testing.T) {
		flagset := flag.NewFlagSet("test", flag.ContinueOnError)
		flagset.String("datasource", "csv", "")

		ctx := cli.NewContext(nil, flagset, nil)

		source, err := getDataSource(ctx)
		assert.Equal(t, ErrMissingCSVFilePath, err)
		assert.Nil(t, source)
	})

	t.Run("Unknown source", func(t *testing.T) {
		flagset := flag.NewFlagSet("test", flag.ContinueOnError)
		flagset.String("datasource", "xml", "")

		ctx := cli.NewContext(nil, flagset, nil)
		source, err := getDataSource(ctx)
		assert.Nil(t, source)
		assert.EqualError(t, err, "datasource xml is unknown")
	})
}

func TestFindClosestListing(t *testing.T) {
	from := Coordinate{51.925146, 4.478617}
	listings := []Listing{
		{
			ID:         "382582",
			Coordinate: Coordinate{37.1768672, -3.608897},
		},
	}

	expectedResults := []Listing{
		{
			ID:         listings[0].ID,
			Coordinate: listings[0].Coordinate,
			Distance:   from.GreatCircleDistance(listings[0].Coordinate),
		},
	}

	inputChan := make(chan Listing)
	go func() {
		defer close(inputChan)
		for _, l := range listings {
			inputChan <- l
		}
	}()

	results := findClosestListings(1, from, inputChan)
	assert.EqualValues(t, expectedResults, results)
}

func TestFormatOutput(t *testing.T) {
	top := 1
	from := Coordinate{51.925146, 4.478617}
	listing := Listing{
		ID:         "1",
		Coordinate: Coordinate{37.1768672, -3.608897},
	}
	listing.Distance = from.GreatCircleDistance(listing.Coordinate)
	expectedMsg := `Top 1 closest listings from coordinate ` + from.String() + ":\n"
	expectedMsg = expectedMsg + fmt.Sprintf("1. id=1 coord=%s distance=%v\n", listing.Coordinate.String(), listing.Distance)

	b := formatOutput(int64(top), from, []Listing{listing})
	assert.Equal(t, expectedMsg, b.String())
}
