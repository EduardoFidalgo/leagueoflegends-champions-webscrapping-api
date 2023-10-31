package main

import (
	"fmt"
	"goWebScrapping/models"
	"goWebScrapping/routes"
	"strings"

	"github.com/gocolly/colly"
)

const (
	version float64 = 1.1
)

func Starting() {
	fmt.Println("Starting server...")
	fmt.Println("Server version: ", version)

}

func main() {
	Starting()
	c := colly.NewCollector()

	var characterList []models.Character
	var currentCharacter models.Character
	var skills []string

	// Getting data from HTML
	c.OnHTML("a.style__Wrapper-sc-n3ovyt-0", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		e.Request.Visit(link)
	})

	// Configurando as novas funções OnHTML para a página de destino
	c.OnHTML("span[data-testid='overview:title']", func(e *colly.HTMLElement) {
		characterName := e.Text
		currentCharacter.Name = characterName
	})

	c.OnHTML("div[data-testid='overview:backgroundimage']>img.style__NoScriptImg-sc-g183su-0", func(e *colly.HTMLElement) {
		characterImgUrl := e.Attr("src")
		currentCharacter.Url = characterImgUrl
	})

	c.OnHTML("p[data-testid='overview:description']", func(e *colly.HTMLElement) {
		characterDescription := e.Text
		currentCharacter.History = characterDescription
	})

	c.OnHTML("div[data-testid='overview:difficulty']", func(e *colly.HTMLElement) {
		characterDifficulty := e.Text
		currentCharacter.Difficulty = characterDifficulty
	})

	c.OnHTML("label.style__CarouselItemText-sc-gky2mu-16", func(e *colly.HTMLElement) {
		skinName := strings.TrimSpace(e.Text)
		currentCharacter.Skins = append(currentCharacter.Skins, skinName)
	})

	c.OnHTML("li.style__AbilityInfoItem-sc-1bu2ash-8", func(e *colly.HTMLElement) {
		skillKey := e.ChildText("h6")
		skillName := e.ChildText("h5")

		skill := fmt.Sprintf("%s: %s", skillKey, skillName)
		skills = append(skills, skill)

		currentCharacter.Skills = skills
	})

	c.OnHTML("div[data-testid='overview:role']", func(e *colly.HTMLElement) {
		characterRole := e.Text
		currentCharacter.Role = characterRole

		characterList = append(characterList, currentCharacter)
		currentCharacter = models.Character{}
		skills = nil
	})

	// Exec request to visit my site
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting.. ", r.URL)
	})

	// Site to be visited..
	c.Visit("https://www.leagueoflegends.com/en-us/champions/")

	routes.HandleRequest(characterList)
}
