import { shapes, dia } from "@joint/core";

export interface GadgetOptions {
  point: { x: number; y: number };
  type: "Class"; // Allowed gadget type
  layer: number; // Layer of the gadget
  size?: { width: number; height: number }; // Default size
  color?: string; // Background color
  outlineColor?: string; // Border color
  name?: string; // Display name
  attributesText?: string;
  methodsText?: string;
}

class UMLClass extends shapes.standard.Rectangle {
constructor(options: GadgetOptions) {
    super({
        position: options.point,
        size: options.size || { width: 200, height: 120 },
        attrs: {
            header: {
                x: 0,
                y: 0,
                width: options.size?.width || 200,
                height: 30,
                fill: "#2ECC71",
                stroke: "#000000",
            },
            headerLabel: {
                ref: "header",
                refX: "50%",
                refY: "50%",
                textAnchor: "middle",
                yAlignment: "middle",
                text: options.name || "Class Name",
                fill: "#FFFFFF",
                fontWeight: "bold",
            },
            attributes: {
                x: 0,
                y: 30,
                width: options.size?.width || 200,
                height: (options.size?.height || 120) / 2,
                fill: "#ECF0F1",
                stroke: "#000000",
            },
            attributesLabel: {
                ref: "attributes",
                refX: 5,
                refY: 5,
                textAnchor: "left",
                yAlignment: "top",
                text: options.attributesText || "Attributes",
                fill: "#333333",
            },
            methods: {
                x: 0,
                y: 60,
                width: options.size?.width || 200,
                height: (options.size?.height || 120) / 2,
                fill: "#ECF0F1",
                stroke: "#000000",
            },
            methodsLabel: {
                ref: "methods",
                refX: 5,
                refY: 5,
                textAnchor: "left",
                yAlignment: "top",
                text: options.methodsText || "Methods",
                fill: "#333333",
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

export function createGadget(type: string, config: GadgetOptions): dia.Element {
  switch (type) {
    case "Class": {
      return new UMLClass(config);
    }
    default:
      throw new Error(`Unknown gadget type: ${type}`);
  }
}

export interface BackendGadget {
  gadgetType: string;
  x: number;
  y: number;
  layer: number;
  height: number;
  width: number;
  color: number;
    attributes: {
        content: string;
        height: number;
        width: number;
        fontSize: number;
        fontStyle: number;
        fontFile: string;
    }[][];
}

// Parse backend gadget JSON and convert it to a dia.Element
export function parseBackendGadget(gadgetData: BackendGadget): dia.Element {
    console.log("Parsing backend gadget data:", gadgetData);
    return createGadget("Class", {
      point: { x: gadgetData.x, y: gadgetData.y },
      type: "Class",
      layer: gadgetData.layer,
      size: { width: gadgetData.width, height: gadgetData.height },
      color: `#${gadgetData.color.toString(16).padStart(6, "0")}`,
      name: gadgetData.attributes[0]?.[0]?.content || "Class Name",
      attributesText: gadgetData.attributes[1]?.map(attr => attr.content).join("\n") || "",
      methodsText: gadgetData.attributes[2]?.map(attr => attr.content).join("\n") || "",
    });
}
