// createAssociation.ts
import { shapes, dia } from '@joint/core';

interface AssociationOptions {
    source: { x: number; y: number };
    target: { x: number; y: number };
    layer: number;
}

interface BackendAssAttribute {
    content: string;
    fontSize: number;
    fontStyle: number;
    fontFile: string;
    ratio: number;
}

interface BackendAssociation {
    assType: string;
    layer: number;
    startX: number;
    startY: number;
    endX: number;
    endY: number;
    attributes: BackendAssAttribute[];
}

export function createAssociation({
    source,
    target,
    layer,
    style = { stroke: '#333333', strokeWidth: 2 }, // Default line style
    marker = {
        type: 'path',
        d: 'M 10 -5 0 0 10 5 Z',
        fill: '#333333',
    }, // Default marker
}: AssociationOptions & { style?: dia.Link.Attributes['line']; marker?: dia.Link.Attributes['line']['targetMarker'] }): dia.Link {
    const link = new shapes.standard.Link({
        source: { x: source.x, y: source.y },
        target: { x: target.x, y: target.y },
        attrs: {
            line: {
                ...style,
                targetMarker: marker,
            },
        },
        z: layer,
    });

    return link;
}

// Parse backend JSON and create an association
export function parseBackendAssociation(json: string): dia.Link {
    const data: BackendAssociation = JSON.parse(json);

    const link = createAssociation({
        source: { x: data.startX, y: data.startY },
        target: { x: data.endX, y: data.endY },
        layer: data.layer,
        style: { stroke: '#333333', strokeWidth: 2 }, // Default style
        marker: {
            type: 'path',
            d: 'M 10 -5 0 0 10 5 Z',
            fill: '#333333',
        },
    });

    // Add attributes as labels
    const labels = data.attributes.map((attr) => ({
        position: attr.ratio,
        attrs: {
            text: {
                text: attr.content,
                fontSize: attr.fontSize,
                fontStyle: attr.fontStyle,
                fontFamily: attr.fontFile,
                fill: '#000000',
            },
        },
    }));

    link.set('labels', labels);

    return link;
}
