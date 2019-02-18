# nestsensors

Golang package which provides a library and command-line utility for fetching Nest thermostat and sensor data not exposed through the official Nest API.

## Requirements

* A Nest account
* Go vX.X

## Usage

```go
package main

import (
    "jaytaylor.com/nestsensors"
)

func main() {
}
```

Also see [examples](...).

## Background and motivation

The [Temperature sensor API was first requested in April 2017](https://nestdevelopers.io/t/temperature-sensor-api/1317/1), and despite promises to deliver, Nest has yet to do so.

Jeff Williams created a go package called [nest-scrape](https://github.com/jeffwilliams/nest-scrape).  I like it, however it depends on headless Firefox annd I couldn't get it working on Windows.  `nestsensors` has no external dependencies.
