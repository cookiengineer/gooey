package geolocation

type GeolocationPositionError int

const (
	GeolocationPositionErrorUnknown             GeolocationPositionError = 0
	GeolocationPositionErrorPermissionDenied    GeolocationPositionError = 1
	GeolocationPositionErrorPositionUnavailable GeolocationPositionError = 2
	GeolocationPositionErrorTimeout             GeolocationPositionError = 3
)

