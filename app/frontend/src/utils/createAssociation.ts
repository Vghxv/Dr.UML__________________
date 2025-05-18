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

        // draw selected association
        // if (isSelected) {
        //     ctx.setLineDash([5, 3]);
        //     ctx.strokeStyle = "#FFA500";
        //     ctx.lineWidth = lineWidth * 2;
        //     ctx.stroke();
        //     ctx.setLineDash([]);
        //     ctx.strokeStyle = "black";
        //     ctx.lineWidth = lineWidth;
        // }

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
    }

    drawNormalAss(ctx: CanvasRenderingContext2D, margin: number, lineWidth: number) {
        ctx.beginPath();
        ctx.moveTo(this.assProps.startX, this.assProps.startY);
        ctx.lineTo(this.assProps.endX, this.assProps.endY);
        ctx.fillStyle = "black";
        ctx.fill();
    }

    drawSelfAss(ctx: CanvasRenderingContext2D, margin: number, lineWidth: number) {
        ctx.beginPath();
        ctx.moveTo(this.assProps.startX, this.assProps.startY);
        ctx.lineTo(this.assProps.startX + this.assProps.deltaX, this.assProps.startY + this.assProps.deltaY);
        ctx.lineTo(this.assProps.deltaX, this.assProps.deltaY);
        ctx.fillStyle = "black";
        ctx.fill();
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