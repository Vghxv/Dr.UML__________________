package main

import (
	"embed"

	"Dr.uml/backend/component"
	"Dr.uml/backend/umldiagram"
	"Dr.uml/backend/umlproject"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// TODO: handle error
	project, _ := umlproject.CreateEmptyUMLProject("NewProject")

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Dr.uml",
		Width:  1500,
		Height: 1000,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        project.Startup,
		Bind: []interface{}{
			project,
		},
		EnumBind: []interface{}{
			umldiagram.AllDiagramTypes,
			component.AllGadgetTypes,
			component.AllAssociationTypes,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
