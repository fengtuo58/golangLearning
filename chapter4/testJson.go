package main

import (
	"fmt"
	"encoding/json"
)

type Movie struct {
	Title string
	Year  int `json:"release"`
	Color bool `json:"color, omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "kkkhfhhf", Year: 9, Color :false,
		Actors:[]string{"fengtuo", "tuofeng"},
	},
	{Title: "kkkhfhhf", Year: 9, Color :false,
		Actors:[]string{"fengtuo"},
	},
	{Title:"BUllitt", Year: 1968, Color:true,
		Actors:[]string{"steve McQueen", "Jacqueline Bisset"},
	},
}

func main() {
	//	data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil{
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("data:%s\n", data)

	var slic []struct{Title string}
	err = json.Unmarshal(data, &slic)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(slic)
}

