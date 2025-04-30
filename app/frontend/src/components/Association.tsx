// Association.tsx
import React, { useCallback } from 'react';
import { createAssociation } from '../utils/createAssociation';
import { dia } from '@joint/core';

interface AssociationProps {
    source: { x: number; y: number };
    target: { x: number; y: number };
    layer: number;
    style?: dia.Link.Attributes['line']; // Optional line style
    marker?: dia.Link.Attributes['line']['targetMarker']; // Optional marker style
    onCreate: (link: dia.Link) => void;
}

const Association: React.FC<AssociationProps> = ({ source, target, layer, style, marker, onCreate }) => {
    const handleCreateAssociation = useCallback(() => {
        const link = createAssociation({ source, target, layer, style, marker });
        if (onCreate && link) {
            onCreate(link);
        } else {
            console.error('Failed to create association or onCreate is not defined.');
        }
    }, [source, target, layer, style, marker, onCreate]);

    return (
        <button
            onClick={handleCreateAssociation}
            style={{
                padding: '10px 20px',
                backgroundColor: '#007BFF',
                color: 'white',
                border: 'none',
                borderRadius: '6px',
                cursor: 'pointer',
                boxShadow: '0 2px 6px rgba(0,0,0,0.3)',
                fontWeight: 'bold',
            }}
        >
            Create Association
        </button>
    );
};

export default Association;
