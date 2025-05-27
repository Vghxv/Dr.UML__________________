import {EventsOff, EventsOn} from "../../wailsjs/runtime";
import {utils} from "../../wailsjs/go/models";

export function ToPoint(x: number, y: number) {
    return new utils.Point({
        X: x,
        Y: y
    });
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
