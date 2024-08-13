package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
)

// go routine to listen to events change the url for app
func timingListener(data chan []byte) {
	resp, err := http.Get("http//localhost:8000/timing/subscribe/irwindale")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// create a new buffered reader to read the stream
	reader := bufio.NewReader(resp.Body)

	// loop to continuously read from the stream
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			break
		}
		data <- []byte(line)
	}

}
func main() { // Create a new request to the SSE endpoint
	msg := make(chan []byte, 10)
	// call the timinglisterner  function async
	go timingListener(msg)

	/// in the case of the first message unmarshal it into session Into this will be the state for the session data
	// in the case of a newly started session the fields will be empty
	// in the case of a already started session some fields might be filed or all fields might be filled
	// sett the by possition variable to true by default
	// will be used to determin where we are sorting by race postion or practice/qualifing time
	byPos := true
	var sd SessionInfo
	data := <-msg
	json.Unmarshal(data, &sd)
	fmt.Println(sd)
	// start listening for other messages
	// create pointer to session data as not not reset the other values
	currentSession := sd
	fmt.Println(currentSession)
	// each time anoter message is recived
	for updateData := range msg {
		// use the coutom json unmarshaling to unmashal the data
		var tm TimingMessage
		tm.Unmarshal(updateData)
		// switch case based on the message type
		switch tm.Type {
		// in the case of a hearbeat message update the data for the
		case "F":
			// set race status to a new heartBead
			currentSession.RaceStatus = tm.Data.(Heartbeat)
		case "COMP":
			// add a new competitor to the session data put the competior at last for now
			comp := tm.Data.(CompInfo)
			currentSession.RacerData = append(currentSession.RacerData, RacerData{Number: comp.Number, RegistrationNumber: comp.RegistrationNumber})
		case "G":
			raceInfo := tm.Data.(RaceInfo)
			// find the racer with the registration number using brute force search use other alg for prod
			for i := range currentSession.RacerData {
				if currentSession.RacerData[i].RegistrationNumber == raceInfo.RegistrationNumber {
					currentSession.RacerData[i].RacePosition = raceInfo
				}
				// sort the racer based on race position if selected race position sorting else leave alone
				if byPos == true {

				}
			}
			// sort the racer based on race position if selected race position sorting else leave alone
			//

			// case H
			// find the racer with the same registation number
			// then update their practice qualifying position
			// if the sort by pos var is false then sort by the practice qualifing position

			// case J
			// find the racer with the same registration number
			// add the passing info to their passes

			// case I
			//clear the session information and reset  the sessiondata to be default with no values

			// continue to handle the other types of messages

		}
		// ie update the state after a change to the session
		fmt.Println(currentSession)

	}

}
