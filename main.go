package main

import (
	"code-fox/app"
	"code-fox/database"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"log"
)

var assets embed.FS

func main() {
	err := database.InitializeDatabaseObjects()
	if err != nil {
		log.Fatal(err)
	}

	application := app.NewApp()
	snippetManager := app.NewSnippetManager()
	tagManager := app.NewTagManager()

	err = wails.Run(&options.App{
		Title:  "CodeFox",
		Width:  1600,
		Height: 900,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        application.Startup,
		Bind: []interface{}{
			application,
			snippetManager,
			tagManager,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
