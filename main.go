package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <url>")
		return
	}
	url := os.Args[1]

	recipe, err := ScrapeRecipe(url)
	if err != nil {
		fmt.Println("Error scrapping recipe:", err)
		return
	}

	fmt.Printf("Scrapped recipe: \n%+v\n", recipe)
}
