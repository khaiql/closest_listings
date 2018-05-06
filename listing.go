package main

// Listing defines struct of a listing that hold data obtained from ListingSource
type Listing struct {
	ID         string
	Coordinate Coordinate
	Distance   float64
}
