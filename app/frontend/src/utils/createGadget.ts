import {GadgetProps} from "./Props";


class ClassElement {
    public gadgetProps: GadgetProps;
    public len: number;
    public headerHeight: number;
    public attributesHeight: number;
    public colorHexStr: string;

    constructor(props: GadgetProps, margin: number) {
        this.gadgetProps = props;
        this.len = props.attributes.length;

        const headerLen = props.attributes?.[0]?.length || 0;
        const attributesLen = props.attributes.length > 1 ? props.attributes[1]?.length || 0 : 0;

        const calculateSectionHeight = (sectionIndex: number, sectionLen: number): number => {
            let height = 0;

            if (Array.isArray(props.attributes[sectionIndex])) {
                props.attributes[sectionIndex].forEach((attr: any) => {
                    if (attr && typeof attr.height === "number" && !isNaN(attr.height)) {
                        height += attr.height;
                    }
                });
            }

            return height + margin * (sectionLen + 1);
        };

        this.headerHeight = calculateSectionHeight(0, headerLen);
        this.attributesHeight = calculateSectionHeight(1, attributesLen);
        this.colorHexStr = props.color;
    }

    draw(ctx: CanvasRenderingContext2D, margin: number, lineWidth: number) {
        ctx.beginPath();
        ctx.fillStyle = this.colorHexStr;
        ctx.fillRect(this.gadgetProps.x, this.gadgetProps.y, this.gadgetProps.width, this.headerHeight);
        ctx.fillStyle = "white";
        ctx.fillRect(this.gadgetProps.x, this.gadgetProps.y + this.headerHeight, this.gadgetProps.width, this.gadgetProps.height - this.headerHeight);
        ctx.fillStyle = "black";
        ctx.strokeRect(this.gadgetProps.x, this.gadgetProps.y, this.gadgetProps.width, this.gadgetProps.height);

        ctx.moveTo(this.gadgetProps.x, this.gadgetProps.y + this.headerHeight);
        ctx.lineTo(this.gadgetProps.x + this.gadgetProps.width, this.gadgetProps.y + this.headerHeight);

        ctx.moveTo(this.gadgetProps.x, this.gadgetProps.y + this.headerHeight + this.attributesHeight);
        ctx.lineTo(this.gadgetProps.x + this.gadgetProps.width, this.gadgetProps.y + this.headerHeight + this.attributesHeight);

        ctx.strokeStyle = "black";
        ctx.lineWidth = lineWidth;
        ctx.stroke();
        const drawText = (sectionIndex: number, yOffset: number) => {
            yOffset += margin;
            if (Array.isArray(this.gadgetProps.attributes[sectionIndex])) {
                this.gadgetProps.attributes[sectionIndex].forEach((attr: any) => {
                    if (attr && typeof attr.content === "string") {
                        yOffset += Math.round(attr.height);
                        ctx.font = `${attr.fontSize}px ${attr.fontFile}`;
                        ctx.textBaseline = "bottom"
                        ctx.fillText(attr.content, this.gadgetProps.x + margin, yOffset);
                        yOffset += margin;
                    }
                });
            }
        }
        drawText(0, this.gadgetProps.y);
        drawText(1, this.gadgetProps.y + this.headerHeight);
        drawText(2, this.gadgetProps.y + this.headerHeight + this.attributesHeight);
        if (this.gadgetProps.isSelected) {
            ctx.setLineDash([5, 3]);
            ctx.strokeStyle = '#FFA500';
            ctx.lineWidth = lineWidth * 2;
            ctx.strokeRect(this.gadgetProps.x, this.gadgetProps.y, this.gadgetProps.width, this.gadgetProps.height);
            ctx.setLineDash([]);
            ctx.strokeStyle = 'black';
            ctx.lineWidth = lineWidth;
        }
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
