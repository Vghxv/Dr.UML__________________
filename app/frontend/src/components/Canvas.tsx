import React, {useCallback, useEffect, useRef} from 'react';
import {AssociationProps, CanvasProps, GadgetProps} from '../utils/Props';
import {createGad} from '../utils/createGadget';
import {createAss} from '../utils/createAssociation';
import {useCanvasMouseEvents} from '../hooks/useCanvasMouseEvents';
import {useSelection} from '../hooks/useSelection';
import { mockAssociation } from '../assets/mock/ass';

const DrawingCanvas: React.FC<{
    backendData: CanvasProps | null,
    reloadBackendData?: () => void,
    onSelectionChange?: (selectedComponent: GadgetProps | AssociationProps | null, selectedComponentCount: number) => void,
    onCanvasClick?: (point: { x: number, y: number }) => void,
    isAddingAssociation?: boolean
}> = ({
          backendData,
          reloadBackendData,
          onSelectionChange,
          onCanvasClick,
          isAddingAssociation
      }) => {
    const canvasRef = useRef<HTMLCanvasElement>(null);
    const {selectedComponentCount, selectedComponent} = useSelection(backendData?.gadgets);

    const redrawCanvas = useCallback(() => {
        const canvas = canvasRef.current;
        if (canvas) {
            const ctx = canvas.getContext('2d');
            if (ctx) {
                ctx.clearRect(0, 0, canvas.width, canvas.height);
                
                backendData?.gadgets?.forEach((gadget: GadgetProps) => {
                    const gad = createGad("Class", gadget, backendData.margin);
                    gad.draw(ctx, backendData.margin, backendData.lineWidth);
                });

                backendData?.associations?.forEach((association: AssociationProps) => {
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
            onSelectionChange(selectedComponent, selectedComponentCount);
        }
    }, [selectedComponent, selectedComponentCount, onSelectionChange]);
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

    // 使用 useCanvasMouseEvents 處理所有 mouse event
    const {handleMouseDown, handleMouseMove, handleMouseUp} = useCanvasMouseEvents(
        canvasRef,
        () => {
            if (reloadBackendData) {
                reloadBackendData();
            }
        },
        {
            onCanvasClick,
            isAddingAssociation
        }
    );

    return (
        <canvas
            ref={canvasRef}
            className="border-2 border-neutral-600 rounded-lg bg-neutral-900 shadow-lg w-full h-[calc(100vh-170px)] m-0 relative"
            onMouseDown={handleMouseDown}
            onMouseMove={handleMouseMove}
            onMouseUp={handleMouseUp}
        />
    );
};

export default DrawingCanvas;
