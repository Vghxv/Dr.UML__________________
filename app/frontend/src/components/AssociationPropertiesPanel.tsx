import React, { useEffect, useRef, useState } from "react";
import { AssociationProps } from "../utils/Props";
import { attribute } from "../../wailsjs/go/models";

// Define type for font files
type FontFile = {
    default: string;
};

let fontFiles: Record<string, FontFile>;
fontFiles = import.meta.glob<FontFile>('../assets/fonts/*.woff2', { eager: true });

const getFontOptions = () => {
    return Object.keys(fontFiles).map(path => {
        const filename = path.split(/[/\\]/).pop() || '';
        return filename.replace('.woff2', '');
    });
};

interface AssociationPropertiesPanelProps {
    selectedAssociation: AssociationProps;
    updateAssociationProperty: (property: string, value: any) => void;
}

const AssociationPropertiesPanel: React.FC<AssociationPropertiesPanelProps> = ({
    selectedAssociation,
    updateAssociationProperty
}) => {
    const [focusedInput, setFocusedInput] = useState<string | null>(null);
    const inputRefs = useRef<{ [key: string]: HTMLInputElement | HTMLSelectElement | null }>({});

    useEffect(() => {
        if (focusedInput && inputRefs.current[focusedInput]) {
            inputRefs.current[focusedInput]?.focus();
        }
    }, [selectedAssociation, focusedInput]);

    return (
        <div className="absolute right-0 top-0 w-[300px] h-full bg-gray-100 p-5 shadow-md overflow-y-auto">
            <h3 className="text-xl font-semibold text-gray-800 mb-4">Association Properties</h3>
            {/* layer */}
            <div className="mb-4">
                <label className="block mb-1 text-sm font-medium text-gray-700">Layer:</label>
                <input
                    type="number"
                    value={selectedAssociation.layer}
                    ref={el => inputRefs.current['layer'] = el}
                    onFocus={() => setFocusedInput('layer')}
                    onChange={e => updateAssociationProperty('layer', parseInt(e.target.value))}
                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                />
            </div>
            {/* startX, startY */}
            <div className="mb-4">
                <label className="block mb-1 text-sm font-medium text-gray-700">Start X:</label>
                <input
                    type="number"
                    value={selectedAssociation.startX}
                    ref={el => inputRefs.current['startX'] = el}
                    onFocus={() => setFocusedInput('startX')}
                    onChange={e => updateAssociationProperty('startX', parseInt(e.target.value))}
                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                />
            </div>
            <div className="mb-4">
                <label className="block mb-1 text-sm font-medium text-gray-700">Start Y:</label>
                <input
                    type="number"
                    value={selectedAssociation.startY}
                    ref={el => inputRefs.current['startY'] = el}
                    onFocus={() => setFocusedInput('startY')}
                    onChange={e => updateAssociationProperty('startY', parseInt(e.target.value))}
                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                />
            </div>
            {/* endX, endY */}
            <div className="mb-4">
                <label className="block mb-1 text-sm font-medium text-gray-700">End X:</label>
                <input
                    type="number"
                    value={selectedAssociation.endX}
                    ref={el => inputRefs.current['endX'] = el}
                    onFocus={() => setFocusedInput('endX')}
                    onChange={e => updateAssociationProperty('endX', parseInt(e.target.value))}
                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                />
            </div>
            <div className="mb-4">
                <label className="block mb-1 text-sm font-medium text-gray-700">End Y:</label>
                <input
                    type="number"
                    value={selectedAssociation.endY}
                    ref={el => inputRefs.current['endY'] = el}
                    onFocus={() => setFocusedInput('endY')}
                    onChange={e => updateAssociationProperty('endY', parseInt(e.target.value))}
                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                />
            </div>
            {/* deltaX, deltaY */}
            <div className="mb-4">
                <label className="block mb-1 text-sm font-medium text-gray-700">Delta X:</label>
                <input
                    type="number"
                    value={selectedAssociation.deltaX}
                    ref={el => inputRefs.current['deltaX'] = el}
                    onFocus={() => setFocusedInput('deltaX')}
                    onChange={e => updateAssociationProperty('deltaX', parseInt(e.target.value))}
                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                />
            </div>
            <div className="mb-4">
                <label className="block mb-1 text-sm font-medium text-gray-700">Delta Y:</label>
                <input
                    type="number"
                    value={selectedAssociation.deltaY}
                    ref={el => inputRefs.current['deltaY'] = el}
                    onFocus={() => setFocusedInput('deltaY')}
                    onChange={e => updateAssociationProperty('deltaY', parseInt(e.target.value))}
                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                />
            </div>
            {/* attributes */}
            <div className="mb-5">
                <h4 className="text-md font-medium text-gray-700 mb-2">Labels</h4>
                {selectedAssociation.attributes.map((attr, attrIndex) => (
                    <div key={`attr-${attrIndex}`} className="mb-4 p-3 border border-gray-300 rounded-md bg-white">
                        <div className="mb-3">
                            <label className="block mb-1 text-sm font-medium text-gray-700">Content:</label>
                            <input
                                type="text"
                                value={attr.content}
                                ref={el => inputRefs.current[`attributes:${attrIndex}.content`] = el}
                                onFocus={() => setFocusedInput(`attributes:${attrIndex}.content`)}
                                onChange={e => updateAssociationProperty(`attributes:${attrIndex}.content`, e.target.value)}
                                className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                            />
                        </div>
                        <div className="mb-3">
                            <label className="block mb-1 text-sm font-medium text-gray-700">Font Size:</label>
                            <input
                                type="number"
                                value={attr.fontSize}
                                ref={el => inputRefs.current[`attributes:${attrIndex}.fontSize`] = el}
                                onFocus={() => setFocusedInput(`attributes:${attrIndex}.fontSize`)}
                                onChange={e => updateAssociationProperty(`attributes:${attrIndex}.fontSize`, parseInt(e.target.value))}
                                className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                            />
                        </div>
                        <div className="mb-3">
                            <label className="block mb-1 text-sm font-medium text-gray-700">Font Style:</label>
                            <div className="flex space-x-2">
                                <button
                                    type="button"
                                    onClick={() => {
                                        const isBold = (attr.fontStyle & attribute.Textstyle.Bold) !== 0;
                                        let newStyle = attr.fontStyle;
                                        if (isBold) {
                                            newStyle &= ~attribute.Textstyle.Bold;
                                        } else {
                                            newStyle |= attribute.Textstyle.Bold;
                                        }
                                        updateAssociationProperty(`attributes:${attrIndex}.fontStyleB`, newStyle);
                                    }}
                                    className={`px-3 py-2 border rounded-md ${(attr.fontStyle & attribute.Textstyle.Bold) !== 0
                                            ? 'bg-blue-500 text-white'
                                            : 'bg-white text-gray-700 border-gray-300'
                                        } hover:bg-blue-600 hover:text-white font-bold`}
                                >
                                    B
                                </button>
                                <button
                                    type="button"
                                    onClick={() => {
                                        const isItalic = (attr.fontStyle & attribute.Textstyle.Italic) !== 0;
                                        let newStyle = attr.fontStyle;
                                        if (isItalic) {
                                            newStyle &= ~attribute.Textstyle.Italic;
                                        } else {
                                            newStyle |= attribute.Textstyle.Italic;
                                        }
                                        updateAssociationProperty(`attributes:${attrIndex}.fontStyleI`, newStyle);
                                    }}
                                    className={`px-3 py-2 border rounded-md ${(attr.fontStyle & attribute.Textstyle.Italic) !== 0
                                            ? 'bg-blue-500 text-white'
                                            : 'bg-white text-gray-700 border-gray-300'
                                        } hover:bg-blue-600 hover:text-white italic`}
                                >
                                    I
                                </button>
                                <button
                                    type="button"
                                    onClick={() => {
                                        const isUnderline = (attr.fontStyle & attribute.Textstyle.Underline) !== 0;
                                        let newStyle = attr.fontStyle;
                                        if (isUnderline) {
                                            newStyle &= ~attribute.Textstyle.Underline;
                                        } else {
                                            newStyle |= attribute.Textstyle.Underline;
                                        }
                                        updateAssociationProperty(`attributes:${attrIndex}.fontStyleU`, newStyle);
                                    }}
                                    className={`px-3 py-2 border rounded-md ${(attr.fontStyle & attribute.Textstyle.Underline) !== 0
                                            ? 'bg-blue-500 text-white'
                                            : 'bg-white text-gray-700 border-gray-300'
                                        } hover:bg-blue-600 hover:text-white underline`}
                                >
                                    U
                                </button>
                            </div>
                        </div>
                        <div className="mb-3">
                            <label className="block mb-1 text-sm font-medium text-gray-700">Font File:</label>
                            <select
                                value={attr.fontFile}
                                ref={el => inputRefs.current[`attributes:${attrIndex}.fontFile`] = el}
                                onFocus={() => setFocusedInput(`attributes:${attrIndex}.fontFile`)}
                                onChange={e => updateAssociationProperty(`attributes:${attrIndex}.fontFile`, e.target.value)}
                                className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                            >
                                {getFontOptions().map((fontName) => (
                                    <option key={fontName} value={fontName}>{fontName}</option>
                                ))}
                            </select>
                        </div>
                        <div className="mb-3">
                            <label className="block mb-1 text-sm font-medium text-gray-700">Ratio (0~1):</label>
                            <input
                                type="number"
                                min={0}
                                max={1}
                                step={0.01}
                                value={attr.ratio}
                                ref={el => inputRefs.current[`attributes:${attrIndex}.ratio`] = el}
                                onFocus={() => setFocusedInput(`attributes:${attrIndex}.ratio`)}
                                onChange={e => updateAssociationProperty(`attributes:${attrIndex}.ratio`, parseFloat(e.target.value))}
                                className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                            />
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default AssociationPropertiesPanel;
