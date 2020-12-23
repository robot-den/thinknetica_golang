package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/crawler"
	"pkg/crawler/webscnr"
	"pkg/engine"
	"pkg/index"
	"pkg/index/hash"
	"pkg/model"
	"pkg/storage"
	"pkg/storage/memory"
	"strings"
	"sync"
)

// Service представляет собой сервер интернет-поисковика
type Service struct {
	storage storage.Service
	index   index.Service
	engine  *engine.Service
	scanner crawler.Scanner
}

func main() {
	service := new()

	go service.scan()
	service.readline()
}

func new() *Service {
	str := memory.NewStorage()
	ind := hash.NewService()
	eng := engine.NewService(ind, str)

	s := Service{
		storage: str,
		index:   ind,
		engine:  eng,
		scanner: &webscnr.WebScnr{},
	}

	return &s
}

type scanJob struct {
	URL   string
	Depth int
}

var scanJobs = []scanJob{
	{"https://habr.com", 1},
	{"https://dev.to", 2},
	{"https://www.postgresql.org/", 2},
	{"https://clickhouse.tech/", 2},
	{"https://redis.io/", 2},
	{"https://memcached.org/", 2},
	{"https://kafka.apache.org/", 2},
	{"https://grafana.com/", 2},
	{"https://www.zabbix.com/", 2},
	{"https://developer.mozilla.org/en-US/", 2},
	{"https://stackoverflow.com/", 2},
	{"https://www.englishgrammar.org/", 2},
}

const scanersPoolSize = 10

// scan сканирует сайты, сохраняет результаты в хранилище и обновляет индекс.
// Для сканирования используется пул воркеров
func (s *Service) scan() {
	toScan := make(chan scanJob)
	rawDocs := make(chan []model.Document)
	wg := &sync.WaitGroup{}

	// Starts workers pool
	for i := 0; i < scanersPoolSize; i++ {
		wg.Add(1)
		go s.scanWorker(wg, toScan, rawDocs)
	}
	// Sends jobs to workers and close toScan channel when jobs are over
	go func(ch chan<- scanJob) {
		for _, job := range scanJobs {
			ch <- job
		}
		close(ch)
	}(toScan)
	// Reads and saves results from workers. This goroutine done when scan() func closes channel it iterates
	go func(ch <-chan []model.Document) {
		for docs := range ch {
			docsWithIds := s.storage.Write(docs)
			s.index.Update(docsWithIds)
		}
	}(rawDocs)

	wg.Wait()
	close(rawDocs)
}

func (s *Service) scanWorker(wg *sync.WaitGroup, jobs <-chan scanJob, results chan<- []model.Document) {
	defer wg.Done()

	for job := range jobs {
		docs, err := s.scanner.Scan(job.URL, job.Depth)
		if err != nil {
			fmt.Println(err)
			continue // NOTE: do not stop worker if scan failed, let it take next job
		}
		results <- docs
	}
}

func (s *Service) readline() {
	for {
		fmt.Println("Enter search word (leave empty to exit):")
		reader := bufio.NewReader(os.Stdin)
		word, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		word = strings.TrimSuffix(word, "\r\n")
		word = strings.TrimSuffix(word, "\n")
		if word == "" {
			break
		}

		found := s.engine.Search(word)
		fmt.Printf("Results for '%s':\n", word)
		for _, v := range found {
			fmt.Println(v)
		}
	}

	fmt.Println("Bye!")
}
