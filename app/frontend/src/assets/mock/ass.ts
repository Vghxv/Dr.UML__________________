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

export const mockSelfAssociation: AssociationProps = {
    assType: 1,
    layer: 0,
    startX: 300,
    startY: 200,
    endX: 300,      // 對於 self-association，endX/endY 可與 startX/startY 相同
    endY: 200,
    deltaX: 60,     // 控制自連線的彎曲程度與方向
    deltaY: 80,
    attributes: [
        {
            content: "self",
            fontSize: 14,
            fontStyle: 0,
            fontFile: "Arial",
            ratio: 0.5
        }
    ]
};