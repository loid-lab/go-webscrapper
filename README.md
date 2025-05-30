# Go Web Scraper

A simple web scraper written in Go using the [Colly](https://github.com/gocolly/colly) framework.  
This scraper extracts recipe information (title, ingredients, prep time, cooking time, servings, etc.) from recipe web pages like [Love and Lemons](https://www.loveandlemons.com).

---

## Features

- Scrapes recipe name, ingredients, and basic specifications  
- Uses CSS selectors to extract data  
- Handles HTTP requests and errors gracefully  
- Easily customizable to scrape other recipe sites by updating selectors

---

## Prerequisites

- Go 1.18+ installed on your machine  
- Internet connection for scraping live websites

---

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/loid-lab/go-webscrapper.git
   cd go-webscrapperr
   ```

2. Download dependencies:

   ```bash
   go mod tidy
   ```

---

## Usage

Run the scraper by providing the recipe URL as a command-line argument:

```bash
go run main.go scraper.go <recipe-url>
```

Example:

```bash
go run main.go scraper.go https://www.loveandlemons.com/homemade-pizza/
```

The program will visit the URL, scrape the recipe details, and print the results in the console.

---

## Code Structure

- `main.go`: Entry point; parses command-line arguments and calls the scraper function  
- `scraper.go`: Contains the scraping logic using Colly and data models for recipes

---

## Customization

To scrape other websites or extract additional data:

- Inspect the target site's HTML structure  
- Update CSS selectors inside `scraper.go` accordingly (in the `ScrapeRecipe` function)  

---

## License

MIT License © 2025 Fabio Gualtieri

---

## Contact

For questions or improvements, feel free to open an issue or contact me.
