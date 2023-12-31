package sparser

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"sber-scrape/internal/model"
	"sber-scrape/internal/store"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func parseItemBlock(itemBlock *goquery.Selection, store store.Store) {
	// Helper function to extract text from an element and trim whitespace
	getText := func(selector string) string {
		return strings.TrimSpace(
			strings.ReplaceAll(
				itemBlock.Find(selector).Text(),
				"\t",
				"",
			),
		)
	}
	convInt := func(str string) (int, error) {
		if str != "" {
			regex := regexp.MustCompile("[^0-9]+")
			result := regex.ReplaceAllString(str, "")
			return strconv.Atoi(result)
		}
		return 0, nil
	}

	itemTitle := getText(".item-title")
	itemPrice, _ := convInt(getText(".item-price"))
	bonusAmount, _ := convInt(getText(".bonus-amount"))
	bonusPercent, _ := convInt(getText(".bonus-percent"))
	discount, _ := convInt(getText(".discount-percentage__value"))

	// Extract productID and link attributes
	productIDText, _ := itemBlock.Find(".ddl_product_link").Attr("data-product-id")
	productID, _ := convInt(productIDText)
	val, _ := itemBlock.Find(".ddl_product_link").Attr("href")
	link := fmt.Sprintf("%s%s", "https://megamarket.ru", val)

	p := &model.Product{
		Title:        itemTitle,
		Price:        itemPrice,
		BonusAmount:  bonusAmount,
		BonusPercent: bonusPercent,
		Discount:     discount,
		ProductID:    productID,
		Link:         link,
	}
	if err := store.Product().Create(p); err != nil {
		log.Fatal(err)
		return
	}
	// Print the extracted data
	fmt.Println("Title: ", itemTitle)
	fmt.Println("Price: ", itemPrice)
	fmt.Println("SBonuses: ", bonusAmount)
	fmt.Println("SBonuses %: ", bonusPercent)
	fmt.Println("Discount: ", discount)
	fmt.Println("Product ID", productID)
	fmt.Println("URL: ", link)
	// fmt.Println("-" * 10)
}

func bufferItemBlock(itemBlock *goquery.Selection) *model.Product {
	// Helper function to extract text from an element and trim whitespace
	getText := func(selector string) string {
		return strings.TrimSpace(
			strings.ReplaceAll(
				itemBlock.Find(selector).Text(),
				"\t",
				"",
			),
		)
	}
	convInt := func(str string) (int, error) {
		if str != "" {
			regex := regexp.MustCompile("[^0-9]+")
			result := regex.ReplaceAllString(str, "")
			return strconv.Atoi(result)
		}
		return 0, nil
	}

	itemTitle := getText(".item-title")
	itemPrice, _ := convInt(getText(".item-price"))
	bonusAmount, _ := convInt(getText(".bonus-amount"))
	bonusPercent, _ := convInt(getText(".bonus-percent"))
	discount, _ := convInt(getText(".discount-percentage__value"))

	// Extract productID and link attributes
	productIDText, _ := itemBlock.Find(".ddl_product_link").Attr("data-product-id")
	productID, _ := convInt(productIDText)
	val, _ := itemBlock.Find(".ddl_product_link").Attr("href")
	link := fmt.Sprintf("%s%s", "https://megamarket.ru", val)

	p := &model.Product{
		Title:        itemTitle,
		Price:        itemPrice,
		BonusAmount:  bonusAmount,
		BonusPercent: bonusPercent,
		Discount:     discount,
		ProductID:    productID,
		Link:         link,
	}
	return p
}

func GetHtml(url string, store store.Store) error {
	var productArray []model.Product
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	r, err := http.NewRequest("GET", url, nil)
	log.Println(url)
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_1) AppleWebKit/534.16 (KHTML, like Gecko) Chrome/52.0.1458.377 Safari/603")
	r.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".item-block").Each(func(index int, itemBlock *goquery.Selection) {
		// parseItemBlock(itemBlock, store)
		productArray = append(productArray, *bufferItemBlock(itemBlock))
	})
	// err = store.Product().BulkInsertProducts(productArray)
	err = store.Product().Transaction(productArray)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func GetLocalHtml(filename string, store store.Store) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".item-block").Each(func(index int, itemBlock *goquery.Selection) {
		parseItemBlock(itemBlock, store)
	})
	return nil
}

func GetPages(url string, store store.Store, pages int) error {
	for page := 1; page <= pages; page++ {
		currurl := strings.Replace(
			url,
			"/catalog",
			fmt.Sprintf("/catalog/page-%d", page),
			1,
		)
		err := GetHtml(currurl, store)
		if err != nil {
			return err
		}
	}
	return nil
}
