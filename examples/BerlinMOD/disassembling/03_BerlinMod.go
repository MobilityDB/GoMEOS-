package main

//MEOS example: meos/examples/05_berlinmod_disassemble.c

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/MobilityDB/GoMEOS/gomeos"
)

const (
	MaxLengthTrip   = 170001
	MaxLengthHeader = 1024
	MaxLengthDate   = 12
	MaxNoTrips      = 64
)

type TripRecord struct {
	TripID int
	VehID  int
	Day    time.Time
	Seq    int
	Trip   *gomeos.TGeomPointSeq
}

func main() {
	// Arrays to compute the results
	var trips [MaxNoTrips]TripRecord
	var currInst [MaxNoTrips]int

	// Get start time
	startTime := time.Now()

	// Initialize MEOS
	gomeos.MeosInitialize("UTC")

	// Open the input CSV file
	file, err := os.Open("data/berlinmod_trips.csv")
	if err != nil {
		log.Fatalf("Error opening input file: %v\n", err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.FieldsPerRecord = -1 // To handle varying number of fields

	// Read the first line (headers)
	_, err = reader.Read()
	if err != nil {
		log.Fatalf("Error reading header: %v\n", err)
	}

	i := 0
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("Error reading record: %v\n", err)
		}

		// Extract values
		tripID, _ := strconv.Atoi(record[0])
		vehID, _ := strconv.Atoi(record[1])
		dateBuffer := record[2]
		seq, _ := strconv.Atoi(record[3])
		tripBuffer := record[4]

		// Transform the date string into a time.Time value
		day, err := time.Parse("2006-01-02", dateBuffer) // Example format: "YYYY-MM-DD"
		if err != nil {
			log.Fatalf("Error parsing date: %v\n", err)
		}

		// Transform the trip string into a Temporal value
		trip := gomeos.NewTGeomPointSeqFromWKB(tripBuffer)

		// Save the trip record
		trips[i] = TripRecord{
			TripID: tripID,
			VehID:  vehID,
			Day:    day,
			Seq:    seq,
			Trip:   trip,
		}

		i++
	}

	recordsIn := i

	fmt.Println("finish reading csv")
	// Open the output CSV file
	outputFile, err := os.Create("data/berlinmod_instants.csv")
	if err != nil {
		log.Fatalf("Error creating output file: %v\n", err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Write the header line
	fmt.Println("start writing csv")
	writer.Write([]string{"tripid", "vehid", "day", "seqno", "geom", "t"})

	// Initialize the current instant for each trip to the first one
	for i := 0; i < MaxNoTrips; i++ {
		currInst[i] = 1
	}

	// Loop until all trips have been processed
	recordsOut := 0
	for {
		// Find the minimum instant
		first := 0
		for first < recordsIn && currInst[first] < 0 {
			first++
		}
		if first == recordsIn {
			// All trips have been processed
			break
		}

		minInst := gomeos.TemporalInstantN(trips[first].Trip, &gomeos.TGeomPointInst{}, currInst[first])
		minTrip := first
		// Loop for the minimum instant among all remaining trips
		for i := first + 1; i < recordsIn; i++ {
			if currInst[i] < 0 {
				continue
			}
			inst := gomeos.TemporalInstantN(trips[i].Trip, &gomeos.TGeomPointInst{}, currInst[i])
			if minInst.Timestamptz().After(inst.Timestamptz()) {
				minInst = inst
				minTrip = i
			}
		}

		// Write line in the CSV file
		dateStr := trips[minTrip].Day.Format("2006-01-02")
		geomStr := gomeos.GeoAsEWKT(minInst, 6)
		timeStr := minInst.TimestampOut()
		writer.Write([]string{
			strconv.Itoa(trips[minTrip].VehID),
			strconv.Itoa(trips[minTrip].VehID),
			dateStr,
			strconv.Itoa(trips[minTrip].Seq),
			geomStr,
			timeStr,
		})

		recordsOut++

		// Advance the current instant of the trip
		currInst[minTrip]++
		if currInst[minTrip] > gomeos.TemporalNumInstants(trips[minTrip].Trip) {
			currInst[minTrip] = -1
		}
	}

	fmt.Printf("%d trip records read from file 'berlimod_trips.csv'.\n", recordsIn)
	fmt.Printf("%d observation records written in file 'berlimod_instants.csv'.\n", recordsOut)

	// Calculate the elapsed time
	elapsedTime := time.Since(startTime)
	fmt.Printf("The program took %f seconds to execute\n", elapsedTime.Seconds())

	// Free memory (handled automatically in Go)
	// Finalize MEOS
	gomeos.MeosFinalize()
}