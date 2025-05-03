import React, { useEffect, useRef, useState } from 'react';
import { dia, shapes } from '@joint/core';
import { parseBackendGadget, ComponentProps} from '../utils/createComponent';

interface CanvasProps {
    graph: dia.Graph;
    backendData: any; // JSON data for gadgets and associations
}

const Canvas: React.FC<CanvasProps> = ({ graph, backendData }) => {
    const paperRef = useRef<HTMLDivElement>(null);
    const [selectedElements, setSelectedElements] = useState<dia.Element[]>([]);

    useEffect(() => {
        if (paperRef.current) {
            // Initialize the Paper
            const paper = new dia.Paper({
                el: paperRef.current,
                model: graph,
                width: 800,
                height: 600,
                gridSize: 10,
                drawGrid: true,
                interactive: true, // Enable interactivity
            });

            // Handle mouse interactions
            paper.on('element:pointerclick', (elementView) => {
                const element = elementView.model;
                if (selectedElements.includes(element)) {
                    setSelectedElements(selectedElements.filter((el) => el !== element));
                } else {
                    setSelectedElements([...selectedElements, element]);
                }
            });

            paper.on('blank:pointerclick', () => {
                setSelectedElements([]); // Deselect all elements when clicking on blank space
            });
        }
    }, [graph, selectedElements]);

    useEffect(() => {
        // Parse backend data and add gadgets and associations to the graph
        if (backendData) {
            backendData.gadgets.forEach((gadgetData: any) => {
                const gadget = parseBackendGadget(gadgetData);
                graph.addCell(gadget);
            });
        }
    }, [backendData, graph]);

    return (
        <div
            ref={paperRef}
            style={{
                border: '2px solid #444',
                borderRadius: '8px',
                backgroundColor: '#1e1e1e',
                boxShadow: '0 4px 8px rgba(0, 0, 0, 0.5)',
                width: '800px',
                height: '600px',
                margin: '20px auto',
                position: 'relative',
            }}
        />
    );
};

export default Canvas;