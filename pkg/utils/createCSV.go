package repeatable

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteInCSV(r []string) {

	f, err := os.Create("report.csv")
	defer f.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range r {
		row := []string{record}
		w.Write(row)
	}

}
