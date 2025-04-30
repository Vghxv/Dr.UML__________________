package main

import (
	_ "fmt"
)

// App struct
// type App struct {
// 	ctx     context.Context
// 	project *umlproject.UMLProject
// }

// // NewApp creates a new App application struct
// func NewApp() *App {
// 	return &App{}
// }

// // startup is called when the app starts. The context is saved
// // so we can call the runtime methods
// func (a *App) startup(ctx context.Context) {
// 	a.ctx = ctx
// 	a.project = umlproject.NewUMLProject("test")
// 	a.project.AddNewDiagram(umldiagram.ClassDiagram, "new class diagram")
// 	a.project.RegisterNotifyDrawUpdate(a.ProcessWithCallback)
// 	a.project.SelectDiagram("new class diagram")
// }

// func (a *App) ProcessWithCallback(callbackID string) duerror.DUError {

// 	// Call the JavaScript callback function with the result
// 	runtime.EventsEmit(a.ctx, callbackID)

// 	// Return a nil error to satisfy the expected signature
// 	return nil
// }

// func (a *App) GetCurrentDiagramName() (string, duerror.DUError) {
// 	if a.project == nil {
// 		return "", duerror.NewInvalidArgumentError("project is nil")
// 	}
// 	if a.project.GetCurrentDiagram() == nil {
// 		return "", duerror.NewInvalidArgumentError("current diagram is nil")
// 	}
// 	return a.project.GetCurrentDiagram().GetName(), nil
// }

// func (a *App) AddGadget(
// 	gadgetType component.GadgetType,
// 	point utils.Point,
// ) duerror.DUError {

// 	err := a.project.AddGadget(gadgetType, point)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
