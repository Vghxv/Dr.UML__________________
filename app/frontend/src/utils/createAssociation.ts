import { AssociationProps } from "./Props";

// Association type constants
enum ASS_TYPE {
    ASS_TYPE_EXTENSION = 1,
    ASS_TYPE_IMPLEMENTATION = 2,
    ASS_TYPE_COMPOSITION = 4,
    ASS_TYPE_DEPENDENCY = 8,
}
class AssociationElement {
    public assProps: AssociationProps;

    constructor(props: AssociationProps, margin: number) {
        this.assProps = props;
    }

    draw(ctx: CanvasRenderingContext2D, margin: number, lineWidth: number) {
        // 根據 assType 決定畫法
        switch (this.assProps.assType) {
            case ASS_TYPE.ASS_TYPE_DEPENDENCY:
                this.drawLine(ctx, margin, lineWidth, true); // 虛線
                this.drawArrow(ctx, margin, lineWidth, false); // 普通箭頭
                break;
            case ASS_TYPE.ASS_TYPE_COMPOSITION:
                this.drawLine(ctx, margin, lineWidth, false); // 實線
                this.drawDiamond(ctx, margin, lineWidth, true); // 實心菱形
                this.drawArrow(ctx, margin, lineWidth, false); // 普通箭頭
                break;
            case ASS_TYPE.ASS_TYPE_EXTENSION:
                this.drawLine(ctx, margin, lineWidth, false); // 實線
                this.drawArrow(ctx, margin, lineWidth, true); // 空心三角
                break;
            case ASS_TYPE.ASS_TYPE_IMPLEMENTATION:
                this.drawLine(ctx, margin, lineWidth, true); // 虛線
                this.drawArrow(ctx, margin, lineWidth, true); // 空心三角
                break;
            default:
                // fallback: normal association
                if (this.assProps.deltaX === 0 && this.assProps.deltaY === 0) {
                    this.drawNormalAss(ctx, margin, lineWidth);
                } else {
                    this.drawSelfAss(ctx, margin, lineWidth);
                }
                this.drawArrow(ctx, margin, lineWidth, false);
        }
        // Draw label text(s) if present
        this.drawText(ctx, margin);
    }

    drawLine(ctx: CanvasRenderingContext2D, margin: number, lineWidth: number, dashed: boolean) {
        if (this.assProps.deltaX === 0 && this.assProps.deltaY === 0) {
            ctx.save();
            ctx.beginPath();
            if (dashed) ctx.setLineDash([8, 6]);
            ctx.moveTo(this.assProps.startX, this.assProps.startY);
            ctx.lineTo(this.assProps.endX, this.assProps.endY);
            ctx.strokeStyle = "black";
            ctx.lineWidth = lineWidth;
            ctx.stroke();
            ctx.setLineDash([]);
            ctx.restore();
        } else {
            // Self association
            ctx.save();
            ctx.beginPath();
            if (dashed) ctx.setLineDash([8, 6]);
            const { startX, startY, endX, endY, deltaX, deltaY } = this.assProps;
            ctx.moveTo(startX, startY);
            ctx.lineTo(startX + deltaX, startY + deltaY);
            ctx.lineTo(endX + deltaX, endY + deltaY);
            ctx.lineTo(endX, endY);
            ctx.strokeStyle = "black";
            ctx.lineWidth = lineWidth;
            ctx.stroke();
            ctx.setLineDash([]);
            ctx.restore();
        }
    }

    drawDiamond(ctx: CanvasRenderingContext2D, margin: number, lineWidth: number, filled: boolean) {
        // 菱形在起點
        const { startX, startY, endX, endY } = this.assProps;
        const dx = endX - startX;
        const dy = endY - startY;
        const len = Math.sqrt(dx * dx + dy * dy);
        if (len === 0) return;
        const unitX = dx / len;
        const unitY = dy / len;
        const size = 18;
        // 四個點
        const p0 = { x: startX, y: startY };
        const p1 = { x: startX + unitX * size / 2 - unitY * size / 3, y: startY + unitY * size / 2 + unitX * size / 3 };
        const p2 = { x: startX + unitX * size, y: startY + unitY * size };
        const p3 = { x: startX + unitX * size / 2 + unitY * size / 3, y: startY + unitY * size / 2 - unitX * size / 3 };

        ctx.save();
        ctx.beginPath();
        ctx.moveTo(p0.x, p0.y);
        ctx.lineTo(p1.x, p1.y);
        ctx.lineTo(p2.x, p2.y);
        ctx.lineTo(p3.x, p3.y);
        ctx.closePath();
        ctx.lineWidth = lineWidth;
        ctx.strokeStyle = "black";
        ctx.stroke();
        if (filled) {
            ctx.fillStyle = "black";
            ctx.fill();
        }
        ctx.restore();
    }

    drawArrow(ctx: CanvasRenderingContext2D, margin: number, lineWidth: number, hollow: boolean) {
        let fromX, fromY, toX, toY;
        if (this.assProps.deltaX === 0 && this.assProps.deltaY === 0) {
            fromX = this.assProps.startX;
            fromY = this.assProps.startY;
            toX = this.assProps.endX;
            toY = this.assProps.endY;
        } else {
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
        const arrowSize = 16;
        // 箭頭三角形
        const baseX = toX - unitX * arrowSize;
        const baseY = toY - unitY * arrowSize;
        const leftX = baseX - unitY * arrowSize / 2;
        const leftY = baseY + unitX * arrowSize / 2;
        const rightX = baseX + unitY * arrowSize / 2;
        const rightY = baseY - unitX * arrowSize / 2;

        ctx.save();
        ctx.beginPath();
        ctx.moveTo(toX, toY);
        ctx.lineTo(leftX, leftY);
        ctx.lineTo(rightX, rightY);
        ctx.closePath();
        ctx.lineWidth = lineWidth;
        ctx.strokeStyle = "black";
        if (hollow) {
            ctx.fillStyle = "white";
            ctx.fill();
            ctx.stroke();
        } else {
            ctx.fillStyle = "black";
            ctx.fill();
            ctx.stroke();
        }
        ctx.restore();
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
        const { startX, startY, endX, endY, deltaX, deltaY } = this.assProps;
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