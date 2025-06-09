import { AssociationProps } from "./Props";
import { component, attribute } from "../../wailsjs/go/models";

class AssociationElement {
    public assProps: AssociationProps;
    private isSelfAssociation: boolean;
    constructor(props: AssociationProps, margin: number) {
        this.assProps = props;
        this.isSelfAssociation = props.deltaX !== 0 || props.deltaY !== 0;
    }

    draw(ctx: CanvasRenderingContext2D, margin: number, lineWidth: number) {
        switch (this.assProps.assType) {
            case component.AssociationType.Dependency:
                this.drawLine(ctx, lineWidth, true);
                this.drawArrow(ctx, lineWidth, false);
                break;
            case component.AssociationType.Composition:
                this.drawLine(ctx, lineWidth, false);
                this.drawDiamond(ctx, lineWidth, true);
                this.drawArrow(ctx, lineWidth, false);
                break;
            case component.AssociationType.Extension:
                this.drawLine(ctx, lineWidth, false);
                this.drawArrow(ctx, lineWidth, true);
                break;
            case component.AssociationType.Implementation:
                this.drawLine(ctx, lineWidth, true);
                this.drawArrow(ctx, lineWidth, true);
                break;
            default:
                if (this.isSelfAssociation) {
                    this.drawSelfAss(ctx, margin, lineWidth);
                } else {
                    this.drawNormalAss(ctx, margin, lineWidth);
                }
                this.drawArrow(ctx, lineWidth, false);
        }
        this.drawText(ctx, margin);

    }

    drawLine(ctx: CanvasRenderingContext2D, lineWidth: number, dashed: boolean) {
        if (this.assProps.deltaX === 0 && this.assProps.deltaY === 0) {
            ctx.save();
            ctx.beginPath();
            if (dashed){
                ctx.setLineDash([8, 6]);
            } 
            ctx.moveTo(this.assProps.startX, this.assProps.startY);
            ctx.lineTo(this.assProps.endX, this.assProps.endY);
            ctx.strokeStyle = "black";
            ctx.lineWidth = lineWidth;
            ctx.stroke();
            ctx.setLineDash([]);
            // isSelected: draw orange highlight
            if (this.assProps.isSelected) {
                ctx.beginPath();
                ctx.setLineDash([5, 3]);
                ctx.moveTo(this.assProps.startX, this.assProps.startY);
                ctx.lineTo(this.assProps.endX, this.assProps.endY);
                ctx.strokeStyle = "#FFA500";
                ctx.lineWidth = lineWidth * 2;
                ctx.stroke();
                ctx.setLineDash([]);
            }
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
            // isSelected: draw orange highlight
            if (this.assProps.isSelected) {
                ctx.beginPath();
                ctx.setLineDash([5, 3]);
                ctx.moveTo(startX, startY);
                ctx.lineTo(startX + deltaX, startY + deltaY);
                ctx.lineTo(endX + deltaX, endY + deltaY);
                ctx.lineTo(endX, endY);
                ctx.strokeStyle = "#FFA500";
                ctx.lineWidth = lineWidth * 2;
                ctx.stroke();
                ctx.setLineDash([]);
            }
            ctx.restore();
        }
    }

    drawDiamond(ctx: CanvasRenderingContext2D, lineWidth: number, filled: boolean) {
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

    drawArrow(ctx: CanvasRenderingContext2D, lineWidth: number, hollow: boolean) {
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
        // isSelected: draw orange border for arrow
        if (this.assProps.isSelected) {
            ctx.beginPath();
            ctx.moveTo(toX, toY);
            ctx.lineTo(leftX, leftY);
            ctx.lineTo(rightX, rightY);
            ctx.closePath();
            ctx.setLineDash([5, 3]);
            ctx.lineWidth = lineWidth * 2;
            ctx.strokeStyle = "#FFA500";
            ctx.stroke();
            ctx.setLineDash([]);
        }
        ctx.restore();
    }

    drawText(ctx: CanvasRenderingContext2D, margin: number) {
        if (Array.isArray(this.assProps.attributes)) {
            this.assProps.attributes.forEach(attr => {
                if (attr && typeof attr.content === "string") {
                    const x = this.assProps.startX + (this.assProps.endX - this.assProps.startX) * (attr.ratio ?? 0.5);
                    const y = this.assProps.startY + (this.assProps.endY - this.assProps.startY) * (attr.ratio ?? 0.5);
                    const boldString = (attr.fontStyle & attribute.Textstyle.Bold) !== 0 ? "bold " : "";
                    const italicString = (attr.fontStyle & attribute.Textstyle.Italic) !== 0 ? "italic " : "";
                    const isUnderline = (attr.fontStyle & attribute.Textstyle.Underline) !== 0;
                    ctx.font = `${boldString}${italicString}${attr.fontSize}px ${attr.fontFile}`;
                    ctx.textBaseline = "middle"
                    ctx.fillText(attr.content, x + margin, y);
                    if (isUnderline) {
                        const underlineHeight = 2;
                        ctx.fillRect(x + margin, y + Math.round(attr.height * 0.6), ctx.measureText(attr.content).width, underlineHeight);
                    }
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