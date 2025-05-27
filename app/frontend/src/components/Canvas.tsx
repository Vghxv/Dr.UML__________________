import React, {useEffect, useRef, useCallback} from 'react';
import {CanvasProps, GadgetProps} from '../utils/Props';
import {createGadget} from '../utils/createGadget';
import {createAss} from '../utils/createAssociation';
import {useCanvasMouseEvents} from '../hooks/useCanvasMouseEvents';
import {useSelection} from '../hooks/useSelection';

const DrawingCanvas: React.FC<{
    backendData: CanvasProps | null,
    reloadBackendData?: () => void,
    onSelectionChange?: (selectedGadget: GadgetProps | null, selectedGadgetCount: number) => void
}> = ({
          backendData,
          reloadBackendData,
          onSelectionChange
      }) => {
    const canvasRef = useRef<HTMLCanvasElement>(null);
    const {selectedGadgetCount, selectedGadget} = useSelection(backendData?.gadgets);

    const redrawCanvas = useCallback(() => {
        const canvas = canvasRef.current;
        if (canvas) {
            const ctx = canvas.getContext('2d');
            if (ctx) {
                ctx.clearRect(0, 0, canvas.width, canvas.height);

                backendData?.gadgets?.forEach((gadget: GadgetProps) => {
                    const gad = createGadget("Class", gadget, backendData.margin);
                    gad.draw(ctx, backendData.margin, backendData.lineWidth);
                });

                backendData?.Association?.forEach((association) => {
                    const ass = createAss("Association", association, backendData.margin);
                    ass.draw(ctx, backendData.margin, backendData.lineWidth);
                });
            }
        }
    }, [backendData]);

    const resizeCanvas = useCallback(() => {
        const canvas = canvasRef.current;
        if (canvas) {
            const parent = canvas.parentElement;
            if (parent) {
                canvas.width = parent.clientWidth;
                canvas.height = parent.clientHeight;
                redrawCanvas();
            }
        }
    }, [canvasRef, redrawCanvas]);

    useEffect(() => {
        redrawCanvas();
    }, [redrawCanvas]); // Only depend on redrawCanvas, which already depends on backendData

    useEffect(() => {
        if (onSelectionChange) {
            onSelectionChange(selectedGadget, selectedGadgetCount);
        }
    }, [selectedGadget, selectedGadgetCount, onSelectionChange]);
    useEffect(() => {
        resizeCanvas();
    }, [resizeCanvas]); // Use resizeCanvas, which already includes redrawCanvas

    // Add a resize event listener to handle viewport changes
    useEffect(() => {
        window.addEventListener('resize', resizeCanvas);

        // Clean up the event listener on a component unmount
        return () => {
            window.removeEventListener('resize', resizeCanvas);
        };
    }, [resizeCanvas]); // Only depend on resizeCanvas, which includes necessary dependencies

    const {handleMouseDown, handleMouseMove, handleMouseUp} = useCanvasMouseEvents(
        canvasRef,
        () => {
            if (reloadBackendData) {
                reloadBackendData();
            }
        }
    );


    return (
        // <div className="relative flex">
            <canvas
                ref={canvasRef}
                className="border-2 border-neutral-600 rounded-lg bg-neutral-900 shadow-lg w-full h-[calc(100vh-170px)] m-0 relative"
                onMouseDown={handleMouseDown}
                onMouseMove={handleMouseMove}
                onMouseUp={handleMouseUp}
            />
        // </div>
    );
};

export default DrawingCanvas;
