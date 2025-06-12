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
        <div className="p-6">
            {/* Industrial Header */}
            <div 
                className="mb-6 p-4 border-2 border-[#4682B4]"
                style={{
                    background: 'linear-gradient(145deg, #4682B4, #333333)',
                    boxShadow: '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                }}
            >
                <h3 className="text-lg font-bold text-[#F2F2F0] tracking-wide uppercase">ASSOCIATION PROPERTIES</h3>
            </div>
            
            {/* Layer Control */}
            <div className="mb-6">
                <label className="block mb-2 text-sm font-bold text-[#F2F2F0] tracking-wide uppercase">LAYER:</label>
                <input
                    type="number"
                    value={getValue('layer', selectedAssociation.layer)}
                    onChange={(e) => handleInputChange('layer', parseInt(e.target.value))}
                    onKeyPress={(e) => handleKeyPress(e, 'layer')}
                    className="w-full p-2 bg-[#1C1C1C] border-2 border-[#918175] text-[#F2F2F0] font-bold focus:border-[#B87333] focus:outline-none"
                    style={{
                        boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.5)'
                    }}
                />
            </div>
            
            {/* Labels Section */}
            <div className="mb-6">
                <div 
                    className="flex justify-between items-center p-3 mb-4 border-2 border-[#556B2F]"
                    style={{
                        background: 'linear-gradient(145deg, #556B2F, #333333)',
                        boxShadow: '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                    }}
                >
                    <h4 className="text-sm font-bold text-[#F2F2F0] tracking-wide uppercase">LABELS</h4>
                    <button
                        className="px-3 py-1 text-[#1C1C1C] font-bold text-xs tracking-wide uppercase transition-all duration-200 border-2 border-[#B87333] hover:border-[#B7410E]"
                        style={{
                            background: 'linear-gradient(145deg, #B87333, #918175)',
                            boxShadow: '2px 2px 4px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                        }}
                        onClick={() => addAttributeToAssociation(0.5, "sample label")}
                    >
                        ADD
                    </button>
                </div>
                {/* Labels List */}
                {selectedAssociation?.attributes?.map((attr, attrIndex) => (
                    <div 
                        key={`attr-${attrIndex}`} 
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
                                value={getValue(`attributes:${attrIndex}.content`, attr.content)}
                                onChange={(e) => handleInputChange(`attributes:${attrIndex}.content`, e.target.value)}
                                onKeyPress={(e) => handleKeyPress(e, `attributes:${attrIndex}.content`)}
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
                                value={getValue(`attributes:${attrIndex}.fontSize`, attr.fontSize)}
                                onChange={(e) => handleInputChange(`attributes:${attrIndex}.fontSize`, parseInt(e.target.value))}
                                onKeyPress={(e) => handleKeyPress(e, `attributes:${attrIndex}.fontSize`)}
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
                                        updateAssociationProperty(`attributes:${attrIndex}.fontStyle${styleType}`, newStyle);
                                    }}
                                />
                            </div>
                        </div>
                        <div className="mb-3">
                            <label className="block mb-2 text-xs font-bold text-[#F2F2F0] tracking-wide uppercase">FONT:</label>
                            <select
                                value={getValue(`attributes:${attrIndex}.fontFile`, attr.fontFile)}
                                onChange={(e) => handleInputChange(`attributes:${attrIndex}.fontFile`, e.target.value)}
                                onBlur={() => handleBlur(`attributes:${attrIndex}.fontFile`)}
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
                        <div className="mb-3">
                            <label className="block mb-2 text-xs font-bold text-[#F2F2F0] tracking-wide uppercase">RATIO (0~1):</label>
                            <input
                                type="number"
                                min={0}
                                max={1}
                                step={0.01}
                                value={getValue(`attributes:${attrIndex}.ratio`, attr.ratio)}
                                onChange={(e) => handleInputChange(`attributes:${attrIndex}.ratio`, parseFloat(e.target.value))}
                                onKeyPress={(e) => handleKeyPress(e, `attributes:${attrIndex}.ratio`)}
                                className="w-full p-2 bg-[#1C1C1C] border-2 border-[#4682B4] text-[#F2F2F0] font-bold focus:border-[#B87333] focus:outline-none"
                                style={{
                                    boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.5)'
                                }}
                            />
                        </div>

                        {/* Delete button */}
                        <div className="flex justify-end">
                            <button
                                type="button"
                                onClick={() => updateAssociationProperty(`attributes:${attrIndex}.delete`, true)}
                                className="px-3 py-1 text-[#F2F2F0] font-bold text-xs tracking-wide uppercase transition-all duration-200 border-2 border-[#B7410E] hover:border-[#B87333]"
                                style={{
                                    background: 'linear-gradient(145deg, #B7410E, #1C1C1C)',
                                    boxShadow: '2px 2px 4px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                                }}
                            >
                                DELETE
                            </button>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default AssociationPropertiesPanel;
