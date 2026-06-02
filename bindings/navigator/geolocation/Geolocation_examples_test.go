//go:build wasm

package geolocation

import "github.com/cookiengineer/gooey/bindings/navigator/geolocation"

func ExampleGeolocation_ClearWatch() {

	// import "github.com/cookiengineer/gooey/bindings/navigator/geolocation"

	console := console.GetConsole()
	geolocation := GetGeolocation()

	handler_id := geolocation.WatchPosition(func(position GeolocationPosition) {
		console.Info(position)
	}, func(err GeolocationError) {
		console.Error(err)
	})

	if handler_id != -1 {

		go func(geolocation *Geolocation) {
			time.Sleep(10 * time.Second)
			console.Log("Clear Position Watcher now")
			geolocation.ClearWatch(handler_id)
		}(geolocation)

	}

}

func ExampleGeolocation_GetCurrentPosition() {

	// import "github.com/cookiengineer/gooey/bindings/navigator/geolocation"

	console := console.GetConsole()
	geolocation := GetGeolocation()

	geolocation.GetCurrentPosition(func(position GeolocationPosition) {

		// GeolocationPosition{
		//   Coords: {
		//     Latitude:  13.37,
		//     Longitude: 13.37,
		//     Altitude:  1.337, // 0 if unavailable
		//     Accuracy:  12,    // 0 if unavailable
		//     Heading:   180,   // degrees clockwise from North
		//     Speed:     1.23,  // meters per second
		//   },
		//   Timestamp: time.Now(),
		// }

		console.Info(position)

	}, func(err GeolocationError) {
		console.Error(err)
	})

}

func ExampleGeolocation_WatchPosition() {

	// import "github.com/cookiengineer/gooey/bindings/navigator/geolocation"

	console := console.GetConsole()
	geolocation := GetGeolocation()

	handler_id := geolocation.WatchPosition(func(position GeolocationPosition) {
		console.Info(position)
	}, func(err GeolocationError) {
		console.Error(err)
	})

	if handler_id != -1 {

		go func(geolocation *Geolocation) {
			time.Sleep(10 * time.Second)
			console.Log("Clear Position Watcher now")
			geolocation.ClearWatch(handler_id)
		}(geolocation)

	}

}

