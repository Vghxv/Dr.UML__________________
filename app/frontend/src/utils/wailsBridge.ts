import { EventsOn, EventsOff } from "../../wailsjs/runtime";

interface WindowWithGo extends Window {
    go: {
        umlproject: {
            UMLProject: {
                GetCurrentDiagramName(): Promise<string>;
                AddNewDiagram(diagramType: number, name: string): Promise<void>;
                SelectDiagram(name: string): Promise<void>;
                AddGadget(
                    gadgetType: number,
                    point: { x: number; y: number },
                    layer: number,
                    color: number
                ): Promise<void>;
                GetDrawData(): Promise<any>;
            };
        };
    };
}

declare var window: WindowWithGo;

// Fetch the current diagram name
export async function getCurrentDiagramName(): Promise<string> {
    try {
        return await window.go.umlproject.UMLProject.GetCurrentDiagramName();
    } catch (error) {
        console.error("Error fetching diagram name:", error);
        throw error;
    }
}

// Add a new gadget to the diagram
export async function addGadget(gadgetType: number, point: { x: number; y: number }, layer: number, color: number): Promise<void> {
    try {
        await window.go.umlproject.UMLProject.AddGadget(gadgetType, point, layer, color);
        console.log("Gadget add to backend:", gadgetType, point);
    } catch (error) {
        console.error("Error adding gadget:", error);
        throw error;
    }
}

// Get canvas data without adding a gadget
export async function getCanvasData(): Promise<any> {
    try {
        return await window.go.umlproject.UMLProject.GetDrawData();
    } catch (error) {
        console.error("Error fetching canvas data:", error);
        throw error;
    }
}

// Register a backend event listener
export function onBackendEvent(eventName: string, callback: (result: any) => void): void {
    console.log("Registering event listener for:", eventName);
    EventsOn(eventName, callback);
}

// Unregister a backend event listener
export function offBackendEvent(eventName: string): void {
    console.log("Unregistering event listener for:", eventName);
    EventsOff(eventName);
}
