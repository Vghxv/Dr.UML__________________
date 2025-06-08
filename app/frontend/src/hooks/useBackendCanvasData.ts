import { useEffect, useState, useCallback } from "react";
import { offBackendEvent, onBackendEvent } from "../utils/wailsBridge";
import { GetDrawData } from "../../wailsjs/go/umlproject/UMLProject";
import { CanvasProps } from "../utils/Props";
// import { mockAssociation } from "../assets/mock/ass";

export function useBackendCanvasData() {
    const [backendData, setBackendData] = useState<CanvasProps | null>(null);

    const reloadBackendData = useCallback(async () => {
        try {
            const diagram = await GetDrawData();
            const canvasData: CanvasProps = {
                margin: diagram.margin,
                color: diagram.color,
                lineWidth: diagram.lineWidth,
                gadgets: diagram.gadgets.map(gadget => ({
                    gadgetType: gadget.gadgetType.toString(),
                    x: gadget.x,
                    y: gadget.y,
                    layer: gadget.layer,
                    height: gadget.height,
                    width: gadget.width,
                    color: gadget.color,
                    isSelected: gadget.isSelected,
                    attributes: gadget.attributes
                })),
                associations: diagram.associations?.map(association => ({
                    assType: association.assType,
                    layer: association.layer,
                    startX: association.startX,
                    startY: association.startY,
                    endX: association.endX,
                    endY: association.endY,
                    deltaX: association.deltaX,
                    deltaY: association.deltaY,
                    isSelected: association.isSelected,
                    attributes: association.attributes
                }))
            };
            // if (!canvasData.associations || canvasData.associations.length === 0) {
            //     canvasData.associations = [mockAssociation];
            // }
            setBackendData(canvasData);
        } catch (error) {
            console.error("Error loading canvas data:", error);
        }
    }, []);

    useEffect(() => {
        reloadBackendData().then(r => console.log("Loaded canvas data:", r));
        onBackendEvent("backend-event", (result) => {
            if (result) {
                const canvasData: CanvasProps = {
                    margin: result.margin,
                    color: result.color,
                    lineWidth: result.lineWidth,
                    gadgets: result.gadgets?.map((gadget: any) => ({
                        gadgetType: gadget.gadgetType.toString(),
                        x: gadget.x,
                        y: gadget.y,
                        layer: gadget.layer,
                        height: gadget.height,
                        width: gadget.width,
                        color: gadget.color,
                        isSelected: gadget.isSelected,
                        attributes: gadget.attributes
                    })),
                    associations: result.associations?.map((association: any) => ({
                        assType: association.assType,
                        layer: association.layer,
                        startX: association.startX,
                        startY: association.startY,
                        endX: association.endX,
                        endY: association.endY,
                        deltaX: association.deltaX,
                        deltaY: association.deltaY,
                        attributes: association.attributes.map((attr: any) => ({
                            content: attr.content,
                            fontSize: attr.fontSize,
                            fontStyle: attr.fontStyle,
                            fontFile: attr.fontFile,
                            ratio: attr.ratio
                        }))
                    }))
                };
                setBackendData(canvasData);
            }
        });
        return () => {
            offBackendEvent("backend-event");
        };
    }, [reloadBackendData]);

    return { backendData, reloadBackendData };
}
