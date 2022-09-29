package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/dctalbot/air/awair"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	if os.Getenv("AWAIR_TOKEN") == "" {
		log.Fatal("AWAIR_TOKEN environment variable not set")
		return
	}
	if os.Getenv("DB_CONN_STR") == "" {
		log.Fatal("DB_CONN_STR environment variable not set")
		return
	}

	devices := awair.GetDevices()
	if len(devices) < 1 {
		log.Fatal("No devices found")
		return
	}

	data := awair.GetLatest(devices[0])

	db, err := sql.Open("pgx", os.Getenv("DB_CONN_STR"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	_, err = db.Exec(`
	INSERT INTO latest (at, score, voc, co2, temp, pm25, humidity, voc_index, co2_index, temp_index, pm25_index, humidity_index) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		data.Timestamp,
		data.Score,

		data.Sensor("voc"),
		data.Sensor("co2"),
		data.Sensor("temp"),
		data.Sensor("pm25"),
		data.Sensor("humid"),

		data.Index("voc"),
		data.Index("co2"),
		data.Index("temp"),
		data.Index("pm25"),
		data.Index("humid"),
	)

	if err != nil {
		log.Fatal(err)
		return
	}
}
