import React from "react";
import { GadgetProps } from "../utils/Props";
import { attribute } from "../../wailsjs/go/models";
import { usePendingChanges } from "../hooks/usePendingChanges";
import { getFontOptions } from "../utils/fontUtils";
import { FontStyleButtons } from "./FontStyleButtons";

interface GadgetPropertiesPanelProps {
    selectedGadget: GadgetProps;
    updateGadgetProperty: (property: string, value: any) => void;
    addAttributeToGadget: (section: number, content: string) => void;
}

const GadgetPropertiesPanel: React.FC<GadgetPropertiesPanelProps> = ({
    selectedGadget,
    updateGadgetProperty,
    addAttributeToGadget
}) => {
    // Use shared pending changes hook
    const { handleInputChange, handleKeyPress, handleBlur, getValue } = usePendingChanges({
        onPropertyUpdate: updateGadgetProperty,
        dependencyArray: [selectedGadget.x, selectedGadget.y],
        formatProperty: (property: string) => {
            // Handle attribute properties specially to match the expected format
            if (property.includes('attributes')) {
                const [attrPath, attrProperty] = property.split('.');
                return `${attrPath}.${attrProperty}`;
            }
            return property;
        }
    });

    return (
        <div className="absolute right-0 top-0 w-[300px] h-full bg-gray-100 p-5 shadow-md overflow-y-auto">
            <h3 className="text-xl font-semibold text-gray-800 mb-4">Gadget Properties</h3>
            {/* x */}
            <div className="mb-4">
                <label className="block mb-1 text-sm font-medium text-gray-700">X Position:</label>
                <input
                    type="number"
                    value={getValue('x', selectedGadget.x)}
                    onChange={(e) => handleInputChange('x', parseInt(e.target.value))}
                    onKeyPress={(e) => handleKeyPress(e, 'x')}
                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                />
            </div>
            {/* y */}
            <div className="mb-4">
                <label className="block mb-1 text-sm font-medium text-gray-700">Y Position:</label>
                <input
                    type="number"
                    value={getValue('y', selectedGadget.y)}
                    onChange={(e) => handleInputChange('y', parseInt(e.target.value))}
                    onKeyPress={(e) => handleKeyPress(e, 'y')}
                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                />
            </div>
            {/* layer */}
            <div className="mb-4">
                <label className="block mb-1 text-sm font-medium text-gray-700">Layer:</label>
                <input
                    type="number"
                    value={getValue('layer', selectedGadget.layer)}
                    onChange={(e) => handleInputChange('layer', parseInt(e.target.value))}
                    onKeyPress={(e) => handleKeyPress(e, 'layer')}
                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                />
            </div>
            {/* color */}
            <div className="mb-4">
                <label className="block mb-1 text-sm font-medium text-gray-700">Color:</label>
                <input
                    type="color"
                    value={getValue('color', selectedGadget.color)}
                    onChange={(e) => handleInputChange('color', e.target.value)}
                    onBlur={() => handleBlur('color')}
                    className="w-full h-10 p-1 border border-gray-300 rounded text-black"
                />
            </div>
            {/* backend attr */}
            {selectedGadget.attributes.map((attrGroup, groupIndex) => (
                <div key={`group-${groupIndex}`} className="mb-5">
                    <div className="flex justify-between items-center mb-2">
                        <h4 className="text-md font-medium text-gray-700">{groupIndex === 0 ? "Class Name" : groupIndex === 1 ? "Attributes" : "Methods"}</h4>
                        {groupIndex === 1 && (
                            <button
                                className="px-2 py-1 bg-blue-500 text-white rounded hover:bg-blue-600 text-sm"
                                onClick={() => addAttributeToGadget(groupIndex, "sample attribute")}
                            >
                                Add
                            </button>
                        )}
                        {groupIndex === 2 && (
                            <button
                                className="px-2 py-1 bg-blue-500 text-white rounded hover:bg-blue-600 text-sm"
                                onClick={() => addAttributeToGadget(groupIndex, "sample method")}
                            >
                                Add
                            </button>
                        )}
                    </div>
                    {attrGroup.map((attr, attrIndex) => (
                        <div key={`attr-${groupIndex}-${attrIndex}`}
                            className="mb-4 p-3 border border-gray-300 rounded-md bg-white">
                            <div className="mb-3">
                                <label className="block mb-1 text-sm font-medium text-gray-700">Content:</label>
                                <input
                                    type="text"
                                    value={getValue(`attributes${groupIndex}:${attrIndex}.content`, attr.content)}
                                    onChange={(e) => handleInputChange(`attributes${groupIndex}:${attrIndex}.content`, e.target.value)}
                                    onKeyPress={(e) => handleKeyPress(e, `attributes${groupIndex}:${attrIndex}.content`)}
                                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                                />
                            </div>
                            <div className="mb-3">
                                <label className="block mb-1 text-sm font-medium text-gray-700">Font Size:</label>
                                <input
                                    type="number"
                                    value={getValue(`attributes${groupIndex}:${attrIndex}.fontSize`, attr.fontSize)}
                                    onChange={(e) => handleInputChange(`attributes${groupIndex}:${attrIndex}.fontSize`, parseInt(e.target.value))}
                                    onKeyPress={(e) => handleKeyPress(e, `attributes${groupIndex}:${attrIndex}.fontSize`)}
                                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                                />
                            </div>
                            <div className="mb-3">
                                <label className="block mb-1 text-sm font-medium text-gray-700">Font Style:</label>
                                <FontStyleButtons
                                    fontStyle={attr.fontStyle}
                                    onStyleChange={(newStyle, styleType) => {
                                        updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.fontStyle${styleType}`, newStyle);
                                    }}
                                />
                            </div>
                            <div className="mb-3">
                                <label className="block mb-1 text-sm font-medium text-gray-700">Font File:</label>
                                <select
                                    value={getValue(`attributes${groupIndex}:${attrIndex}.fontFile`, attr.fontFile)}
                                    onChange={(e) => handleInputChange(`attributes${groupIndex}:${attrIndex}.fontFile`, e.target.value)}
                                    onBlur={() => handleBlur(`attributes${groupIndex}:${attrIndex}.fontFile`)}
                                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                                >
                                    {getFontOptions().map((fontName) => (
                                        <option key={fontName} value={fontName}>{fontName}</option>
                                    ))}
                                </select>
                            </div>

                            {/* Delete button - only show for attributes and methods sections (not class name) */}
                            {groupIndex > 0 && (
                                <div className="flex justify-end">
                                    <button
                                        type="button"
                                        onClick={() => updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.delete`, true)}
                                        className="px-3 py-1 bg-red-500 text-white rounded hover:bg-red-600 text-sm"
                                    >
                                        Delete
                                    </button>
                                </div>
                            )}
                        </div>
                    ))}
                </div>
            ))}
        </div>
    );
};

export default GadgetPropertiesPanel;
