package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	log.Println("Go Time Zone")
	currentTime := time.Now()
	// currentTimeUTC := time.Now().UTC()

	// Load "Europe/Berlin" timezone
	locBerlin, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	locLocal, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}
	currentTimeLocal := currentTime.In(locLocal)
	log.Println("currentTimeLocal:", currentTimeLocal)

	// Get the current time in Berlin
	currentTimeBerlin := currentTime.In(locBerlin)

	// Calculate the offset from UTC
	_, offset := currentTimeBerlin.Zone()
	offsetHours := offset / 3600          // Convert offset from seconds to hours
	offsetMinutes := (offset % 3600) / 60 // Convert remaining seconds to minutes

	// Get the time zone of the current local device
	// loc, _ := currentTime.Zone()
	location := currentTime.Location()
	log.Println("location :", location)

	timeZone, offset := currentTime.In(location).Zone()

	// Print the time zone and offset
	log.Println("Local device Time Zone:", timeZone)
	log.Println("Local device Offset:", offset)

	// Print the offset
	fmt.Printf("Offset from UTC for Berlin time: %02d:%02d\n", offsetHours, offsetMinutes)

	// Extract year, month, and day
	year := currentTime.Year()
	month := currentTime.Month()
	day := currentTime.Day()

	// get current date
	date := fmt.Sprint(day) + "-" + fmt.Sprint(month) + "-" + fmt.Sprint(year)
	timeZoneOffSet := fmt.Sprint(offsetHours) + fmt.Sprint(offsetMinutes)
	// get the offset of the timezone
	timeFormat := date + " " + "15:00:00" + " " + timeZoneOffSet

	log.Println("timeFormat: ", timeFormat)
	log.Println("year: ", year)
	log.Println("month: ", month)
	log.Println("day: ", day)

	// Extract the layout from the current time
	// timeLayout := currentTime.Format("02-01-2006 15:04:05 -0700")
	timeLayout := "02-01-2006 15:04:05 -0700"
	// timeLayout := currentTime.Format(timeFormat)
	log.Println("timeLayout===>", timeLayout)

	// Parse the current time using the extracted layout
	givenTime, err := time.Parse(timeLayout, currentTime.Format(timeLayout))
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	fmt.Printf("givenTime time in locally: %s\n", givenTime)
	// Convert the given time to Berlin time
	berlinTime := givenTime.In(locBerlin)

	// Print the times for comparison
	fmt.Printf("Time in code (Berlin Time): %s\n", berlinTime)
	fmt.Printf("Current time in Berlin: %s\n", currentTimeBerlin)

	// Check if they are equal
	if berlinTime.Equal(currentTimeBerlin) {
		fmt.Println("The time in the code matches the current time in Berlin.")
	} else {
		fmt.Println("The time in the code does not match the current time in Berlin.")
	}

}

// ===OUTPUT===

// 2024/05/08 09:59:41 Go Time Zone
// Offset from UTC for Berlin time: 02:00
// 2024/05/08 09:59:41 timeFormat:  8-May-2024 15:00:00 20
// 2024/05/08 09:59:41 year:  2024
// 2024/05/08 09:59:41 month:  May
// 2024/05/08 09:59:41 day:  8
// Time in code (Berlin Time): 0000-01-01 00:54:09 +0053 LMT
// Current time in Berlin: 2024-05-08 08:59:41.021492047 +0200 CEST
// The time in the code does not match the current time in Berlin.
