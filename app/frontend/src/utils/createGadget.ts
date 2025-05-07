import { shapes, dia } from "@joint/core";
import { GadgetProps } from "../components/Gadget";

export interface BackendGadgetProps {
    gadgetType: string;
    x: number;
    y: number;
    layer: number;
    height: number;
    width: number;
    color: string;
    attributes: {
        content: string;
        height: number;
        width: number;
        fontSize: number;
        fontStyle: number;
        fontFile: string;
    }[][];
}

class UMLClass extends shapes.standard.Rectangle {
    constructor(options: GadgetProps) {
        console.log("Initializing UMLClass with options:", options);

        // Dynamically generate markup and attrs for attributes
        const attributeTexts = options.attributes.map((attr, i) => ({
            tagName: "text",
            selector: `attributeLabel${i}`,
            attributes: {}
        }));
        const attributeAttrs = options.attributes.reduce((acc, attr, i) => {
            acc[`attributeLabel${i}`] = {
                text: attr.content,
                refX: 5,
                refY: 35 + i * 20,
                textAnchor: "left",
                yAlignment: "top",
                fill: "#333333", // You can change this to attr.fontStyle -> color mapping
                fontFamily: attr.fontFile || "Arial",
                fontSize: attr.fontSize || 12,
            };
            return acc;
        }, {} as Record<string, any>);

        // Dynamically generate markup and attrs for methods
        const methodTexts = options.methods.map((method, i) => ({
            tagName: "text",
            selector: `methodLabel${i}`,
            attributes: {}
        }));
        const methodAttrs = options.methods.reduce((acc, method, i) => {
            acc[`methodLabel${i}`] = {
                text: method.content,
                refX: 5,
                refY: 30 + (options.height / 2) + 5 + i * 20,
                textAnchor: "left",
                yAlignment: "top",
                fill: "#333333",
                fontFamily: method.fontFile || "Arial",
                fontSize: method.fontSize || 12,
            };
            return acc;
        }, {} as Record<string, any>);

        super({
            position: { x: options.x, y: options.y },
            size: { width: options.width, height: options.height },
            attrs: {
                header: {
                    x: 0,
                    y: 0,
                    width: options.width,
                    height: 30,
                    fill: options.color || "#2ECC71",
                    stroke: "#000000",
                },
                headerLabel: {
                    ref: "header",
                    refX: "50%",
                    refY: "50%",
                    textAnchor: "middle",
                    yAlignment: "middle",
                    text: options.header || "Class Name",
                    fill: "#FFFFFF",
                    fontWeight: "bold",
                    fontFamily: options.header_atrributes.fontFile || "Arial",
                    fontSize: options.header_atrributes.fontSize || 12,
                },
                attributes: {
                    x: 0,
                    y: 30,
                    width: options.width,
                    height: options.height / 2,
                    fill: "#ECF0F1",
                    stroke: "#000000",
                },
                methods: {
                    x: 0,
                    y: 30 + options.height / 2,
                    width: options.width,
                    height: options.height / 2,
                    fill: "#ECF0F1",
                    stroke: "#000000",
                },
                ...attributeAttrs,
                ...methodAttrs
            },
            markup: [
                { tagName: "rect", selector: "header", attributes: {} },
                { tagName: "text", selector: "headerLabel", attributes: {} },
                { tagName: "rect", selector: "attributes", attributes: {} },
                { tagName: "rect", selector: "methods", attributes: {} },
                ...attributeTexts,
                ...methodTexts,
            ],
            z: options.layer,
        });
    }
}

export function createGadget(type: string, config: GadgetProps): dia.Element {
    switch (type) {
        case "Class": {
            return new UMLClass(config);
        }
        default:
            throw new Error(`Unknown gadget type: ${type}`);
    }
}

export function parseBackendGadget(gadgetData: BackendGadgetProps): dia.Element {
    console.log("Before Parse gadget:", gadgetData);
    const gadget = createGadget("Class", {
        gadgetType: gadgetData.gadgetType,
        x: gadgetData.x,
        y: gadgetData.y,
        layer: gadgetData.layer,
        height: gadgetData.height,
        width: gadgetData.width,
        color: gadgetData.color,
        header: gadgetData.attributes[0][0].content,
        header_atrributes: {
            content: gadgetData.attributes[0][0].content,
            height: gadgetData.attributes[0][0].height,
            width: gadgetData.attributes[0][0].width,
            fontSize: gadgetData.attributes[0][0].fontSize,
            fontStyle: gadgetData.attributes[0][0].fontStyle,
            fontFile: gadgetData.attributes[0][0].fontFile,
        },
        attributes: gadgetData.attributes[1].map(attr => ({
            content: attr.content,
            height: attr.height,
            width: attr.width,
            fontSize: attr.fontSize,
            fontStyle: attr.fontStyle,
            fontFile: attr.fontFile,
        })),
        methods: gadgetData.attributes[2].map(method => ({
            content: method.content,
            height: method.height,
            width: method.width,
            fontSize: method.fontSize,
            fontStyle: method.fontStyle,
            fontFile: method.fontFile,
        })),
    });
    console.log("Parsed gadget:", gadget);
    return gadget;
}
