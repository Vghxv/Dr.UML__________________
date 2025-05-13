import { GadgetProps } from "../utils/Props";


class ClassElement {
    public op: GadgetProps;
    public len: number;
    public headerHeight: number;
    public attributesHeight: number;

    constructor(options: GadgetProps, margin: number) {
        this.op = options;
        this.len = options.attributes.length;

        const headerLen = options.attributes?.[0]?.length || 0;
        const attributesLen = options.attributes.length > 1 ? options.attributes[1]?.length || 0 : 0;
        const methodsLen = options.attributes.length > 2 ? options.attributes[2]?.length || 0 : 0;
        console.log(`Header: ${headerLen}, Attributes: ${attributesLen}, Methods: ${methodsLen}`);

        const calculateSectionHeight = (sectionIndex: number, sectionLen: number): number => {
            let height = 0;

            if (Array.isArray(options.attributes[sectionIndex])) {
                options.attributes[sectionIndex].forEach((attr: any) => {
                    if (attr && typeof attr.height === "number" && !isNaN(attr.height)) {
                        height += attr.height;
                    }
                });
            }

            return height + margin * (sectionLen + 1);
        };

        this.headerHeight = calculateSectionHeight(0, headerLen);
        this.attributesHeight = calculateSectionHeight(1, attributesLen);
    }

    draw(ctx: CanvasRenderingContext2D, margin: number, lineWidth: number) {
        ctx.fillStyle = this.op.color;
        ctx.fillRect(this.op.x, this.op.y, this.op.width, this.headerHeight);
        ctx.fillStyle = "white";
        ctx.fillRect(this.op.x, this.op.y + this.headerHeight, this.op.width, this.op.height - this.headerHeight);
        ctx.fillStyle = "black";
        ctx.strokeRect(this.op.x, this.op.y, this.op.width, this.op.height);
        ctx.beginPath();

        ctx.moveTo(this.op.x, this.op.y + this.headerHeight);
        ctx.lineTo(this.op.x + this.op.width, this.op.y + this.headerHeight);

        ctx.moveTo(this.op.x, this.op.y + this.headerHeight + this.attributesHeight);
        ctx.lineTo(this.op.x + this.op.width, this.op.y + this.headerHeight + this.attributesHeight);

        ctx.strokeStyle = "black";
        ctx.lineWidth = lineWidth;
        ctx.stroke();
        // ctx.font = "12px Georgia";
        const drawText = (sectionIndex: number, yOffset: number) => {
            yOffset += margin;
            if (Array.isArray(this.op.attributes[sectionIndex])) {
                this.op.attributes[sectionIndex].forEach((attr: any) => {
                    if (attr && typeof attr.content === "string") {
                        yOffset += attr.height / 2;
                        ctx.font = `${attr.fontSize}px ${attr.fontFile}`;
                        ctx.fillText(attr.content, this.op.x + margin, yOffset);
                        yOffset += attr.height / 2 + margin;
                    }
                });
            }
        }
        drawText(0, this.op.y);
        drawText(1, this.op.y + this.headerHeight);
        drawText(2, this.op.y + this.headerHeight + this.attributesHeight);
    }
}

export function createGadget(type: string, config: GadgetProps, margin: number) {
    switch (type) {
        case "Class": {
            return new ClassElement(config, margin);
        }
        default:
            throw new Error(`Unknown gadget type: ${type}`);
    }
}
