package main

import "github.com/cookiengineer/gooey/pkg"
import "github.com/cookiengineer/gooey/pkg/navigator/geolocation"
import "strconv"
import "time"

func main() {

	element_latitude := gooey.Document.QuerySelector("#latitude")
	element_longitude := gooey.Document.QuerySelector("#longitude")
	element_altitude := gooey.Document.QuerySelector("#altitude")
	element_accuracy := gooey.Document.QuerySelector("#accuracy")
	element_error := gooey.Document.QuerySelector("#error")

	geolocation.Geolocation.GetCurrentPosition(func(position geolocation.GeolocationPosition) {

		element_latitude.SetInnerHTML(strconv.FormatFloat(position.Coords.Latitude, 'g', -1, 64))
		element_longitude.SetInnerHTML(strconv.FormatFloat(position.Coords.Longitude, 'g', -1, 64))
		element_altitude.SetInnerHTML(strconv.FormatFloat(position.Coords.Altitude, 'g', -1, 64))
		element_accuracy.SetInnerHTML(strconv.FormatFloat(position.Coords.Accuracy, 'g', -1, 64) + " meters")

	}, func(err geolocation.GeolocationPositionError) {

		switch err {
		case geolocation.GeolocationPositionErrorUnknown:
			element_error.SetInnerHTML("Unknown Error")
			break
		case geolocation.GeolocationPositionErrorPermissionDenied:
			element_error.SetInnerHTML("Permission Denied")
			break
		case geolocation.GeolocationPositionErrorPositionUnavailable:
			element_error.SetInnerHTML("Position Unavailable")
			break
		case geolocation.GeolocationPositionErrorTimeout:
			element_error.SetInnerHTML("Timeout")
			break
		}

	})

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
