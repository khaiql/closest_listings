# Find top closest listings
[![Build Status](https://travis-ci.org/khaiql/closest_listings.svg?branch=master)](https://travis-ci.org/khaiql/closest_listings)

## Prerequisites
1. Having go (1.10.x) installed and set up
1. Having `dep` installed: https://golang.github.io/dep/docs/installation.html

## Getting started

1. `git clone git@github.com:khaiql/closest_listings.git $GOPATH/src/github.com/khaiql/closest_listings`
1. `cd $GOPATH/src/github.com/khaiql/closest_listings`
1. `dep ensure`
1. `go build`
1. `./closest_listings -filepath data/geoData.csv`

## Flags

- --top [5]
- --datasource [csv]
- --filepath 
- --long [4.478617]
- --lat [51.925146]
