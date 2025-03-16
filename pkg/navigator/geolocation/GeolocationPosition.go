//go:build wasm

package geolocation

import "syscall/js"
import "time"

type GeolocationPosition struct {
	Coords struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Altitude  float64 `json:"altitude"`
		Accuracy  float64 `json:"accuracy"`
		Heading   float64 `json:"heading"`
		Speed     float64 `json:"speed"`
	} `json:"coords"`
	Timestamp time.Time `json:"timestamp"`
}

func ToGeolocationPosition(value js.Value) GeolocationPosition {

	var position GeolocationPosition

	coords := value.Get("coords")

	if !coords.IsNull() && !coords.IsUndefined() {

		latitude := coords.Get("latitude")

		if !latitude.IsUndefined() && !latitude.IsNull() {
			position.Coords.Latitude = latitude.Float()
		}

		longitude := coords.Get("longitude")

		if !longitude.IsUndefined() && !longitude.IsNull() {
			position.Coords.Longitude = longitude.Float()
		}

		altitude := coords.Get("altitude")

		if !altitude.IsUndefined() && !altitude.IsNull() {
			position.Coords.Altitude = altitude.Float()
		}

		accuracy := coords.Get("accuracy")

		if !accuracy.IsUndefined() && !accuracy.IsNull() {
			position.Coords.Accuracy = accuracy.Float()
		}

		heading := coords.Get("heading")

		if !heading.IsUndefined() && !heading.IsNull() && !heading.IsNaN() {
			position.Coords.Heading = heading.Float()
		}

		speed := coords.Get("speed")

		if !speed.IsUndefined() && !speed.IsNull() && !speed.IsNaN() {
			position.Coords.Speed = speed.Float()
		}

	}

	timestamp := value.Get("timestamp")

	if !timestamp.IsNull() && !timestamp.IsUndefined() {
		position.Timestamp = time.Unix(int64(timestamp.Int()/1000), 0)
	}

	return position

}
