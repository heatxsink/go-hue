# go-hue

[![Go Reference](https://pkg.go.dev/badge/github.com/heatxsink/go-hue.svg)](https://pkg.go.dev/github.com/heatxsink/go-hue)
[![Go Report Card](https://goreportcard.com/badge/github.com/heatxsink/go-hue)](https://goreportcard.com/report/github.com/heatxsink/go-hue)

A comprehensive Go library for interacting with Philips Hue lights via the Hue Hub HTTP API. This library provides a clean, idiomatic Go interface for controlling lights, groups, sensors, schedules, and rules.

## Demo

[![Alt Demo of go-hue unit tests](http://img.youtube.com/vi/3zMky_9xdJs/0.jpg)](http://www.youtube.com/watch?v=3zMky_9xdJs)

## Features

- **Lights**: Control individual lights (on/off, brightness, color, effects)
- **Groups**: Manage light groups and scenes
- **Sensors**: Read sensor data and manage sensor configurations
- **Schedules**: Create and manage scheduled actions
- **Rules**: Set up automation rules
- **Portal Discovery**: Automatic Hue bridge discovery
- **Configuration**: User management and bridge configuration

## Installation

```bash
go get github.com/heatxsink/go-hue
```

## Quick Start

### 1. Discover and Connect to Your Hue Bridge

First, run the discovery example to find your Hue bridge and create a user:

```bash
go run examples/discover.go
```

This will:
- Automatically discover your Hue bridge on the network
- Prompt you to press the link button on your bridge
- Create a new user and return your API key

### 2. Basic Usage

```go
package main

import (
    "fmt"
    "github.com/heatxsink/go-hue/lights"
    "github.com/heatxsink/go-hue/portal"
)

func main() {
    // Discover bridge
    bridges, err := portal.GetPortal()
    if err != nil {
        panic(err)
    }
    
    // Connect to lights API
    ll := lights.New(bridges[0].InternalIPAddress, "your-api-key-here")
    
    // Get all lights
    allLights, err := ll.GetAllLights()
    if err != nil {
        panic(err)
    }
    
    // Print light information
    for _, light := range allLights {
        fmt.Printf("Light %d: %s\n", light.ID, light.Name)
    }
    
    // Turn on a light
    state := lights.State{On: true, Brightness: 254}
    _, err = ll.SetLightState(1, state)
    if err != nil {
        panic(err)
    }
}
```

## API Reference

### Lights

```go
import "github.com/heatxsink/go-hue/lights"

// Create lights client
ll := lights.New(bridgeIP, apiKey)

// Get all lights
allLights, err := ll.GetAllLights()

// Get specific light
light, err := ll.GetLight(1)

// Set light state
state := lights.State{
    On:         true,
    Brightness: 200,
    Hue:        25500,
    Saturation: 200,
}
response, err := ll.SetLightState(1, state)
```

### Groups

```go
import "github.com/heatxsink/go-hue/groups"

// Create groups client
gg := groups.New(bridgeIP, apiKey)

// Get all groups
allGroups, err := gg.GetAllGroups()

// Set group state
state := lights.State{On: true, Brightness: 254}
response, err := gg.SetGroupState(1, state)
```

### Configuration

```go
import "github.com/heatxsink/go-hue/configuration"

// Create configuration client
cc := configuration.New(bridgeIP)

// Create new user (requires link button press)
response, err := cc.CreateUser("MyApp", "MyDevice")

// Get bridge configuration
config, err := cc.GetConfiguration(apiKey)
```

## Examples

The `examples/` directory contains working examples:

- `discover.go` - Bridge discovery and user creation
- `get-light-state.go` - List lights and groups, demonstrate state changes

Run examples:

```bash
# Discover bridge and create user
go run examples/discover.go

# List lights and groups (requires API key)
go run examples/get-light-state.go -key=your-api-key
```

## Testing

To run the tests, you'll need to set the following environment variables:

```bash
export HUE_TEST_USERNAME=your-api-key
export HUE_TEST_HOSTNAME=your-bridge-ip
```

You can obtain a whitelisted username (API key) by running `examples/discover.go`.

Then run the tests:

```bash
go test ./...
```

## Package Structure

- `configuration/` - Bridge configuration and user management
- `groups/` - Light group management
- `hue/` - Core API response types
- `lights/` - Individual light control
- `portal/` - Bridge discovery via portal API
- `rules/` - Automation rules
- `schedules/` - Scheduled actions  
- `sensors/` - Sensor data and management

## Important Notes

- **User Deletion**: As of API version 1.31, Philips/Hue has disabled the `deleteUser` function, and it will always return the error `unauthorized user`. To delete an app from your bridge, you must login to https://account.meethue.com/apps and manually delete using the website.

- **Rate Limiting**: The Hue API has rate limits. Avoid making too many requests in rapid succession.

- **Network Discovery**: The library uses both SSDP discovery and the Philips portal API for bridge discovery.

## Contributing

Issues and pull requests are always welcome! Please feel free to:

- Report bugs
- Suggest new features
- Submit pull requests
- Improve documentation

## License

This project is licensed under the MIT License - see the LICENSE file for details.


