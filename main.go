package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Event struct {
	Title       string
	Description string
	StartDate   string
	EndDate     string
	Country     string
	Category    string // Nouveau champ pour la catégorie
}

type Battle struct {
	Name        string
	Description string
	Location    string
	Latitude    float64
	Longitude   float64
	Date        string
}

type Biography struct {
	Name        string `json:"name"`
	Biography   string `json:"biography"`
	Anecdotes   string `json:"anecdotes"`
	DateOfBirth string `json:"date_of_birth"`
	DateOfDeath string `json:"date_of_death"`
	References  string `json:"references"`
}

type Technology struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"` // Ajout du champ Image
}

// Structure pour les stratégies militaires
type Strategy struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func dbConn() (db *sql.DB) {
	dbDriver := "sql7.freemysqlhosting.net"
	dbUser := "sql7734553"
	dbPass := "PTKXi89sf1"
	dbName := "sql7734553"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/timeline", timeline)
	http.HandleFunc("/battles", battles)
	http.HandleFunc("/api/battles", battlesAPI)
	http.HandleFunc("/biographies", biographies)
	http.HandleFunc("/api/biographies", biographiesAPI)
	http.HandleFunc("/technology", technologyHandler)
	http.HandleFunc("/api/technologies", technologiesAPI)
	http.HandleFunc("/api/strategies", strategiesAPI)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func timeline(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	// Fetch timeline events
	rows, err := db.Query("SELECT title, description, start_date, end_date, country, category FROM timeline_events ORDER BY start_date")
	if err != nil {
		log.Fatal(err)
	}
	var events []Event
	for rows.Next() {
		var event Event
		err = rows.Scan(&event.Title, &event.Description, &event.StartDate, &event.EndDate, &event.Country, &event.Category)
		if err != nil {
			log.Fatal(err)
		}
		events = append(events, event)
	}

	templates.ExecuteTemplate(w, "timeline.html", events)
}

func battles(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	// Fetch battles
	rows, err := db.Query("SELECT name, description, location, latitude, longitude, date FROM battles")
	if err != nil {
		log.Fatal(err)
	}
	var battles []Battle
	for rows.Next() {
		var battle Battle
		err = rows.Scan(&battle.Name, &battle.Description, &battle.Location, &battle.Latitude, &battle.Longitude, &battle.Date)
		if err != nil {
			log.Fatal(err)
		}
		battles = append(battles, battle)
	}

	templates.ExecuteTemplate(w, "battles.html", battles)
}

func battlesAPI(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	// Fetch battles
	rows, err := db.Query("SELECT name, description, location, latitude, longitude, date FROM battles")
	if err != nil {
		log.Fatal(err)
	}
	var battles []Battle
	for rows.Next() {
		var battle Battle
		err = rows.Scan(&battle.Name, &battle.Description, &battle.Location, &battle.Latitude, &battle.Longitude, &battle.Date)
		if err != nil {
			log.Fatal(err)
		}
		battles = append(battles, battle)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(battles)
}

func biographies(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "biographies.html", nil)
}

// Handler API pour récupérer les biographies
func biographiesAPI(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	// Met à jour la requête SQL pour inclure les références
	rows, err := db.Query("SELECT name, biography, anecdotes, date_of_birth, date_of_death, referencesbio FROM biographies")
	if err != nil {
		log.Fatal(err)
	}
	var bios []Biography
	for rows.Next() {
		var bio Biography
		err = rows.Scan(&bio.Name, &bio.Biography, &bio.Anecdotes, &bio.DateOfBirth, &bio.DateOfDeath, &bio.References)
		if err != nil {
			log.Fatal(err)
		}
		bios = append(bios, bio)
	}

	// Envoyer les données sous forme de JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bios)
}

func technologyHandler(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	// Récupérer les technologies militaires
	techRows, err := db.Query("SELECT id, name, description, image FROM technologies")
	if err != nil {
		log.Fatal(err)
	}
	defer techRows.Close()

	var technologies []Technology
	for techRows.Next() {
		var tech Technology
		err = techRows.Scan(&tech.ID, &tech.Name, &tech.Description, &tech.Image)
		if err != nil {
			log.Fatal(err)
		}
		technologies = append(technologies, tech)
	}

	// Récupérer les stratégies militaires
	stratRows, err := db.Query("SELECT id, name, description FROM strategies")
	if err != nil {
		log.Fatal(err)
	}
	defer stratRows.Close()

	var strategies []Strategy
	for stratRows.Next() {
		var strat Strategy
		err = stratRows.Scan(&strat.ID, &strat.Name, &strat.Description)
		if err != nil {
			log.Fatal(err)
		}
		strategies = append(strategies, strat)
	}

	// Passer les données à la template
	data := struct {
		Technologies []Technology
		Strategies   []Strategy
	}{
		Technologies: technologies,
		Strategies:   strategies,
	}

	templates.ExecuteTemplate(w, "technology.html", data)
}

func technologiesAPI(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	techRows, err := db.Query("SELECT id, name, description, image FROM technologies")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des technologies", http.StatusInternalServerError)
		return
	}
	defer techRows.Close()

	var technologies []Technology
	for techRows.Next() {
		var tech Technology
		err = techRows.Scan(&tech.ID, &tech.Name, &tech.Description, &tech.Image)
		if err != nil {
			http.Error(w, "Erreur lors du traitement des technologies", http.StatusInternalServerError)
			return
		}
		technologies = append(technologies, tech)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(technologies)
}

// Gestionnaire pour les stratégies
func strategiesAPI(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	stratRows, err := db.Query("SELECT id, name, description FROM strategies")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des stratégies", http.StatusInternalServerError)
		return
	}
	defer stratRows.Close()

	var strategies []Strategy
	for stratRows.Next() {
		var strat Strategy
		err = stratRows.Scan(&strat.ID, &strat.Name, &strat.Description)
		if err != nil {
			http.Error(w, "Erreur lors du traitement des stratégies", http.StatusInternalServerError)
			return
		}
		strategies = append(strategies, strat)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(strategies)
}
