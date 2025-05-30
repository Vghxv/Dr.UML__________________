import React, {useEffect, useRef, useState} from "react";
import {GadgetProps} from "../utils/Props";
import {attribute} from "../../wailsjs/go/models";

interface GadgetPropertiesPanelProps {
    selectedGadget: GadgetProps | null;
    updateGadgetProperty: (property: string, value: any) => void;
    addAttributeToGadget: (section: number, content: string) => void;
}

const GadgetPropertiesPanel: React.FC<GadgetPropertiesPanelProps> = ({
        selectedGadget,
        updateGadgetProperty,
        addAttributeToGadget
    }) => {
    const [focusedInput, setFocusedInput] = useState<string | null>(null);
    const inputRefs = useRef<{ [key: string]: HTMLInputElement | HTMLSelectElement | null }>({});

    // Restore focus after re-render
    useEffect(() => {
        if (focusedInput && inputRefs.current[focusedInput]) {
            inputRefs.current[focusedInput]?.focus();
        }
    }, [selectedGadget, focusedInput]);

    if (!selectedGadget) return null;

    return (
        <div className="absolute right-0 top-0 w-[300px] h-full bg-gray-100 p-5 shadow-md overflow-y-auto">
            <h3 className="text-xl font-semibold text-gray-800 mb-4">Gadget Properties</h3>

            {/* x */}
            <div className="mb-4">
                <label className="block mb-1 text-sm font-medium text-gray-700">X Position:</label>
                <input
                    type="number"
                    value={selectedGadget.x}
                    ref={(el) => inputRefs.current['x'] = el}
                    onFocus={() => setFocusedInput('x')}
                    onChange={(e) => updateGadgetProperty('x', parseInt(e.target.value))}
                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                />
            </div>

            {/* y */}
            <div className="mb-4">
                <label className="block mb-1 text-sm font-medium text-gray-700">Y Position:</label>
                <input
                    type="number"
                    value={selectedGadget.y}
                    ref={(el) => inputRefs.current['y'] = el}
                    onFocus={() => setFocusedInput('y')}
                    onChange={(e) => updateGadgetProperty('y', parseInt(e.target.value))}
                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                />
            </div>


            {/* layer */}
            <div className="mb-4">
                <label className="block mb-1 text-sm font-medium text-gray-700">Layer:</label>
                <input
                    type="number"
                    value={selectedGadget.layer}
                    ref={(el) => inputRefs.current['layer'] = el}
                    onFocus={() => setFocusedInput('layer')}
                    onChange={(e) => updateGadgetProperty('layer', parseInt(e.target.value))}
                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                />
            </div>

            {/* color */}
            <div className="mb-4">
                <label className="block mb-1 text-sm font-medium text-gray-700">Color:</label>
                <input
                    type="color"
                    value={selectedGadget.color}
                    ref={(el) => inputRefs.current['color'] = el}
                    onFocus={() => setFocusedInput('color')}
                    onChange={(e) => updateGadgetProperty('color', e.target.value)}
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
                                    value={attr.content}
                                    ref={(el) => inputRefs.current[`attributes${groupIndex}:${attrIndex}.content`] = el}
                                    onFocus={() => setFocusedInput(`attributes${groupIndex}:${attrIndex}.content`)}
                                    onChange={(e) => updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.content`, e.target.value)}
                                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                                />
                            </div>

                            <div className="mb-3">
                                <label className="block mb-1 text-sm font-medium text-gray-700">Font Size:</label>
                                <input
                                    type="number"
                                    value={attr.fontSize}
                                    ref={(el) => inputRefs.current[`attributes${groupIndex}:${attrIndex}.fontSize`] = el}
                                    onFocus={() => setFocusedInput(`attributes${groupIndex}:${attrIndex}.fontSize`)}
                                    onChange={(e) => updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.fontSize`, parseInt(e.target.value))}
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
                                                newStyle &= ~attribute.Textstyle.Bold; // Remove bold bit
                                            } else {
                                                newStyle |= attribute.Textstyle.Bold; // Add bold bit
                                            }
                                            updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.fontStyleB`, newStyle);
                                        }}
                                        className={`px-3 py-2 border rounded-md ${
                                            (attr.fontStyle & attribute.Textstyle.Bold) !== 0 
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
                                                newStyle &= ~attribute.Textstyle.Italic; // Remove italic bit
                                            } else {
                                                newStyle |= attribute.Textstyle.Italic; // Add italic bit
                                            }
                                            updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.fontStyleI`, newStyle);
                                        }}
                                        className={`px-3 py-2 border rounded-md ${
                                            (attr.fontStyle & attribute.Textstyle.Italic) !== 0 
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
                                                newStyle &= ~attribute.Textstyle.Underline; // Remove underline bit
                                            } else {
                                                newStyle |= attribute.Textstyle.Underline; // Add underline bit
                                            }
                                            updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.fontStyleU`, newStyle);
                                        }}
                                        className={`px-3 py-2 border rounded-md ${
                                            (attr.fontStyle & attribute.Textstyle.Underline) !== 0 
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
                                    ref={(el) => inputRefs.current[`attributes${groupIndex}:${attrIndex}.fontFile`] = el}
                                    onFocus={() => setFocusedInput(`attributes${groupIndex}:${attrIndex}.fontFile`)}
                                    onChange={(e) => updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.fontFile`, e.target.value)}
                                    className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-black"
                                >
                                    <option value="Arial">Arial</option>
                                    <option value="Helvetica">Helvetica</option>
                                    <option value="Times New Roman">Times New Roman</option>
                                    <option value="Courier New">Courier New</option>
                                    <option value="Georgia">Georgia</option>
                                    <option value="Verdana">Verdana</option>
                                </select>
                            </div>
                        </div>
                    ))}
                </div>
            ))}
        </div>
    );
};

export default GadgetPropertiesPanel;
