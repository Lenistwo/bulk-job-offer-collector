package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	BaseUrl               = "https://www.pracuj.pl"
	JobOfferURL           = BaseUrl + "/praca/java%20developer;kw/warszawa;wp?rd=30"
	OfferSection          = "window.__INITIAL_STATE__ ="
	OpeningBracket        = "{"
	ClosingBracket        = "}"
	NewLine               = "\n"
	NoBrakingSpace        = " "
	Space                 = " "
	EmptyString           = ""
	ReplaceAllAppearances = -1
	OutputFileName        = "results.json"
	ReadPermission        = 0644
)

var (
	foundJobs []JobOffer
	outputs   []Output
)

func main() {
	startTime := time.Now()
	startPage := 1
	stillActive := make(chan bool)
	go bulkLoadJobOffers(startPage, startPage+1, stillActive)
	if !<-stillActive {
		prepareOutput()
		writeOutputToResultFile()
	}
	fmt.Printf("Job Offer Load Duration %v", time.Since(startTime))
}

func prepareOutput() {
	for _, job := range foundJobs {
		applicationURL := ""
		if len(job.Offers) > 0 {
			applicationURL = job.Offers[0].OfferUrl
		}

		outputs = append(outputs, Output{
			Title:           job.JobTitle,
			EmploymentLevel: job.EmploymentLevel,
			Employer:        job.Employer,
			Description:     job.JobDescription,
			Salary:          job.Salary,
			TypesOfContract: job.TypesOfContract,
			ApplicationURL:  applicationURL,
			ExpiresAt:       job.ExpirationDate,
		})
	}
}

func writeOutputToResultFile() {
	indent, err := json.MarshalIndent(outputs, EmptyString, Space)
	checkError(err)
	err = ioutil.WriteFile(OutputFileName, indent, ReadPermission)
	checkError(err)
}

func bulkLoadJobOffers(currentPage, lastPage int, stillActive chan bool) {
	if currentPage <= 0 || currentPage > lastPage {
		stillActive <- false
		return
	}
	parsedBody := sendRequest(JobOfferURL + fmt.Sprint("&pn=", currentPage))
	offers := getJobOffersSection(parsedBody)
	prettyOffers := extractJSON(offers)
	jsonMap := unmarshallToMap(prettyOffers)
	jobOffers := unmarshallJobOffers(jsonMap)
	pagination := unmarshallPagination(jsonMap)
	foundJobs = append(foundJobs, jobOffers...)
	go bulkLoadJobOffers(currentPage+1, pagination.MaxPages, stillActive)
}

func unmarshallJobOffers(jsonMap map[string]json.RawMessage) []JobOffer {
	var jobOffers []JobOffer
	err := json.Unmarshal(jsonMap["offers"], &jobOffers)
	checkError(err)
	return jobOffers
}

func unmarshallPagination(jsonMap map[string]json.RawMessage) Pagination {
	var pagination Pagination
	err := json.Unmarshal(jsonMap["pagination"], &pagination)
	checkError(err)
	return pagination
}

func unmarshallToMap(extractedOffers string) map[string]json.RawMessage {
	objectMap := make(map[string]json.RawMessage)
	err := json.Unmarshal([]byte(extractedOffers), &objectMap)
	checkError(err)
	return objectMap
}

func getJobOffersSection(parsedBody string) string {
	offerIndex := strings.Index(parsedBody, OfferSection)
	offer := parsedBody[offerIndex:]
	newLineIndex := strings.Index(offer, NewLine)
	return offer[:newLineIndex]
}

func sendRequest(url string) string {
	getResponse, err := http.Get(url)
	checkError(err)
	responseBody := getResponse.Body
	defer responseBody.Close()
	rawBody, err := ioutil.ReadAll(responseBody)
	checkError(err)
	return strings.Replace(string(rawBody), NoBrakingSpace, Space, ReplaceAllAppearances)
}

func extractJSON(json string) string {
	firstOpeningBracketIndex := strings.Index(json, OpeningBracket)
	lastClosingBracketIndex := strings.LastIndex(json, ClosingBracket) + 1
	return json[firstOpeningBracketIndex:lastClosingBracketIndex]
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
