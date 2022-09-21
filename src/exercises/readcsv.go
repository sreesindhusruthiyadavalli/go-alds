package exercises

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCsv(filePath string)[][]string{
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file " + filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for " + filePath, err)
	}
	return records
}

func GetRecords(){
	fmt.Println(os.Getwd())
	records := readCsv("./problems.csv")
	print(records)
}


