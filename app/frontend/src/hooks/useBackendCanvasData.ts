import { useEffect, useState, useCallback } from "react";
import { offBackendEvent, onBackendEvent } from "../utils/wailsBridge";
import { GetDrawData } from "../../wailsjs/go/umlproject/UMLProject";
import { CanvasProps, GadgetProps, AssociationProps } from "../utils/Props";

// Type definitions for backend data
interface BackendGadget {
    gadgetType: any;
    x: number;
    y: number;
    layer: number;
    height: number;
    width: number;
    color: string;
    isSelected: boolean;
    attributes: any[][];
}

interface BackendAssociation {
    assType: number;
    layer: number;
    startX: number;
    startY: number;
    endX: number;
    endY: number;
    deltaX: number;
    deltaY: number;
    isSelected: boolean;
    attributes?: any[];
}

interface BackendDiagram {
    margin: number;
    color: string;
    lineWidth: number;
    gadgets: BackendGadget[];
    associations?: BackendAssociation[];
}

/**
 * Transforms a backend association attribute to frontend format
 */
const transformAssociationAttribute = (attr: any) => ({
    content: attr.content,
    fontSize: attr.fontSize,
    fontStyle: attr.fontStyle,
    fontFile: attr.fontFile,
    ratio: attr.ratio,
    height: attr.height
});

/**
 * Transforms a backend gadget to frontend format
 */
const transformGadget = (gadget: BackendGadget): GadgetProps => ({
    gadgetType: gadget.gadgetType.toString(),
    x: gadget.x,
    y: gadget.y,
    layer: gadget.layer,
    height: gadget.height,
    width: gadget.width,
    color: gadget.color,
    isSelected: gadget.isSelected,
    attributes: gadget.attributes
});

/**
 * Transforms a backend association to frontend format
 */
const transformAssociation = (association: BackendAssociation): AssociationProps => ({
    assType: association.assType,
    layer: association.layer,
    startX: association.startX,
    startY: association.startY,
    endX: association.endX,
    endY: association.endY,
    deltaX: association.deltaX,
    deltaY: association.deltaY,
    isSelected: association.isSelected,
    attributes: association.attributes?.map(transformAssociationAttribute) || []
});

/**
 * Transforms backend diagram data to frontend canvas props format
 */
const transformDiagramToCanvasProps = (diagram: BackendDiagram): CanvasProps => ({
    margin: diagram.margin,
    color: diagram.color,
    lineWidth: diagram.lineWidth,
    gadgets: diagram.gadgets?.map(transformGadget) || [],
    associations: diagram.associations?.map(transformAssociation) || []
});

export function useBackendCanvasData() {
    const [backendData, setBackendData] = useState<CanvasProps | null>(null);

    const reloadBackendData = useCallback(async () => {
        try {
            const diagram = await GetDrawData();
            const canvasData = transformDiagramToCanvasProps(diagram);
            setBackendData(canvasData);
        } catch (error) {
            console.error("Error loading canvas data:", error);
        }
    }, []);

    useEffect(() => {
        // Initial data load
        reloadBackendData().then(() => console.log("Canvas data loaded successfully"));

        // Set up backend event listener
        const handleBackendEvent = (result: BackendDiagram) => {
            if (result) {
                console.log("Received backend event data:", result);
                const canvasData = transformDiagramToCanvasProps(result);
                setBackendData(canvasData);
            }
        };

        onBackendEvent("backend-event", handleBackendEvent);

        // Cleanup function
        return () => {
            offBackendEvent("backend-event");
        };
    }, [reloadBackendData]);

    return { backendData, reloadBackendData };
}
