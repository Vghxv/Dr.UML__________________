// createAssociation.ts
import { shapes, dia } from '@joint/core';

interface AssociationOptions {
    source: { x: number; y: number };
    target: { x: number; y: number };
    layer: number;
}

export function createAssociation({ source, target, layer }: AssociationOptions): dia.Link {
    const link = new shapes.standard.Link({
        source: { x: source.x, y: source.y },
        target: { x: target.x, y: target.y },
        attrs: {
            line: {
                stroke: '#333333',
                strokeWidth: 2,
                targetMarker: {
                    type: 'path',
                    d: 'M 10 -5 0 0 10 5 Z',
                    fill: '#333333',
                },
            },
        },
        z: layer,
    });

    return link;
}
