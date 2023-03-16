package phase2

import (
	"CA1/phase1"
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type Job struct {
	id   int
	line string
}

type Result struct {
	job           Job
	formattedLine string
}

func worker(wg *sync.WaitGroup, jobs chan Job, results chan Result) {
	for job := range jobs {
		output := Result{job, phase1.FormatText(job.line)}
		results <- output
	}
	wg.Done()
}

func allocate(inputFilePath string, jobs chan Job) {
	inputFile, _ := os.Open(inputFilePath)
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	i := 0

	for scanner.Scan() {
		text := scanner.Text()
		job := Job{i, text}
		jobs <- job
		i++
	}

	close(jobs)
}

func createWorkerPool(noOfWorkers int, jobs chan Job, results chan Result) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg, jobs, results)
	}

	wg.Wait()
	close(results)
}

func result(results chan Result, done chan bool, outputFilePath string) {
	outputFile, _ := os.Create(outputFilePath)
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)
	sortedResult := make(map[int]string)

	for res := range results {
		sortedResult[res.job.id] = res.formattedLine
	}

	for i := 0; i < len(sortedResult); i++ {
		writer.WriteString(sortedResult[i] + "\n")
	}

	writer.Flush()
	done <- true
}

func Run(inputFilePath string, outputFilePath string, noOfWorkers int) {
	startTime := time.Now()
	jobs := make(chan Job, noOfWorkers)
	results := make(chan Result, noOfWorkers)
	go allocate(inputFilePath, jobs)
	done := make(chan bool)
	go result(results, done, outputFilePath)
	createWorkerPool(noOfWorkers, jobs, results)
	<-done
	elapsed := time.Since(startTime)
	fmt.Println("total time taken ", elapsed.Seconds(), "seconds")
}
