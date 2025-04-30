import { shapes, dia } from '@joint/core';

export interface GadgetOptions {
    point: { x: number; y: number };
    type: 'Class';  // 允許的 gadget 類型
    layer: number;  // 元件的層級
    size?: { width: number; height: number };  // 預設大小
    color?: string;  // 背景顏色
    outlineColor?: string;  // 邊框顏色
    name?: string;  // 顯示名稱
}

const UMLClass = dia.Element.define('uml.Class', {
    size: { width: 200, height: 120 },
    attrs: {
        header: {
            x: 0,
            y: 0,
            width: 200,
            height: 30,
            fill: '#2ECC71',
            stroke: '#000000',
        },
        headerLabel: {
            ref: 'header',
            refX: '50%',
            refY: '50%',
            textAnchor: 'middle',
            yAlignment: 'middle',
            text: 'Class Name',
            fill: '#FFFFFF',
            fontWeight: 'bold',
        },
        attributes: {
            x: 0,
            y: 30,  // 緊接在 header 底下
            width: 200,
            height: 45,
            fill: '#ECF0F1',
            stroke: '#000000',
        },
        attributesLabel: {
            ref: 'attributes',
            refX: 5,
            refY: 5,
            textAnchor: 'left',
            yAlignment: 'top',
            text: 'Attributes',
            fill: '#333333',
        },
        methods: {
            x: 0,
            y: 75,  // 緊接在 attributes 底下 (30 + 45)
            width: 200,
            height: 45,
            fill: '#ECF0F1',
            stroke: '#000000',
        },
        methodsLabel: {
            ref: 'methods',
            refX: 5,
            refY: 5,
            textAnchor: 'left',
            yAlignment: 'top',
            text: 'Methods',
            fill: '#333333',
        },
    }
}, {
    markup: [
        { tagName: 'rect', selector: 'header' },
        { tagName: 'text', selector: 'headerLabel' },
        { tagName: 'rect', selector: 'attributes' },
        { tagName: 'text', selector: 'attributesLabel' },
        { tagName: 'rect', selector: 'methods' },
        { tagName: 'text', selector: 'methodsLabel' },
    ]
});


export function createGadget({
    point,
    type,
    layer,
    size = { width: 150, height: 100 },  // 預設大小
    color = '#FFFFFF',  // 預設顏色
    outlineColor = '#000000',  // 預設邊框顏色
    name = '',  // 預設名稱
}: GadgetOptions): dia.Element {

    switch (type) {
        case 'Class': {
            
            return new UMLClass({
                size: { width: size.width || 200, height: size.height || 150 },
                position: point,
                z: layer,
                attrs: {
                    header: { fill: '#3498DB' },
                    headerLabel: { text: name || 'MyClass' },
                    attributesLabel: { text: 'id: Int\nname: String' },
                    methodsLabel: { text: '+getId(): Int\n+getName(): String' },
                },
            });
        }
        default:
            throw new Error(`Unknown gadget type: ${type}`);  // 如果是未知的類型，拋出錯誤
    }
}
