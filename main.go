package main

import (
	"fmt"
	"net/http"
	"time"
)

func countdownHandler(w http.ResponseWriter, r *http.Request) {
	// Zielzeitpunkt festlegen
	end := time.Date(2024, time.February, 4, 12, 7, 0, 0, time.UTC)
	now := time.Now()

	// Zeitdifferenz berechnen
	diff := end.Sub(now)

	// Tage, Stunden, Minuten und Sekunden berechnen
	days := int(diff.Hours() / 24)
	hours := int(diff.Hours()) % 24
	minutes := int(diff.Minutes()) % 60
	seconds := int(diff.Seconds()) % 60

	// Countdown als String formatieren und ausgeben
	countdown := fmt.Sprintf("%d Tage %02d:%02d:%02d", days, hours, minutes, seconds)
	fmt.Fprintf(w, countdown)
}

func main() {
	// Statischer Dateihandler
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handler für den Countdown und die Hauptseite
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/countdown", countdownHandler)

	// Starte den Server
	http.ListenAndServe(":8080", nil)
}
