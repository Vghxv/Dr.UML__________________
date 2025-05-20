import { RefObject } from "react";
import { ToPoint } from "../utils/wailsBridge";
import { SelectComponent } from "../../wailsjs/go/umlproject/UMLProject";

export function useCanvasMouseEvents(
    canvasRef: RefObject<HTMLCanvasElement>,
    onSelect: () => void
) {
    const handleMouseDown = (event: React.MouseEvent<HTMLCanvasElement>) => {
        const rect = canvasRef.current?.getBoundingClientRect();
        const x = Math.round(event.clientX - (rect?.left || 0));
        const y = Math.round(event.clientY - (rect?.top || 0));
        SelectComponent(ToPoint(x, y))
            .then(() => {
                onSelect();
            })
            .catch((error) => {
                console.error("Error selecting component:", error);
            });
    };

    const handleMouseMove = (event: React.MouseEvent<HTMLCanvasElement>) => {
        // 可擴充 hover 效果
    };

    const handleMouseUp = () => {
        // 可擴充拖曳等功能
    };

    return {
        handleMouseDown,
        handleMouseMove,
        handleMouseUp
    };
}
