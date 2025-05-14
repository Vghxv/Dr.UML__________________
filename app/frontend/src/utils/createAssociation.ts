import { AssociationProps } from "./Props";

class AssociationElement {
    public op: AssociationProps;

    constructor(options: AssociationProps) {
        this.op = options;
    }

    draw(ctx: CanvasRenderingContext2D, lineWidth: number) {
        ctx.save();
        ctx.lineWidth = lineWidth;
        ctx.strokeStyle = "black";

        // 處理不同的 assType
        switch (this.op.assType) {
            case 0: // 普通實線
                ctx.setLineDash([]);
                break;
            case 1: // 虛線
                ctx.setLineDash([5, 5]);
                break;
            case 2: // 繼承（箭頭）
                ctx.setLineDash([]);
                this.drawArrow(ctx);
                break;
            // 可擴充其他關係如聚合、依賴等
            default:
                ctx.setLineDash([]);
        }

        // 畫主線
        ctx.beginPath();
        ctx.moveTo(this.op.startX, this.op.startY);
        ctx.lineTo(this.op.endX, this.op.endY);
        ctx.stroke();
        ctx.setLineDash([]);

        // 畫屬性文字
        this.op.attributes.forEach(attr => {
            const midX = this.op.startX + (this.op.endX - this.op.startX) * attr.ratio;
            const midY = this.op.startY + (this.op.endY - this.op.startY) * attr.ratio;

            ctx.font = `${attr.fontSize}px ${attr.fontFile}`;
            ctx.fillStyle = "black";
            ctx.fillText(attr.content, midX + 5, midY - 5); // 偏移避免覆蓋線條
        });

        ctx.restore();
    }

    private drawArrow(ctx: CanvasRenderingContext2D) {
        const { startX, startY, endX, endY } = this.op;
        const headlen = 10;

        const angle = Math.atan2(endY - startY, endX - startX);

        const arrowX = endX;
        const arrowY = endY;

        const x1 = arrowX - headlen * Math.cos(angle - Math.PI / 6);
        const y1 = arrowY - headlen * Math.sin(angle - Math.PI / 6);
        const x2 = arrowX - headlen * Math.cos(angle + Math.PI / 6);
        const y2 = arrowY - headlen * Math.sin(angle + Math.PI / 6);

        ctx.beginPath();
        ctx.moveTo(x1, y1);
        ctx.lineTo(arrowX, arrowY);
        ctx.lineTo(x2, y2);
        ctx.stroke();
    }
}

export function createAssociation(config: AssociationProps) {
    return new AssociationElement(config);
}
