export interface CanvasProps {
    margin: number;
    color: number;
    lineWidth: number;
    gadgets?: GadgetProps[];
    Association?: AssociationProps[];
}

export interface GadgetProps {
    gadgetType: string;
    x: number;
    y: number;
    layer: number;
    height: number;
    width: number;
    color: string;
    isSelected: boolean;
    attributes: {
        content: string;
        height: number;
        width: number;
        fontSize: number;
        fontStyle: number;
        fontFile: string;
    }[][];
}

export interface AssociationProps {
    assType: number;
    layer: number;
    startX: number;
    startY: number;
    endX: number;
    endY: number;
    deltaX: number;
    deltaY: number;
    attributes: {
        content: string;
        fontSize: number;
        fontStyle: number;
        fontFile: string;
        ratio: number;
    }[];
}