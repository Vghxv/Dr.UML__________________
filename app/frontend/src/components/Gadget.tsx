import React, { useCallback } from 'react';
import { createGadget } from '../utils/createGadget';
import { dia } from '@joint/core';

interface GadgetProps {
    point: { x: number; y: number };
    type: 'Class';
    layer: number;
    size?: { width: number; height: number };
    color?: string;
    outlineColor?: string;
    name?: string;
    onDrop: (gadget: dia.Element) => void;
}

const Gadget: React.FC<GadgetProps> = ({
    point,
    type,
    layer,
    size = { width: 100, height: 60 },
    color = '#f0f0f0',
    outlineColor = '#007BFF',
    name = '',
    onDrop,
}) => {
    const handleCreateGadget = useCallback(() => {
        const gadget = createGadget({
            point,
            type,
            layer,
            size,
            color,
            outlineColor,
            name,
        });

        if (onDrop && gadget) {
            onDrop(gadget); // Ensure the onDrop function is called with the created gadget
        } else {
            console.error('Failed to create gadget or onDrop is not defined.');
        }
    }, [point, type, layer, size, color, outlineColor, name, onDrop]);

    return (
        <div
            onClick={handleCreateGadget}
            style={{
                width: size.width,
                height: size.height,
                backgroundColor: color,
                border: `2px solid ${outlineColor}`,
                borderRadius:  '8px',
                cursor: 'pointer',
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
                boxShadow: '0 4px 8px rgba(0, 0, 0, 0.2)',
                transition: 'transform 0.2s, box-shadow 0.2s',
            }}
            onMouseEnter={(e) => {
                e.currentTarget.style.transform = 'scale(1.05)';
                e.currentTarget.style.boxShadow = '0 6px 12px rgba(0, 0, 0, 0.3)';
            }}
            onMouseLeave={(e) => {
                e.currentTarget.style.transform = 'scale(1)';
                e.currentTarget.style.boxShadow = '0 4px 8px rgba(0, 0, 0, 0.2)';
            }}
        >
            <span style={{ color: '#333', fontSize: '14px', fontWeight: 'bold' }}>{name}</span>
        </div>
    );
};

export default Gadget;
