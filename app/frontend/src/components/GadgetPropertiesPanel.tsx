import React, { useState, useRef, useEffect } from "react";
import { GadgetProps } from "../utils/Props";

interface GadgetPropertiesPanelProps {
    selectedGadget: GadgetProps | null;
    updateGadgetProperty: (property: string, value: any) => void;
}

const TEXT_COLOR = 'black';

const GadgetPropertiesPanel: React.FC<GadgetPropertiesPanelProps> = ({ selectedGadget, updateGadgetProperty }) => {
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
        <div style={{
            position: 'absolute',
            right: 0,
            top: 0,
            width: '300px',
            height: '100%',
            backgroundColor: '#f0f0f0',
            padding: '20px',
            boxShadow: '-2px 0 5px rgba(0,0,0,0.2)',
            overflowY: 'auto'
        }}>
            <h3 style={{color: TEXT_COLOR}}>Gadget Properties</h3>

            {/* Basic properties */}
            <div style={{marginBottom: '15px'}}>
                <label style={{display: 'block', marginBottom: '5px', color: TEXT_COLOR}}>X Position:</label>
                <input
                    type="number"
                    value={selectedGadget.x}
                    ref={(el) => inputRefs.current['x'] = el}
                    onFocus={() => setFocusedInput('x')}
                    onChange={(e) => updateGadgetProperty('x', parseInt(e.target.value))}
                    style={{width: '100%', padding: '5px'}}
                />
            </div>

            <div style={{marginBottom: '15px'}}>
                <label style={{display: 'block', marginBottom: '5px', color: TEXT_COLOR}}>Y Position:</label>
                <input
                    type="number"
                    value={selectedGadget.y}
                    ref={(el) => inputRefs.current['y'] = el}
                    onFocus={() => setFocusedInput('y')}
                    onChange={(e) => updateGadgetProperty('y', parseInt(e.target.value))}
                    style={{width: '100%', padding: '5px'}}
                />
            </div>

            <div style={{marginBottom: '15px'}}>
                <label style={{display: 'block', marginBottom: '5px', color: TEXT_COLOR}}>Layer:</label>
                <input
                    type="number"
                    value={selectedGadget.layer}
                    ref={(el) => inputRefs.current['layer'] = el}
                    onFocus={() => setFocusedInput('layer')}
                    onChange={(e) => updateGadgetProperty('layer', parseInt(e.target.value))}
                    style={{width: '100%', padding: '5px'}}
                />
            </div>

            <div style={{marginBottom: '15px'}}>
                <label style={{display: 'block', marginBottom: '5px', color: TEXT_COLOR}}>Color:</label>
                <input
                    type="color"
                    value={selectedGadget.color}
                    ref={(el) => inputRefs.current['color'] = el}
                    onFocus={() => setFocusedInput('color')}
                    onChange={(e) => updateGadgetProperty('color', e.target.value)}
                    style={{width: '100%', padding: '5px'}}
                />
            </div>

            {/* Attributes */}
            <h4 style={{color: TEXT_COLOR}}>Attributes</h4>
            {selectedGadget.attributes.map((attrGroup, groupIndex) => (
                <div key={`group-${groupIndex}`} style={{marginBottom: '20px'}}>
                    <h5 style={{color: TEXT_COLOR}}>Group {groupIndex + 1}</h5>
                    {attrGroup.map((attr, attrIndex) => (
                        <div key={`attr-${groupIndex}-${attrIndex}`} style={{
                            marginBottom: '15px',
                            padding: '10px',
                            border: '1px solid #ddd',
                            borderRadius: '5px'
                        }}>
                            <div style={{marginBottom: '10px'}}>
                                <label style={{display: 'block', marginBottom: '5px', color: TEXT_COLOR}}>Content:</label>
                                <input
                                    type="text"
                                    value={attr.content}
                                    ref={(el) => inputRefs.current[`attributes${groupIndex}:${attrIndex}.content`] = el}
                                    onFocus={() => setFocusedInput(`attributes${groupIndex}:${attrIndex}.content`)}
                                    onChange={(e) => updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.content`, e.target.value)}
                                    style={{width: '100%', padding: '5px'}}
                                />
                            </div>

                            <div style={{marginBottom: '10px'}}>
                                <label style={{display: 'block', marginBottom: '5px', color: TEXT_COLOR}}>Font Size:</label>
                                <input
                                    type="number"
                                    value={attr.fontSize}
                                    ref={(el) => inputRefs.current[`attributes${groupIndex}:${attrIndex}.fontSize`] = el}
                                    onFocus={() => setFocusedInput(`attributes${groupIndex}:${attrIndex}.fontSize`)}
                                    onChange={(e) => updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.fontSize`, parseInt(e.target.value))}
                                    style={{width: '100%', padding: '5px'}}
                                />
                            </div>

                            <div style={{marginBottom: '10px'}}>
                                <label style={{display: 'block', marginBottom: '5px', color: TEXT_COLOR}}>Font Style:</label>
                                <select
                                    value={attr.fontStyle}
                                    ref={(el) => inputRefs.current[`attributes${groupIndex}:${attrIndex}.fontStyle`] = el}
                                    onFocus={() => setFocusedInput(`attributes${groupIndex}:${attrIndex}.fontStyle`)}
                                    onChange={(e) => updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.fontStyle`, parseInt(e.target.value))}
                                    style={{width: '100%', padding: '5px'}}
                                >
                                    <option value={0} style={{color: TEXT_COLOR}}>Normal</option>
                                    <option value={1} style={{color: TEXT_COLOR}}>Italic</option>
                                    <option value={2} style={{color: TEXT_COLOR}}>Bold</option>
                                    <option value={3} style={{color: TEXT_COLOR}}>Bold Italic</option>
                                </select>
                            </div>

                            <div style={{marginBottom: '10px'}}>
                                <label style={{display: 'block', marginBottom: '5px', color: TEXT_COLOR}}>Font File:</label>
                                <select
                                    value={attr.fontFile}
                                    ref={(el) => inputRefs.current[`attributes${groupIndex}:${attrIndex}.fontFile`] = el}
                                    onFocus={() => setFocusedInput(`attributes${groupIndex}:${attrIndex}.fontFile`)}
                                    onChange={(e) => updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.fontFile`, e.target.value)}
                                    style={{width: '100%', padding: '5px'}}
                                >
                                    <option value="Arial" style={{color: TEXT_COLOR}}>Arial</option>
                                    <option value="Helvetica" style={{color: TEXT_COLOR}}>Helvetica</option>
                                    <option value="Times New Roman" style={{color: TEXT_COLOR}}>Times New Roman</option>
                                    <option value="Courier New" style={{color: TEXT_COLOR}}>Courier New</option>
                                    <option value="Georgia" style={{color: TEXT_COLOR}}>Georgia</option>
                                    <option value="Verdana" style={{color: TEXT_COLOR}}>Verdana</option>
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
