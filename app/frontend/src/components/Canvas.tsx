import React, { useEffect, useRef } from 'react';
import { CanvasProps, GadgetProps } from '../utils/Props';
import { createGadget } from '../utils/createGadget';
import { createAss } from '../utils/createAssociation';
import { useCanvasMouseEvents } from '../hooks/useCanvasMouseEvents';
import { useSelection } from '../hooks/useSelection';

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
    const { selectedGadgetCount, selectedGadget } = useSelection(backendData?.gadgets);

    const redrawCanvas = () => {
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
    };

    useEffect(() => {
        redrawCanvas();
    }, [backendData]);

    useEffect(() => {
        if (onSelectionChange) {
            onSelectionChange(selectedGadget, selectedGadgetCount);
        }
    }, [selectedGadget, selectedGadgetCount, onSelectionChange]);

    const { handleMouseDown, handleMouseMove, handleMouseUp } = useCanvasMouseEvents(
        canvasRef,
        () => {
            if (reloadBackendData) {
                reloadBackendData();
            }
        }
    );


    return (
        <div style={{position: 'relative', display: 'flex'}}>
            <canvas
                ref={canvasRef}
                style={{
                    border: '2px solid #444',
                    borderRadius: '8px',
                    backgroundColor: '#1e1e1e',
                    boxShadow: '0 4px 8px rgba(0, 0, 0, 0.5)',
                    margin: '20px auto',
                    position: 'relative',
                }}
                width="1200"
                height="700"
                onMouseDown={handleMouseDown}
                onMouseMove={handleMouseMove}
                onMouseUp={handleMouseUp}
            />
        </div>
    );
};

export default DrawingCanvas;
