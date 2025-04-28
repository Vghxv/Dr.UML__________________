import React, {useEffect, useRef} from 'react';
import {useDrop} from 'react-dnd';
import {ItemTypes} from '../types';
import {dia, shapes} from '@joint/core';

const Canvas: React.FC<{ graph: dia.Graph }> = ({ graph }) => {
    const paperRef = useRef<HTMLDivElement>(null);

    useEffect(() => {
        if (paperRef.current) {
            // Initialize the Paper
            new dia.Paper({
                el: paperRef.current,
                model: graph,
                width: 800, // Set the width of the canvas
                height: 600, // Set the height of the canvas
                gridSize: 10, // Set the grid size
                drawGrid: true // Enable grid drawing
            });
        }
    }, [graph]);

    return (
        <div
            ref={paperRef}
            style={{
                border: '2px solid #444', // Dark gray border
                borderRadius: '8px', // Rounded corners
                backgroundColor: '#1e1e1e', // Dark background
                boxShadow: '0 4px 8px rgba(0, 0, 0, 0.5)', // Stronger shadow
                width: '800px', // Fixed width
                height: '600px', // Fixed height
                margin: '20px auto' // Center alignment
            }}
        />
    );
};

export default Canvas;