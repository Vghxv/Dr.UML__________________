import React, { useEffect, useRef, useState } from 'react';
import { dia, shapes } from '@joint/core';
import { parseBackendGadget, BackendGadgetProps} from '../utils/createGadget';
import { BackendCanvasProps, createCanvas } from '../utils/createCanvas';

export interface CanvasProps {
    graph: dia.Graph;
    paper: dia.Paper;
    backendData?: BackendCanvasProps; // JSON data for gadgets and associations
}

const Canvas: React.FC<{ backendData : BackendCanvasProps}> = ({ backendData }) => {
    const paperRef = useRef<HTMLDivElement>(null);
    const [graph, setGraph] = useState<dia.Graph | null>(null);
    const [paper, setPaper] = useState<dia.Paper | null>(null);
    const [selectedElements, setSelectedElements] = useState<dia.Element[]>([]);

    // TODO: remove this part, since this moniters individual element, should monitor canvas onclick instead
    useEffect(() => {
        if (graph && paper) {
            paper.on('element:pointerclick', (elementView) => {
                const element = elementView.model;
                if (selectedElements.includes(element)) {
                    setSelectedElements(selectedElements.filter((el) => el !== element));
                } else {
                    setSelectedElements([...selectedElements, element]);
                }
            });

            paper.on('blank:pointerclick', () => {
                setSelectedElements([]);
            });
        }
    }, [graph, paper, selectedElements]);

    useEffect(() => {
        if (paperRef.current) {
            const { graph, paper } = createCanvas(paperRef.current, backendData);
            setGraph(graph);
            setPaper(paper);
        }
    }, [backendData]);

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