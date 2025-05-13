import React, { useRef, useEffect, useState } from 'react';
import { CanvasProps, GadgetProps } from '../utils/Props';
import { createGadget } from './createGadget';

const DrawingCanvas: React.FC<{ backendData: CanvasProps | null }> = ({ backendData }) => {
    const canvasRef = useRef<HTMLCanvasElement>(null);

    const redrawCanvas = () => {
        const canvas = canvasRef.current;
        if (canvas) {
            const ctx = canvas.getContext('2d');
            if (ctx) {
                ctx.clearRect(0, 0, canvas.width, canvas.height);
                backendData?.gadgets?.forEach((gadget: GadgetProps) => {
                    const gad = createGadget("Class", gadget, backendData.margin);
                    gad.draw(ctx, backendData.margin, backendData.lineWidth);
                }
                );
            }
        }
    }

    useEffect(() => {
        redrawCanvas();

    }, [backendData]);

    const handleMouseDown = (event: React.MouseEvent<HTMLCanvasElement>) => {
        // TODO: call backend
    };

    const handleMouseMove = (event: React.MouseEvent<HTMLCanvasElement>) => {
        // TODO: add some hover things
    };

    const handleMouseUp = () => {
    };

    return (
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
            width="800"
            height="600"
            onMouseDown={handleMouseDown}
            onMouseMove={handleMouseMove}
            onMouseUp={handleMouseUp}
        />
    );
};

export default DrawingCanvas;
