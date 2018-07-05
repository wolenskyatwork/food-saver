package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"time"
)

type DietDay struct {
	Date    time.Time `json:"date"`
	Choice	string `json:"choice"`
}

var dietDays []DietDay

func getDietDaysHandler(w http.ResponseWriter, r *http.Request) {
	//Convert the "birds" variable to json
	dietDaysBytes, err := json.Marshal(dietDays)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of birds to the response
	w.Write(dietDaysBytes)
}

func createDietDayHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Bird
	// bird := Bird{}
	fmt.Println("you hit a thing")
	// Create a new instance of DietDay
	dietDay := DietDay{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dietDay.Choice = r.Form.Get("choice")
	dietDay.Date = time.Now() // r.Form.Get("date") // this probably has to be typed

	// Get the information about the bird from the form info
	//bird.Species = r.Form.Get("species")
	//bird.Description = r.Form.Get("description")

	// Append our existing list of birds with a new entry
	//birds = append(birds, bird)
	dietDays = append(dietDays, dietDay)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	// I don't think I need to redirect for now?
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
