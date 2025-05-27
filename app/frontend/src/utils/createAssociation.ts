import {AssociationProps} from "./Props";

class AssociationElement {
    public assProps: AssociationProps;

    constructor(props: AssociationProps, margin: number) {
        this.assProps = props;
    }

    draw(ctx: CanvasRenderingContext2D, margin: number, lineWidth: number) {

        if (this.assProps.deltaX === 0 && this.assProps.deltaY === 0) {
            this.drawNormalAss(ctx, margin, lineWidth);
        } else {
            this.drawSelfAss(ctx, margin, lineWidth);
        }

        this.drawArrow(ctx, margin, lineWidth);

        // Draw label text(s) if present
        this.drawText(ctx, margin);
    }

    drawText(ctx: CanvasRenderingContext2D, margin: number) {
        if (Array.isArray(this.assProps.attributes)) {
            this.assProps.attributes.forEach(attr => {
                if (attr) {
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
        const {startX, startY, endX, endY, deltaX, deltaY} = this.assProps;

        const p0x = startX, p0y = startY;
        const p1x = startX + deltaX, p1y = startY + deltaY;
        const p2x = endX + deltaX, p2y = endY + deltaY;
        const p3x = endX, p3y = endY;

        ctx.beginPath();
        ctx.moveTo(p0x, p0y);
        ctx.lineTo(p1x, p1y);
        ctx.lineTo(p2x, p2y);
        ctx.lineTo(p3x, p3y);
        ctx.strokeStyle = "black";
        ctx.lineWidth = lineWidth;
        ctx.stroke();
    }

    drawArrow(ctx: CanvasRenderingContext2D, margin: number, lineWidth: number) {
        let fromX, fromY, toX, toY;

        if (this.assProps.deltaX === 0 && this.assProps.deltaY === 0) {
            // Normal association: arrow at end of straight line
            fromX = this.assProps.startX;
            fromY = this.assProps.startY;
            toX = this.assProps.endX;
            toY = this.assProps.endY;
        } else {
            // Self-association: arrow at end of last segment (from end+delta to end)
            fromX = this.assProps.endX + this.assProps.deltaX;
            fromY = this.assProps.endY + this.assProps.deltaY;
            toX = this.assProps.endX;
            toY = this.assProps.endY;
        }

        const dx = toX - fromX;
        const dy = toY - fromY;
        const len = Math.sqrt(dx * dx + dy * dy);
        if (len === 0) return;

        const unitX = dx / len;
        const unitY = dy / len;
        const arrowSize = 10;

        // Arrow tip at (toX, toY)
        // Two base points of the arrowhead
        const baseX = toX - unitX * arrowSize;
        const baseY = toY - unitY * arrowSize;

        ctx.beginPath();
        ctx.moveTo(toX, toY);
        ctx.lineTo(baseX - unitY * 5, baseY + unitX * 5);
        ctx.lineTo(baseX + unitY * 5, baseY - unitX * 5);
        ctx.closePath();
        ctx.fillStyle = "black";
        ctx.fill();
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