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
        <div className="p-6">
            {/* Industrial Header */}
            <div 
                className="mb-6 p-4 border-2 border-[#4682B4]"
                style={{
                    background: 'linear-gradient(145deg, #4682B4, #333333)',
                    boxShadow: '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                }}
            >
                <h3 className="text-lg font-bold text-[#F2F2F0] tracking-wide uppercase">GADGET PROPERTIES</h3>
            </div>
            
            {/* Position Controls */}
            <div className="grid grid-cols-2 gap-4 mb-6">
                <div>
                    <label className="block mb-2 text-sm font-bold text-[#F2F2F0] tracking-wide uppercase">X POS:</label>
                    <input
                        type="number"
                        value={getValue('x', selectedGadget.x)}
                        onChange={(e) => handleInputChange('x', parseInt(e.target.value))}
                        onKeyPress={(e) => handleKeyPress(e, 'x')}
                        className="w-full p-2 bg-[#1C1C1C] border-2 border-[#918175] text-[#F2F2F0] font-bold focus:border-[#B87333] focus:outline-none"
                        style={{
                            boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.5)'
                        }}
                    />
                </div>
                <div>
                    <label className="block mb-2 text-sm font-bold text-[#F2F2F0] tracking-wide uppercase">Y POS:</label>
                    <input
                        type="number"
                        value={getValue('y', selectedGadget.y)}
                        onChange={(e) => handleInputChange('y', parseInt(e.target.value))}
                        onKeyPress={(e) => handleKeyPress(e, 'y')}
                        className="w-full p-2 bg-[#1C1C1C] border-2 border-[#918175] text-[#F2F2F0] font-bold focus:border-[#B87333] focus:outline-none"
                        style={{
                            boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.5)'
                        }}
                    />
                </div>
            </div>
            
            {/* Layer Control */}
            <div className="mb-6">
                <label className="block mb-2 text-sm font-bold text-[#F2F2F0] tracking-wide uppercase">LAYER:</label>
                <input
                    type="number"
                    value={getValue('layer', selectedGadget.layer)}
                    onChange={(e) => handleInputChange('layer', parseInt(e.target.value))}
                    onKeyPress={(e) => handleKeyPress(e, 'layer')}
                    className="w-full p-2 bg-[#1C1C1C] border-2 border-[#918175] text-[#F2F2F0] font-bold focus:border-[#B87333] focus:outline-none"
                    style={{
                        boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.5)'
                    }}
                />
            </div>
            
            {/* Color Control */}
            <div className="mb-6">
                <label className="block mb-2 text-sm font-bold text-[#F2F2F0] tracking-wide uppercase">COLOR:</label>
                <div 
                    className="p-2 border-2 border-[#918175] inline-block"
                    style={{
                        background: 'linear-gradient(145deg, #333333, #1C1C1C)',
                        boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.5)'
                    }}
                >
                    <input
                        type="color"
                        value={getValue('color', selectedGadget.color)}
                        onChange={(e) => handleInputChange('color', e.target.value)}
                        onBlur={() => handleBlur('color')}
                        className="w-16 h-8 border-0 cursor-pointer"
                    />
                </div>
            </div>
            {/* Attributes Sections */}
            {selectedGadget.attributes.map((attrGroup, groupIndex) => (
                <div key={`group-${groupIndex}`} className="mb-6">
                    {/* Section Header */}
                    <div 
                        className="flex justify-between items-center p-3 mb-4 border-2 border-[#556B2F]"
                        style={{
                            background: 'linear-gradient(145deg, #556B2F, #333333)',
                            boxShadow: '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                        }}
                    >
                        <h4 className="text-sm font-bold text-[#F2F2F0] tracking-wide uppercase">
                            {groupIndex === 0 ? "CLASS NAME" : groupIndex === 1 ? "ATTRIBUTES" : "METHODS"}
                        </h4>
                        {groupIndex === 1 && (
                            <button
                                className="px-3 py-1 text-[#1C1C1C] font-bold text-xs tracking-wide uppercase transition-all duration-200 border-2 border-[#B87333] hover:border-[#B7410E]"
                                style={{
                                    background: 'linear-gradient(145deg, #B87333, #918175)',
                                    boxShadow: '2px 2px 4px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                                }}
                                onClick={() => addAttributeToGadget(groupIndex, "sample attribute")}
                            >
                                ADD
                            </button>
                        )}
                        {groupIndex === 2 && (
                            <button
                                className="px-3 py-1 text-[#1C1C1C] font-bold text-xs tracking-wide uppercase transition-all duration-200 border-2 border-[#B87333] hover:border-[#B7410E]"
                                style={{
                                    background: 'linear-gradient(145deg, #B87333, #918175)',
                                    boxShadow: '2px 2px 4px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                                }}
                                onClick={() => addAttributeToGadget(groupIndex, "sample method")}
                            >
                                ADD
                            </button>
                        )}
                    </div>
                    
                    {/* Attributes List */}
                    {attrGroup.map((attr, attrIndex) => (
                        <div 
                            key={`attr-${groupIndex}-${attrIndex}`}
                            className="mb-4 p-4 border-2 border-[#918175]"
                            style={{
                                background: 'linear-gradient(145deg, #918175, #1C1C1C)',
                                boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.3)'
                            }}
                        >
                            <div className="mb-3">
                                <label className="block mb-2 text-xs font-bold text-[#F2F2F0] tracking-wide uppercase">CONTENT:</label>
                                <input
                                    type="text"
                                    value={getValue(`attributes${groupIndex}:${attrIndex}.content`, attr.content)}
                                    onChange={(e) => handleInputChange(`attributes${groupIndex}:${attrIndex}.content`, e.target.value)}
                                    onKeyPress={(e) => handleKeyPress(e, `attributes${groupIndex}:${attrIndex}.content`)}
                                    className="w-full p-2 bg-[#1C1C1C] border-2 border-[#4682B4] text-[#F2F2F0] font-bold focus:border-[#B87333] focus:outline-none"
                                    style={{
                                        boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.5)'
                                    }}
                                />
                            </div>
                            <div className="mb-3">
                                <label className="block mb-2 text-xs font-bold text-[#F2F2F0] tracking-wide uppercase">FONT SIZE:</label>
                                <input
                                    type="number"
                                    value={getValue(`attributes${groupIndex}:${attrIndex}.fontSize`, attr.fontSize)}
                                    onChange={(e) => handleInputChange(`attributes${groupIndex}:${attrIndex}.fontSize`, parseInt(e.target.value))}
                                    onKeyPress={(e) => handleKeyPress(e, `attributes${groupIndex}:${attrIndex}.fontSize`)}
                                    className="w-full p-2 bg-[#1C1C1C] border-2 border-[#4682B4] text-[#F2F2F0] font-bold focus:border-[#B87333] focus:outline-none"
                                    style={{
                                        boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.5)'
                                    }}
                                />
                            </div>
                            <div className="mb-3">
                                <label className="block mb-2 text-xs font-bold text-[#F2F2F0] tracking-wide uppercase">FONT STYLE:</label>
                                <div 
                                    className="p-2 border-2 border-[#918175]"
                                    style={{
                                        background: 'linear-gradient(145deg, #333333, #1C1C1C)',
                                        boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.5)'
                                    }}
                                >
                                    <FontStyleButtons
                                        fontStyle={attr.fontStyle}
                                        onStyleChange={(newStyle, styleType) => {
                                            updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.fontStyle${styleType}`, newStyle);
                                        }}
                                    />
                                </div>
                            </div>
                            <div className="mb-3">
                                <label className="block mb-2 text-xs font-bold text-[#F2F2F0] tracking-wide uppercase">FONT:</label>
                                <select
                                    value={getValue(`attributes${groupIndex}:${attrIndex}.fontFile`, attr.fontFile)}
                                    onChange={(e) => handleInputChange(`attributes${groupIndex}:${attrIndex}.fontFile`, e.target.value)}
                                    onBlur={() => handleBlur(`attributes${groupIndex}:${attrIndex}.fontFile`)}
                                    className="w-full p-2 bg-[#1C1C1C] border-2 border-[#4682B4] text-[#F2F2F0] font-bold focus:border-[#B87333] focus:outline-none"
                                    style={{
                                        boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.5)'
                                    }}
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
                                        className="px-3 py-1 text-[#F2F2F0] font-bold text-xs tracking-wide uppercase transition-all duration-200 border-2 border-[#B7410E] hover:border-[#B87333]"
                                        style={{
                                            background: 'linear-gradient(145deg, #B7410E, #1C1C1C)',
                                            boxShadow: '2px 2px 4px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                                        }}
                                    >
                                        DELETE
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
