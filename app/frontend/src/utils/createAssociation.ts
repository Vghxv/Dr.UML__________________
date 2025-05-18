import { AssociationProps } from "./Props";

class AssociationElement {
    public assProps: AssociationProps;

    constructor(props: AssociationProps, margin: number) {
        this.assProps = props;
    }

    draw(ctx: CanvasRenderingContext2D, margin: number, lineWidth: number) {

        if (this.assProps.deltaX === 0 && this.assProps.deltaY === 0) {
            this.drawNormalAss(ctx, margin, lineWidth);
        }
        else{
            this.drawSelfAss(ctx, margin, lineWidth);
        }

        const drawArrow = () => {
            const dx = this.assProps.deltaX;
            const dy = this.assProps.deltaY;
            const len = Math.sqrt(dx * dx + dy * dy);
            const unitX = dx / len;
            const unitY = dy / len;
            const arrowSize = 10;

            const arrowX = this.assProps.endX - unitX * arrowSize;
            const arrowY = this.assProps.endY - unitY * arrowSize;

            ctx.beginPath();
            ctx.moveTo(this.assProps.endX, this.assProps.endY);
            ctx.lineTo(arrowX - unitY * 5, arrowY + unitX * 5);
            ctx.lineTo(arrowX + unitY * 5, arrowY - unitX * 5);
            ctx.closePath();
            ctx.fillStyle = "black";
            ctx.fill();
        };

        drawArrow();

        // Draw label text(s) if present
        this.drawText(ctx, margin);
    }

    drawText(ctx: CanvasRenderingContext2D, margin: number) {
        // attributes: { content, fontSize, fontStyle, fontFile, ratio }[]
        if (Array.isArray(this.assProps.attributes)) {
            this.assProps.attributes.forEach(attr => {
                if (attr && typeof attr.content === "string") {
                    // 計算線段上顯示位置
                    const x = this.assProps.startX + (this.assProps.endX - this.assProps.startX) * (attr.ratio ?? 0.5);
                    const y = this.assProps.startY + (this.assProps.endY - this.assProps.startY) * (attr.ratio ?? 0.5);
                    ctx.font = `${attr.fontSize || 12}px ${attr.fontFile || "Arial"}`;
                    ctx.fillStyle = "black";
                    ctx.fillText(attr.content, x + margin, y);
                }
            });
        }
    }

    drawNormalAss(ctx: CanvasRenderingContext2D, margin: number, lineWidth: number) {
        ctx.beginPath();
        ctx.moveTo(this.assProps.startX, this.assProps.startY);
        ctx.lineTo(this.assProps.endX, this.assProps.endY);
        ctx.strokeStyle = "black";
        ctx.lineWidth = lineWidth;
        ctx.stroke();
    }

    drawSelfAss(ctx: CanvasRenderingContext2D, margin: number, lineWidth: number) {
        ctx.beginPath();
        ctx.moveTo(this.assProps.startX, this.assProps.startY);
        ctx.lineTo(this.assProps.startX + this.assProps.deltaX, this.assProps.startY + this.assProps.deltaY);
        ctx.lineTo(this.assProps.deltaX, this.assProps.deltaY);
        ctx.strokeStyle = "black";
        ctx.lineWidth = lineWidth;
        ctx.stroke();
    }
}



export function createAss(type: string, config: AssociationProps, margin: number) {
    switch (type) {
        case "Association": {
            return new AssociationElement(config, margin);
        }
        default:
            throw new Error(`Unknown association type: ${type}`);
    }
}