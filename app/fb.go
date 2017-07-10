package main

import (
	"fmt"

	fb "github.com/huandu/facebook"
)

var (
	appId       = "xxx"
	appSecret   = "xxxxx"
	accessToken = appId + "|" + appSecret
	globalApp   *fb.App
	session     *fb.Session
)

func init() {
	globalApp = fb.New(appId, appSecret)
	globalApp.RedirectUri = "http://localhost:8000/"
	session = globalApp.Session(accessToken)

	err := session.Validate()
	if err != nil {
		fmt.Println("Session validation error: ", err.Error())
	}
}

func getPlacesByLocation(latitude string, longitude string, distance string, query string) []Place {
	var places []Place

	response, err := session.Get("/search", fb.Params{
		"access_token": accessToken,
		"q":            query,
		"type":         "place",
		"center":       latitude + "," + longitude,
		"distance":     distance,
	})

	if err != nil {
		fmt.Println("Fb Request error: ", err.Error())
	}

	if paging, err := response.Paging(session); err != nil {
		fmt.Println("Fb paging error: ", err.Error())
	} else {
		for paging.HasNext() {
			results := paging.Data()
			for _, item := range results {
				if str, ok := item["name"].(string); ok {
					p := Place{Name: str}
					places = append(places, p)
				} else {
					fmt.Println("Can not parse fb result item to string.")
				}
			}
			if _, err := paging.Next(); err != nil {
				fmt.Println("Paging error: ", err.Error())
			}
		}
	}

	return places
}
