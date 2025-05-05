import { shapes, dia } from "@joint/core";
import { CanvasProps } from "../components/Canvas";
import { BackendGadgetProps, parseBackendGadget } from "./createGadget";

export interface BackendCanvasProps {
    margin: number;
    color: number;
    lineWidth: number;
    gadgets?: BackendGadgetProps[];
}

export const createCanvas = (
    el: HTMLElement,
    canvasProps: BackendCanvasProps
): { graph: dia.Graph; paper: dia.Paper } => {
    const graph = new dia.Graph();

    const paper = new dia.Paper({
        el: el,
        model: graph,
        width: 800,
        height: 600,
        gridSize: 10,
        drawGrid: true,
        interactive: true,
    });

    // Add gadgets to the canvas
    if (canvasProps.gadgets) {
        canvasProps.gadgets.forEach((gadgetData: BackendGadgetProps) => {
            const gadget = parseBackendGadget(gadgetData);
            graph.addCell(gadget);
        });
    }
    return { graph, paper };
};
