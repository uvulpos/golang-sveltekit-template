package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

var appversion string = "dev"
var appbuildcommit string = "_"
var appbuilddate string

//go:embed frontend/*
var svelteFS embed.FS

//go:embed pokedata/*
var pokeData embed.FS

type PokemonJSON struct {
	Id               int    `json="id"`
	Number           string `json="number"`
	Slug             string `json="slug"`
	CollectiblesSlug string `json="collectibles_slug"`
	Name             string `json="name"`

	Height float64  `json="height"`
	Weight float64  `json="weight"`
	Type   []string `json="type"`

	Abilities []string `json="abilities"`
	Weakness  []string `json="weakness"`

	ThumbnailImage   string `json="ThumbnailImage"`
	ThumbnailAltText string `json="ThumbnailAltText"`
	DetailPageURL    string `json="detailPageURL"`
	Featured         string `json="featured"`
}

type NewPokemonStruct struct {
	Number           string `json="number"`
	Name             string `json="name"`
	ThumbnailImage   string `json="image"`
	ThumbnailAltText string `json="imageAlt"`
	DetailPageURL    string `json="detailPageURL"`
}

func main() {

	fmt.Println("Version ", appversion)
	fmt.Println("Build-Commit ", appbuildcommit)
	fmt.Println("Build-Date ", appbuilddate)

	app := fiber.New(fiber.Config{
		DisableStartupMessage: appversion == "",
	})

	publicFS, err := fs.Sub(svelteFS, "frontend")
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/api/:lang<regex((en|de))>?", func(c *fiber.Ctx) error {

		pokemons, pokemonErr := getPokemon(c.Params("lang", "en"))
		if pokemonErr != nil {
			c.SendString("internal server error")
			c.SendStatus(http.StatusInternalServerError)
			return pokemonErr
		}

		var newPokemon []NewPokemonStruct
		var existingIDs []int
		for _, pokemon := range pokemons {
			if !existsID(pokemon.Id, existingIDs) {
				existingIDs = append(existingIDs, pokemon.Id)
				newPokemon = append(newPokemon, NewPokemonStruct{
					Number:           pokemon.Number,
					Name:             pokemon.Name,
					ThumbnailImage:   pokemon.ThumbnailImage,
					ThumbnailAltText: pokemon.ThumbnailAltText,
					DetailPageURL:    pokemon.DetailPageURL,
				})
			}
		}

		return c.JSON(newPokemon)
	})
	app.Get("/api/:lang<regex((en|de))>/:name<max(50)>", func(c *fiber.Ctx) error {

		pokemons, pokemonErr := getPokemon(c.Params("lang", "en"))
		if pokemonErr != nil {
			c.SendString("internal server error")
			c.SendStatus(http.StatusInternalServerError)
			return pokemonErr
		}

		var newPokemon = []NewPokemonStruct{}
		var existingIDs []int
		for _, pokemon := range pokemons {

			if !strings.Contains(pokemon.Name, c.Params("name", "")) {
				continue
			}

			if !existsID(pokemon.Id, existingIDs) {
				existingIDs = append(existingIDs, pokemon.Id)
				newPokemon = append(newPokemon, NewPokemonStruct{
					Number:           pokemon.Number,
					Name:             pokemon.Name,
					ThumbnailImage:   pokemon.ThumbnailImage,
					ThumbnailAltText: pokemon.ThumbnailAltText,
					DetailPageURL:    pokemon.DetailPageURL,
				})
			}
		}

		return c.JSON(newPokemon)
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(publicFS),
		NotFoundFile: "index.html",
	}))

	app.Listen(":3000")
}

func existsID(id int, ids []int) bool {
	for _, searchingID := range ids {
		if searchingID == id {
			return true
		}
	}
	return false
}

func getPokemon(lang string) ([]PokemonJSON, error) {
	pokemons := []PokemonJSON{}
	pokeFile, pokeErr := pokeData.ReadFile(fmt.Sprintf("pokedata/pokedex.%s.json", lang))
	if pokeErr != nil {
		return nil, pokeErr
	}

	err := json.Unmarshal(pokeFile, &pokemons)
	if err != nil {
		return nil, err
	}

	return pokemons, nil
}
