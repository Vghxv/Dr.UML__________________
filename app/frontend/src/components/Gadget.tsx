import React, { useEffect, useRef } from 'react';
import { dia } from '@joint/core';
import { createGadget } from '../utils/createGadget';

export interface GadgetProps {
    gadgetType: string;
    x: number;
    y: number;
    layer: number;
    height: number;
    width: number;
    color: string;
    header: string;
    header_atrributes: {
        content: string;
        height: number;
        width: number;
        fontSize: number;
        fontStyle: number;
        fontFile: string;
    }
    attributes: {
        content: string;
        height: number;
        width: number;
        fontSize: number;
        fontStyle: number;
        fontFile: string;
    }[];
    methods: {
        content: string;
        height: number;
        width: number;
        fontSize: number;
        fontStyle: number;
        fontFile: string;
    }[];
}

const Gadget: React.FC<GadgetProps> = (props) => {
    const containerRef = useRef<HTMLDivElement>(null);

    useEffect(() => {
        if (containerRef.current) {
            // Create the gadget using the createGadget function
            const gadget = createGadget(props.gadgetType, props);

            // Render the gadget inside the container
            const paper = new dia.Paper({
                el: containerRef.current,
                model: new dia.Graph(),
                width: props.width,
                height: props.height,
                interactive: false, // Disable interactivity for rendering purposes
            });

            paper.model.addCell(gadget);
        }
    }, [props]);

    return (
        <div
            ref={containerRef}
            style={{
                position: 'absolute',
                left: props.x,
                top: props.y,
                width: props.width,
                height: props.height,
            }}
        />
    );
};

export default Gadget;


