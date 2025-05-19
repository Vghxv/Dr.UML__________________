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

        this.drawArrow(ctx, margin, lineWidth);

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

    drawArrow(ctx: CanvasRenderingContext2D, margin: number, lineWidth: number){
        let dx, dy, ex, ey;
        if (this.assProps.deltaX === 0 && this.assProps.deltaY === 0) {
            // 一般連線
            dx = this.assProps.endX - this.assProps.startX;
            dy = this.assProps.endY - this.assProps.startY;
            ex = this.assProps.endX;
            ey = this.assProps.endY;
        } else {
            // self-association，箭頭要根據最後一段
            const sx = this.assProps.startX + this.assProps.deltaX;
            const sy = this.assProps.startY + this.assProps.deltaY;
            dx = this.assProps.endX - sx;
            dy = this.assProps.endY - sy;
            ex = this.assProps.endX;
            ey = this.assProps.endY;
        }
        const len = Math.sqrt(dx * dx + dy * dy);
        if (len === 0) return;

        const unitX = dx / len;
        const unitY = dy / len;
        const arrowSize = 10;

        const arrowX = ex - unitX * arrowSize;
        const arrowY = ey - unitY * arrowSize;

        ctx.beginPath();
        ctx.moveTo(ex, ey);
        ctx.lineTo(arrowX - unitY * 5, arrowY + unitX * 5);
        ctx.lineTo(arrowX + unitY * 5, arrowY - unitX * 5);
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