package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	// ErrMissingCSVFilePath is returned when missing filepath flag in case datasource=csv
	ErrMissingCSVFilePath = errors.New("missing filepath")
)

func configureApp() *cli.App {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Name = "Find Closest Listings"
	app.Flags = []cli.Flag{
		cli.Int64Flag{
			Name:  "top",
			Value: 5,
			Usage: "number of places to print out",
		},
		cli.StringFlag{
			Name:  "datasource, d",
			Value: "csv",
			Usage: "where to get list of coordinates, could be from csv, maybe sql in the future",
		},
		cli.StringFlag{
			Name:  "filepath",
			Usage: "path to csv file, only used when datasource=csv",
		},
		cli.Float64Flag{
			Name:  "long",
			Value: 4.478617,
			Usage: "Longitude of the place",
		},
		cli.Float64Flag{
			Name:  "lat",
			Value: 51.925146,
			Usage: "Longitude of the place",
		},
	}
	app.Action = applogic

	return app
}

func applogic(c *cli.Context) error {
	top := c.Int64("top")
	fromCoord := Coordinate{
		Lat: c.Float64("lat"),
		Lng: c.Float64("long"),
	}
	datasource, err := getDataSource(c)
	if err != nil {
		return err
	}

	go datasource.Fetch()

	nearestPlaces := findClosestListings(top, fromCoord, datasource.DataChan())
	if err = datasource.FetchError(); err != nil {
		return err
	}

	b := formatOutput(top, fromCoord, nearestPlaces)
	b.WriteTo(os.Stdout)

	return nil
}

func findClosestListings(top int64, from Coordinate, input chan Listing) []Listing {
	sortedList := NewSortedList(top)

	for listing := range input {
		distance := from.GreatCircleDistance(listing.Coordinate)
		sortedList.Insert(distance, listing)
	}

	sortedResults := sortedList.TopWithScore(-1)
	clostestPlaces := make([]Listing, len(sortedResults))
	for i, el := range sortedResults {
		clostestPlaces[i] = el.Value.(Listing)
		clostestPlaces[i].Distance = el.Score
	}

	return clostestPlaces
}

func formatOutput(top int64, fromCoord Coordinate, clostestPlaces []Listing) *bytes.Buffer {
	output := &bytes.Buffer{}
	fmt.Fprintf(output, "Top %d closest listings from coordinate %s:\n", top, fromCoord.String())

	for i, place := range clostestPlaces {
		fmt.Fprintf(output, "%d. id=%s coord=%s distance=%v\n", i+1, place.ID, place.Coordinate.String(), place.Distance)
	}

	return output
}

func getDataSource(c *cli.Context) (ListingsSource, error) {
	source := c.String("datasource")
	switch source {
	case "csv":
		filepath := c.String("filepath")
		if filepath == "" {
			return nil, ErrMissingCSVFilePath
		}
		return NewCSVSource(filepath), nil
	default:
		return nil, fmt.Errorf("datasource %s is unknown", source)
	}
}
