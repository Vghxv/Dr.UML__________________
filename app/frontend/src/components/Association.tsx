// Association.tsx
import React, { useCallback } from 'react';
import { createAssociation, parseBackendAssociation } from '../utils/createAssociation';
import { dia } from '@joint/core';

interface AssociationProps {
    source?: { x: number; y: number }; // Optional if backendJson is provided
    target?: { x: number; y: number }; // Optional if backendJson is provided
    layer?: number; // Optional if backendJson is provided
    style?: dia.Link.Attributes['line']; // Optional line style
    marker?: dia.Link.Attributes['line']['targetMarker']; // Optional marker style
    backendJson?: string; // Optional backend JSON string
    onCreate: (link: dia.Link) => void;
}

const Association: React.FC<AssociationProps> = ({
    source,
    target,
    layer,
    style,
    marker,
    backendJson,
    onCreate,
}) => {
    const handleCreateAssociation = useCallback(() => {
        let link: dia.Link | null = null;

        if (backendJson) {
            // Parse backend JSON if provided
            link = parseBackendAssociation(backendJson);
        } else if (source && target && layer !== undefined) {
            // Otherwise, create an association normally
            link = createAssociation({ source, target, layer, style, marker });
        } else {
            console.error('Insufficient data to create an association.');
        }

        if (link) {
            onCreate(link); // Ensure the onCreate function is called with the created link
        } else {
            console.error('Failed to create association.');
        }
    }, [source, target, layer, style, marker, backendJson, onCreate]);

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

