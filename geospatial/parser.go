package geospatial

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
)

func GeoPostcode(search string) (string, string, error) {
	// Open the CSV file
	file, err := os.Open("data/uk_postcodes.csv")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read in the CSV records line by line
	for {
		// Read one record (line) from CSV
		record, err := reader.Read()
		if err != nil {
			// Check for end-of-file or other errors
			if err.Error() == "EOF" {
				break // End of file reached
			}
			log.Fatal("Error reading CSV:", err)
		}

		if record[1] == search {
			//fmt.Println(record)
			return record[2], record[3], nil
		}
		//fmt.Println() // Move to the next line for the next record
	}

	return "", "", errors.New("nothing found")
}

func GeoDistrict(search string) (string, string, error) {
	// Open the CSV file
	file, err := os.Open("data/uk_districts.csv")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read in all the CSV records
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV:", err)
	}

	// Iterate over each record and print the values
	for _, record := range records {
		if record[1] == search {
			//fmt.Println(record)
			return record[2], record[3], nil
		}
		//fmt.Println() // Move to the next line for the next record
	}

	return "", "", errors.New("nothing found")
}
