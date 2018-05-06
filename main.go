package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func find(top int64, from Coordinate, input chan Listing) []Listing {
	list := NewSortedList(top)
	for listing := range input {
		distance := from.GreatCircleDistance(listing.Coordinate)
		listing.Distance = distance
		list.Insert(distance, listing)
	}

	sortedResults := list.TopWithScore(-1)
	clostestPlaces := make([]Listing, len(sortedResults))
	for i, el := range sortedResults {
		clostestPlaces[i] = el.Value.(Listing)
	}

	return clostestPlaces
}

func output(top int64, fromCoord Coordinate, clostestPlaces []Listing) *bytes.Buffer {
	output := &bytes.Buffer{}
	fmt.Fprintf(output, "Top %d closest places from coordinate %s:\n", top, fromCoord.String())

	for i, place := range clostestPlaces {
		fmt.Fprintf(output, "%d. Coord: %s, Distance: %v (km)\n", i+1, place.Coordinate.String(), place.Distance)
	}

	return output
}

func getDataSource(c *cli.Context) (ListingsSource, error) {
	switch c.String("datasource") {
	case "csv":
		filepath := c.String("filepath")
		if filepath == "" {
			return nil, errors.New("missing filepath")
		}
		return NewCSVSource(filepath), nil
	default:
		return nil, errors.New("unknow datasource")
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "Find Closest Places"
	app.Usage = "make an explosive entrance"
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
	app.Action = func(c *cli.Context) error {
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

		nearestPlaces := find(top, fromCoord, datasource.DataChan())
		if err = datasource.FetchError(); err != nil {
			return err
		}

		b := output(top, fromCoord, nearestPlaces)
		b.WriteTo(os.Stdout)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
