package models

type RequestAddressSearch struct {
	Address string `json:"address"`
}

type RequestAddressGeocode struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type ResponseAddressInfo struct {
	Info []interface{} `json:"info"`
}
