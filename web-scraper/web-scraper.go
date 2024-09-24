package webscraper

import (
	// Set up Rod for Chrome browser automation
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	m "github.com/njh18/tcg-tracker-discord-bot/model"
)

func Main() {
	launcherPath := launcher.New().Headless(true).NoSandbox(true).MustLaunch()
	browser := rod.New().ControlURL(launcherPath).MustConnect()
	defer browser.MustClose()

	version := "eb01"
	url := "https://yuyu-tei.jp/sell/opc/s/" + version

	page := browser.MustPage(url)

	// Wait for the elements to be loaded
	time.Sleep(5 * time.Second)

	// Find all card list blocks
	cardListBlocks := page.MustElements("div.cards-list")

	cardInfoMap := make(map[string][]m.CardInfo)

	for _, block := range cardListBlocks {
		title := block.MustElement("h3").MustText()
		fmt.Println(string(title))

		if !strings.Contains(title, "Card List") {
			continue
		}

		// Extract all card-product elements from the SR block
		cardProducts := block.MustElements("div.card-product")

		var cardInfoList []m.CardInfo
		currentTimestamp := time.Now()

		for _, cardProduct := range cardProducts {
			// Extract the card image URL
			cardInfo := buildCardInfoFromHtml(cardProduct, currentTimestamp)
			cardInfoList = append(cardInfoList, cardInfo)
		}
		cardInfoMap[title] = cardInfoList
	}

	infoMap, err := json.Marshal(cardInfoMap)
	if err != nil {
		log.Fatalf("Error converting to JSON: %s", err)
	}

	// Print the JSON as a string
	fmt.Println(string(infoMap))

	fmt.Println("Job Done.")

}

func buildCardInfoFromHtml(cardProduct *rod.Element, currentTimeStamp time.Time) m.CardInfo {
	imgElement := cardProduct.MustElement("div.product-img img")
	imageURL := imgElement.MustAttribute("src")

	code := cardProduct.MustElement("span").MustText()
	cardName := cardProduct.MustElement("h4").MustText()

	hrefLink := cardProduct.MustElement("a").MustAttribute("href")
	yenPrice := cardProduct.MustElement("strong").MustText()

	return m.CardInfo{
		ImageURL:    *imageURL,
		Code:        code,
		CardName:    cardName,
		HrefLink:    *hrefLink,
		YenPrice:    yenPrice,
		UpdatedTime: currentTimeStamp,
	}
}
