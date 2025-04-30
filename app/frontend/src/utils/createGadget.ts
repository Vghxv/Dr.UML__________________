import { shapes, dia } from '@joint/core';

export interface GadgetOptions {
    point: { x: number; y: number };
    type: 'Class';  // 允許的 gadget 類型
    layer: number;  // 元件的層級
    size?: { width: number; height: number };  // 預設大小
    color?: string;  // 背景顏色
    outlineColor?: string;  // 邊框顏色
    name?: string;  // 顯示名稱
    attributesText?: string;
    methodsText?: string;
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
            y: 30,
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
            y: 75,
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
    size = { width: 150, height: 100 },
    color = '#FFFFFF',
    outlineColor = '#000000',
    name = '',
    attributesText = 'id: Int\nname: String',
    methodsText = '+getId(): Int\n+getName(): String',
}: GadgetOptions): dia.Element {

    switch (type) {
        case 'Class': {
            return new UMLClass({
                size: { width: size.width || 200, height: size.height || 150 },
                position: point,
                z: layer,
                attrs: {
                    header: { fill: '#3498DB', stroke: outlineColor },
                    headerLabel: { text: name || 'MyClass' },
                    attributesLabel: { text: attributesText },
                    methodsLabel: { text: methodsText },
                },
            });
        }
        default:
            throw new Error(`Unknown gadget type: ${type}`);
    }
}

// ---- 新增：將後端 JSON 字串轉為 GadgetOptions 並呼叫 createGadget ----

export interface BackendAttribute {
    content: string;
    height: number;
    width: number;
    fontSize: number;
    fontStyle: number;
    fontFile: string;
}

export interface BackendGadget {
    gadgetType: string;
    x: number;
    y: number;
    layer: number;
    height: number;
    width: number;
    color: number;
    attributes: BackendAttribute[][];
}

// 解析後端 Gadget JSON 字串並轉換格式
export function parseBackendGadget(json: string): dia.Element {
    const data: BackendGadget = JSON.parse(json);

    const colorHex = `#${data.color.toString(16).padStart(6, '0')}`;

    const attributesText = (data.attributes[0] || [])
        .map(attr => attr.content)
        .join('\n');

    const methodsText = (data.attributes[1] || [])
        .map(attr => attr.content)
        .join('\n');

    const options: GadgetOptions = {
        type: data.gadgetType as 'Class',
        point: { x: data.x, y: data.y },
        layer: data.layer,
        size: { width: data.width, height: data.height },
        color: colorHex,
        name: 'MyClass', // 可從 attr.content 或其他欄位拉取
        attributesText,
        methodsText
    };

    return createGadget(options);
}
