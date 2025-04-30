import { shapes, dia } from '@joint/core';  // 引入 JointJS 库

export interface GadgetOptions {
    point: { x: number; y: number };
    type: 'Class' | 'Rect' | 'Circle';  // 允許的 gadget 類型
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
            refWidth: '100%',
            height: '30%',
            fill: '#2ECC71',
            stroke: '#000000',
            strokeWidth: 2,
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
            refY: '30%',
            refWidth: '100%',
            height: '35%',
            fill: '#ECF0F1',
            stroke: '#000000',
            strokeWidth: 1,
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
            refY: '65%',
            refWidth: '100%',
            height: '35%',
            fill: '#ECF0F1',
            stroke: '#000000',
            strokeWidth: 1,
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
    markup: [{
        tagName: 'rect',
        selector: 'header',
    }, {
        tagName: 'text',
        selector: 'headerLabel',
    }, {
        tagName: 'rect',
        selector: 'attributes',
    }, {
        tagName: 'text',
        selector: 'attributesLabel',
    }, {
        tagName: 'rect',
        selector: 'methods',
    }, {
        tagName: 'text',
        selector: 'methodsLabel',
    }]
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
        case 'Rect': {
            // 創建矩形類型的 gadget
            return new shapes.standard.Rectangle({
                position: point,
                size,
                attrs: {
                    body: {
                        fill: color,
                        stroke: outlineColor,
                        strokeWidth: 2,
                    },
                    label: {
                        text: name,
                        fill: '#000000',
                    },
                },
                z: layer,  // 設定元件的層級
            });
        }
        case 'Circle': {
            // 創建圓形類型的 gadget，根據 size 計算出最小的寬度和高度以保持圓形
            const minSize = Math.min(size.width, size.height);
            return new shapes.standard.Circle({
                position: point,
                size: { width: minSize, height: minSize },  // 確保圓形的寬高相等
                attrs: {
                    body: {
                        fill: color,
                        stroke: outlineColor,
                        strokeWidth: 2,
                    },
                    label: {
                        text: name,
                        fill: '#000000',
                    },
                },
                z: layer,  // 設定元件的層級
            });
        }
        default:
            throw new Error(`Unknown gadget type: ${type}`);  // 如果是未知的類型，拋出錯誤
    }
}
