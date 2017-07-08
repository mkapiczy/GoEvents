package main

import (
	"fmt"

	fb "github.com/huandu/facebook"
)

var (
	appId       = "xxxx"
	appSecret   = "xxxx"
	accessToken = appId + "|" + appSecret
)

func getPlacesByLocation(latitude string, longitude string, distance string, query string) []Place {

	res, err := fb.Get("/search", fb.Params{
		"access_token": accessToken,
		"type":         "place",
		"q":            query,
		"center":       latitude + "," + longitude,
		"distance":     distance,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	var items []fb.Result

	err = res.DecodeField("data", &items)

	if err != nil {
		fmt.Printf("An error has happened %v", err)
		return nil
	}

	var places []Place
	for _, item := range items {
		fmt.Println(item["name"])
		if str, ok := item["name"].(string); ok {
				p := Place{Name: str}
				places = append(places, p)
		} else {
			fmt.Println("Not a string!")
		}

	}
	return places
}
