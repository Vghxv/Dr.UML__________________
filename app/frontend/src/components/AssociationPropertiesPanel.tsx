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
    addAttributeToAssociation: (ratio: number, content: string) => void;
}

const AssociationPropertiesPanel: React.FC<AssociationPropertiesPanelProps> = ({
    selectedAssociation,
    updateAssociationProperty,
    addAttributeToAssociation
}) => {
    // Local state to track pending changes for each input
    const [pendingChanges, setPendingChanges] = useState<Record<string, any>>({});

    // Clear pending changes when association selection changes
    useEffect(() => {
        setPendingChanges({});
    }, [selectedAssociation.endX, selectedAssociation.endY]);

    // Handle input changes locally without calling API
    const handleInputChange = (property: string, value: any) => {
        setPendingChanges(prev => ({
            ...prev,
            [property]: value
        }));
    };

    // Handle Enter key press to commit changes
    const handleKeyPress = (e: React.KeyboardEvent, property: string) => {
        if (e.key === 'Enter') {
            const value = pendingChanges[property];
            if (value !== undefined) {
                updateAssociationProperty(property, value);
                setPendingChanges(prev => {
                    const updated = { ...prev };
                    delete updated[property];
                    return updated;
                });
            }
        }
    };

    // Get current value (pending change or original value)
    const getValue = (property: string, originalValue: any) => {
        return pendingChanges[property] !== undefined ? pendingChanges[property] : originalValue;
    };

    return (
        <div className="absolute right-0 top-0 w-[300px] h-full bg-gray-100 p-5 shadow-md overflow-y-auto">
            <h3 className="text-xl font-semibold text-gray-800 mb-4">Association Properties</h3>
            {/* layer */}
            <div className="mb-4">
                <label className="block mb-1 text-sm font-medium text-gray-700">Layer:</label>
                <input
                    type="number"
                    value={getValue('layer', selectedAssociation.layer)}
                    onChange={(e) => handleInputChange('layer', parseInt(e.target.value))}
                    onKeyPress={(e) => handleKeyPress(e, 'layer')}
                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                />
            </div>
            {/* attributes */}
            <div className="mb-5">
                <div className="flex justify-between items-center mb-2">
                    <h4 className="text-md font-medium text-gray-700">Labels</h4>
                    <button
                        className="px-2 py-1 bg-blue-500 text-white rounded hover:bg-blue-600 text-sm"
                        onClick={() => addAttributeToAssociation(0.5, "sample label")}
                    >
                        Add
                    </button>
                </div>
                {selectedAssociation?.attributes?.map((attr, attrIndex) => (
                    <div key={`attr-${attrIndex}`} className="mb-4 p-3 border border-gray-300 rounded-md bg-white">
                        <div className="mb-3">
                            <label className="block mb-1 text-sm font-medium text-gray-700">Content:</label>
                            <input
                                type="text"
                                value={getValue(`attributes:${attrIndex}.content`, attr.content)}
                                onChange={(e) => handleInputChange(`attributes:${attrIndex}.content`, e.target.value)}
                                onKeyPress={(e) => handleKeyPress(e, `attributes:${attrIndex}.content`)}
                                className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                            />
                        </div>
                        <div className="mb-3">
                            <label className="block mb-1 text-sm font-medium text-gray-700">Font Size:</label>
                            <input
                                type="number"
                                value={getValue(`attributes:${attrIndex}.fontSize`, attr.fontSize)}
                                onChange={(e) => handleInputChange(`attributes:${attrIndex}.fontSize`, parseInt(e.target.value))}
                                onKeyPress={(e) => handleKeyPress(e, `attributes:${attrIndex}.fontSize`)}
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
                                value={getValue(`attributes:${attrIndex}.fontFile`, attr.fontFile)}
                                onChange={(e) => handleInputChange(`attributes:${attrIndex}.fontFile`, e.target.value)}
                                onBlur={() => {
                                    const property = `attributes:${attrIndex}.fontFile`;
                                    const value = pendingChanges[property];
                                    if (value !== undefined) {
                                        updateAssociationProperty(property, value);
                                        setPendingChanges(prev => {
                                            const updated = { ...prev };
                                            delete updated[property];
                                            return updated;
                                        });
                                    }
                                }}
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
                                value={getValue(`attributes:${attrIndex}.ratio`, attr.ratio)}
                                onChange={(e) => handleInputChange(`attributes:${attrIndex}.ratio`, parseFloat(e.target.value))}
                                onKeyPress={(e) => handleKeyPress(e, `attributes:${attrIndex}.ratio`)}
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
