package main

import (
	"fmt"

	fb "github.com/huandu/facebook"
)

var (
	appId       = "xxx"
	appSecret   = "xxx"
	accessToken = appId + "|" + appSecret
)

func getPlacesByLocation(latitude string, longitude string, distance string, query string) []Place {

	res, err := fb.Get("/search", fb.Params{
		"access_token": accessToken,
		"q":            query,
		"type":         "place",
		"center":       latitude + "," + longitude,
		"distance":     distance,
	})

	if err != nil {
		fmt.Println("Error during querying fb graph api", err.Error())
		return nil
	}

	var items []fb.Result

	err = res.DecodeField("data", &items)

	if err != nil {
		fmt.Println("Error while decoding fb graph api response", err)
		return nil
	}

	var places []Place
	for _, item := range items {
		fmt.Println(item["name"])
		if str, ok := item["name"].(string); ok {
			p := Place{Name: str}
			places = append(places, p)
		} else {
			fmt.Println("Can not parse fb result item to string.")
		}

	}
	return places
}
