import React, { RefObject } from "react";
import { ToPoint } from "../utils/wailsBridge";
import { SelectComponent } from "../../wailsjs/go/umlproject/UMLProject";

interface ExtraHandlers {
    onCanvasClick?: (point: { x: number, y: number }) => void;
    isAddingAssociation?: boolean;
}

export function useCanvasMouseEvents(
    canvasRef: RefObject<HTMLCanvasElement>,
    onSelect: () => void,
    extraHandlers?: ExtraHandlers
) {
    const handleMouseDown = (event: React.MouseEvent<HTMLCanvasElement>) => {
        const canvas = canvasRef.current;
        if (!canvas) return;

        const rect = canvas.getBoundingClientRect();
        // Calculate the scaling factor between the CSS size and the canvas's internal size
        const scaleX = canvas.width / rect.width;
        const scaleY = canvas.height / rect.height;

        // Apply the scaling to get the correct coordinates in the canvas's coordinate system
        const x = Math.round((event.clientX - rect.left) * scaleX);
        const y = Math.round((event.clientY - rect.top) * scaleY);

        if (extraHandlers?.isAddingAssociation && extraHandlers?.onCanvasClick) {
            extraHandlers.onCanvasClick({ x, y });
        } else {
            SelectComponent(ToPoint(x, y))
                .then(() => {
                    onSelect();
                })
                .catch((error) => {
                    console.error("Error selecting component:", error);
                });
        }
    };

    const handleMouseMove = (event: React.MouseEvent<HTMLCanvasElement>) => {
        // TODO: hovering
    };

    const handleMouseUp = () => {
        // TODO
    };

    return {
        handleMouseDown,
        handleMouseMove,
        handleMouseUp
    };
}
