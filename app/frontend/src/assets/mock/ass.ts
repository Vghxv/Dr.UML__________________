import { AssociationProps } from "../../utils/Props";

export const mockAssociation: AssociationProps = {
    assType: 1,
    layer: 0,
    startX: 100,
    startY: 100,
    endX: 500,
    endY: 500,
    deltaX: 0,
    deltaY: 0,
    attributes: [
        {
            content: "1",
            fontSize: 14,
            fontStyle: 0,
            fontFile: "Arial",
            ratio: 0.1
        },
        {
            content: "0..*",
            fontSize: 14,
            fontStyle: 0,
            fontFile: "Arial",
            ratio: 0.9
        },
        {
            content: "Order",
            fontSize: 16,
            fontStyle: 1,
            fontFile: "Arial",
            ratio: 0.5
        }
    ]
};
