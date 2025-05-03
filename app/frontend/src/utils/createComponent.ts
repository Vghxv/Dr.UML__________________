import { shapes, dia } from "@joint/core";
import { GadgetProps } from "../components/Gadget";

export interface ComponentProps {
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
                    fontFamily: options.header_atrributes.fontFile || "normal",
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
                attributesLabel: {
                    ref: "attributes",
                    refX: 5,
                    refY: 5,
                    textAnchor: "left",
                    yAlignment: "top",
                    text: options.attributes
                        .map(attr => attr.content)
                        .join("\n") || "Attributes",
                    fill: "#333333",
                    fontFamily: options.attributes[0]?.fontFile || "normal",
                    fontSize: options.attributes[0]?.fontSize || 12,
                },
                methods: {
                    x: 0,
                    y: 30 + options.height / 2,
                    width: options.width,
                    height: options.height / 2,
                    fill: "#ECF0F1",
                    stroke: "#000000",
                },
                methodsLabel: {
                    ref: "methods",
                    refX: 5,
                    refY: 5,
                    textAnchor: "left",
                    yAlignment: "top",
                    text: options.methods
                        .map(method => method.content)
                        .join("\n") || "Methods",
                    fill: "#333333",
                    fontFamily: options.methods[0]?.fontFile || "normal",
                    fontSize: options.methods[0]?.fontSize || 12,
                },
            },
            markup: [
                { tagName: "rect", selector: "header", attributes: {} },
                { tagName: "text", selector: "headerLabel", attributes: {} },
                { tagName: "rect", selector: "attributes", attributes: {} },
                { tagName: "text", selector: "attributesLabel", attributes: {} },
                { tagName: "rect", selector: "methods", attributes: {} },
                { tagName: "text", selector: "methodsLabel", attributes: {} },
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

// Parse backend gadget JSON and convert it to a dia.Element
export function parseBackendGadget(gadgetData: ComponentProps): dia.Element {
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
        attributes: gadgetData.attributes.map(attr => ({
            content: attr[1].content,
            height: attr[1].height,
            width: attr[1].width,
            fontSize: attr[1].fontSize,
            fontStyle: attr[1].fontStyle,
            fontFile: attr[1].fontFile,
        })),
        methods: gadgetData.attributes.map(method => ({
            content: method[2].content,
            height: method[2].height,
            width: method[2].width,
            fontSize: method[2].fontSize,
            fontStyle: method[2].fontStyle,
            fontFile: method[2].fontFile,
        })),
    });
    console.log("Parsed gadget:", gadget);
    return gadget;
}
