package sparser

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sber-scrape/internal/model"
	"sber-scrape/internal/store"
	"strconv"
	"sync"
)

func GetJson(store store.Store, searchText string, page int) error {
	// Request body
	log.Printf("Searching %s on page %v", searchText, page)
	requestBody := []byte(fmt.Sprintf(`{"requestVersion":10,"limit":44,"offset":%s,"collectionId":"","selectedAssumedCollectionId":"","isMultiCategorySearch":false,"searchByOriginalQuery":false,"selectedSuggestParams":[],"expandedFiltersIds":[],"sorting":0,"ageMore18":null,"showNotAvailable":true,"searchText":"%s","auth":{"locationId":"50","appPlatform":"WEB","appVersion":1707385735,"experiments":{"8":"1","55":"2","58":"2","68":"2","69":"1","79":"3","98":"1","99":"1","107":"2","109":"2","119":"2","120":"2","121":"2","122":"2","128":"1","132":"1","144":"3","154":"1","173":"1","184":"3","186":"2","190":"1","192":"2","194":"3","200":"2","205":"2","209":"1","218":"1","235":"2","237":"2","243":"1","5779":"2","20121":"2","43568":"2","70070":"2","85160":"2"},"os":"UNKNOWN_OS"}}`, strconv.Itoa(page*44+2), searchText))
	req, err := http.NewRequest("POST", "https://megamarket.ru/api/mobile/v1/catalogService/catalog/search", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("Error creating request:", err)
		return err
	}

	// Request headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate")
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	response := &model.StripResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return err
	}

	err = store.Product().JSONTransaction(*response)
	if err != nil {
		log.Print("Transaction error: ", err)
		return err
	}
	return nil
}

func GetPagesJson(searchText string, store store.Store, pages int) error {
	u, err := url.Parse(searchText)
	if err != nil {
		log.Println("Error parsing URL:", err)
		return err
	}

	// Get the value of the "q" query parameter
	queryValue := u.Query().Get("q")

	// Creating a wait group for concurent page parsing
	var wg sync.WaitGroup
	wg.Add(pages)

	// Context for goroutine termination
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Goroutine for waitgroup (JSON parsing)
	parsePage := func(pendingPage int) {
		defer wg.Done()

		// Check for errors in other goroutines
		select {
		case <-ctx.Done():
			log.Print("Error occured")
			return
		default:
		}

		err := GetJson(store, queryValue, pendingPage)
		if err != nil {
			log.Printf("Page %d failed", pendingPage)
			cancel()
			return
		}

		log.Printf("Page %d parsed", pendingPage)
	}

	for page := 1; page <= pages; page++ {
		go parsePage(page)
	}

	wg.Wait()
	log.Printf("Finished, parsed %d pages", pages)
	return nil
}
