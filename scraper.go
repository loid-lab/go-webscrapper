package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Dictionary map[string]string

type RecipeSpecs struct {
	difficulty, prepTime, cookingTime, servingSize, priceTier string
}

type Recipe struct {
	url, name      string
	ingredients    []Dictionary
	specifications RecipeSpecs
}

var recipes []Recipe

func ScrapeRecipe(url string) (Recipe, error) {
	collector := colly.NewCollector()
	var scrapedRecipe Recipe
	scrapedRecipe.url = url

	collector.OnHTML("h1.entry-title", func(e *colly.HTMLElement) {
		scrapedRecipe.name = e.Text
		fmt.Println("Scraped name:", e.Text)
	})

	collector.OnHTML("li.wprm-recipe-ingredient", func(h *colly.HTMLElement) {
		ingredient := Dictionary{
			"name": h.Text,
		}
		scrapedRecipe.ingredients = append(scrapedRecipe.ingredients, ingredient)
		fmt.Println("Scraped ingredient:", h.Text)
	})

	collector.OnHTML("span.wprm-recipe-prep_time", func(h *colly.HTMLElement) {
		scrapedRecipe.specifications.prepTime = h.Text
		fmt.Println("Scraped prep time:", h.Text)
	})

	collector.OnHTML("span.wprm-recipe-cook_time", func(h *colly.HTMLElement) {
		scrapedRecipe.specifications.cookingTime = h.Text
		fmt.Println("Scraped cooking time:", h.Text)
	})

	collector.OnHTML("span.wprm-recipe-servings", func(h *colly.HTMLElement) {
		scrapedRecipe.specifications.servingSize = h.Text
		fmt.Println("Scraped servings:", h.Text)
	})

	collector.OnHTML("span.difficulty", func(h *colly.HTMLElement) {
		scrapedRecipe.specifications.difficulty = h.Text
		fmt.Println("Scraped difficulty:", h.Text)
	})

	collector.OnHTML("span.price-tier", func(h *colly.HTMLElement) {
		scrapedRecipe.specifications.priceTier = h.Text
		fmt.Println("Scraped price tier:", h.Text)
	})

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error", err)
	})

	err := collector.Visit(url)
	if err != nil {
		return Recipe{}, err
	}

	recipes = append(recipes, scrapedRecipe)

	fmt.Printf("Scraped recipe %+v\n", scrapedRecipe)

	return scrapedRecipe, nil
}
