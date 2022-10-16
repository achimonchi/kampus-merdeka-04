package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sesi6-gin/config"
	"strings"
	"sync"
	"time"
)

const worker = 10

var dataHeades = make([]string, 0)

func main() {

	start := time.Now()
	db, err := config.ConnectPostgres()
	if err != nil {
		panic(err)
	}

	jobs := make(chan []interface{})
	wg := sync.WaitGroup{}

	reader, file, err := openCsvFile("./cmd/data.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	go dispatchWorkers(db, jobs, &wg)

	readCsvFilePerLineThenSendToWorker(reader, jobs, &wg)

	wg.Wait()

	end := time.Since(start).Milliseconds()
	fmt.Println("time execution:", end)
}

func openCsvFile(csvFile string) (*csv.Reader, *os.File, error) {
	log.Println("open csv file")

	f, err := os.Open(csvFile)
	if err != nil {
		return nil, nil, err
	}
	read := csv.NewReader(f)
	return read, f, nil
}

func readCsvFilePerLineThenSendToWorker(reader *csv.Reader, jobs chan<- []interface{}, wg *sync.WaitGroup) {
	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if len(dataHeades) == 0 {
			dataHeades = row
			continue
		}

		rowOrdered := make([]interface{}, 0)
		for _, each := range row {
			rowOrdered = append(rowOrdered, each)
		}

		wg.Add(1)
		jobs <- rowOrdered
	}
	close(jobs)
}

func dispatchWorkers(db *sql.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
	for i := 0; i < worker; i++ {
		go func(index int, db *sql.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
			counter := 0
			for job := range jobs {
				doTheJob(i, counter, db, job)
				wg.Done()
				counter++
			}
		}(i, db, jobs, wg)
	}
}

func doTheJob(index, counter int, db *sql.DB, val []interface{}) {
	if counter%10 == 0 {
		log.Println("> worker", index, "inserted", counter, "data")
	}
	for {
		time.Sleep(50 * time.Millisecond)
		var outerError error
		func(outerError *error) {
			headers := strings.Join(dataHeades, ",")
			headers = headers[:len(headers)-1]

			question := strings.Join(generateQuestionMark(len(dataHeades)), ",")
			question = question[:len(question)]

			query := fmt.Sprintf(`INSERT INTO domain (%s) 
			VALUES (%s)`,
				headers,
				question,
			)
			stmt, err := db.Prepare(query)

			if err != nil {
				panic(err)
			}
			defer stmt.Close()

			_, err = stmt.Exec(val...)

			if err != nil {
				panic(err)
			}
		}(&outerError)
		if outerError != nil {
			break
		}
	}

}

func generateQuestionMark(n int) []string {
	s := []string{}
	for i := 0; i < n; i++ {
		s = append(s, fmt.Sprintf("$%d", i+1))
	}
	return s
}
