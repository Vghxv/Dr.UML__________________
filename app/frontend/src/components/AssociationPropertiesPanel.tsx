import React from "react";
import { AssociationProps } from "../utils/Props";
import { usePendingChanges } from "../hooks/usePendingChanges";
import { getFontOptions } from "../utils/fontUtils";
import { FontStyleButtons } from "./FontStyleButtons";

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
    // Use shared pending changes hook
    const { handleInputChange, handleKeyPress, handleBlur, getValue } = usePendingChanges({
        onPropertyUpdate: updateAssociationProperty,
        dependencyArray: [selectedAssociation.endX, selectedAssociation.endY]
    });

    return (
        <div className="absolute right-0 top-35 w-[300px] h-full bg-gray-100 p-5 shadow-md overflow-y-auto">
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
                            <FontStyleButtons
                                fontStyle={attr.fontStyle}
                                onStyleChange={(newStyle, styleType) => {
                                    updateAssociationProperty(`attributes:${attrIndex}.fontStyle${styleType}`, newStyle);
                                }}
                            />
                        </div>
                        <div className="mb-3">
                            <label className="block mb-1 text-sm font-medium text-gray-700">Font File:</label>
                            <select
                                value={getValue(`attributes:${attrIndex}.fontFile`, attr.fontFile)}
                                onChange={(e) => handleInputChange(`attributes:${attrIndex}.fontFile`, e.target.value)}
                                onBlur={() => handleBlur(`attributes:${attrIndex}.fontFile`)}
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

                        {/* Delete button */}
                        <div className="flex justify-end">
                            <button
                                type="button"
                                onClick={() => updateAssociationProperty(`attributes:${attrIndex}.delete`, true)}
                                className="px-3 py-1 bg-red-500 text-white rounded hover:bg-red-600 text-sm"
                            >
                                Delete
                            </button>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default AssociationPropertiesPanel;
