import React from "react";
import { GadgetProps } from "../utils/Props";

interface GadgetPropertiesPanelProps {
    selectedGadget: GadgetProps | null;
    updateGadgetProperty: (property: string, value: any) => void;
}

const GadgetPropertiesPanel: React.FC<GadgetPropertiesPanelProps> = ({ selectedGadget, updateGadgetProperty }) => {
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
            <h3>Gadget Properties</h3>

            {/* Basic properties */}
            <div style={{marginBottom: '15px'}}>
                <label style={{display: 'block', marginBottom: '5px'}}>X Position:</label>
                <input
                    type="number"
                    value={selectedGadget.x}
                    onChange={(e) => updateGadgetProperty('x', parseInt(e.target.value))}
                    style={{width: '100%', padding: '5px'}}
                />
            </div>

            <div style={{marginBottom: '15px'}}>
                <label style={{display: 'block', marginBottom: '5px'}}>Y Position:</label>
                <input
                    type="number"
                    value={selectedGadget.y}
                    onChange={(e) => updateGadgetProperty('y', parseInt(e.target.value))}
                    style={{width: '100%', padding: '5px'}}
                />
            </div>

            <div style={{marginBottom: '15px'}}>
                <label style={{display: 'block', marginBottom: '5px'}}>Layer:</label>
                <input
                    type="number"
                    value={selectedGadget.layer}
                    onChange={(e) => updateGadgetProperty('layer', parseInt(e.target.value))}
                    style={{width: '100%', padding: '5px'}}
                />
            </div>

            <div style={{marginBottom: '15px'}}>
                <label style={{display: 'block', marginBottom: '5px'}}>Color:</label>
                <input
                    type="color"
                    value={selectedGadget.color}
                    onChange={(e) => updateGadgetProperty('color', e.target.value)}
                    style={{width: '100%', padding: '5px'}}
                />
            </div>

            {/* Attributes */}
            <h4>Attributes</h4>
            {selectedGadget.attributes.map((attrGroup, groupIndex) => (
                <div key={`group-${groupIndex}`} style={{marginBottom: '20px'}}>
                    <h5>Group {groupIndex + 1}</h5>
                    {attrGroup.map((attr, attrIndex) => (
                        <div key={`attr-${groupIndex}-${attrIndex}`} style={{
                            marginBottom: '15px',
                            padding: '10px',
                            border: '1px solid #ddd',
                            borderRadius: '5px'
                        }}>
                            <div style={{marginBottom: '10px'}}>
                                <label style={{display: 'block', marginBottom: '5px'}}>Content:</label>
                                <input
                                    type="text"
                                    value={attr.content}
                                    onChange={(e) => updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.content`, e.target.value)}
                                    style={{width: '100%', padding: '5px'}}
                                />
                            </div>

                            <div style={{marginBottom: '10px'}}>
                                <label style={{display: 'block', marginBottom: '5px'}}>Font Size:</label>
                                <input
                                    type="number"
                                    value={attr.fontSize}
                                    onChange={(e) => updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.fontSize`, parseInt(e.target.value))}
                                    style={{width: '100%', padding: '5px'}}
                                />
                            </div>

                            <div style={{marginBottom: '10px'}}>
                                <label style={{display: 'block', marginBottom: '5px'}}>Font Style:</label>
                                <select
                                    value={attr.fontStyle}
                                    onChange={(e) => updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.fontStyle`, parseInt(e.target.value))}
                                    style={{width: '100%', padding: '5px'}}
                                >
                                    <option value={0}>Normal</option>
                                    <option value={1}>Italic</option>
                                    <option value={2}>Bold</option>
                                    <option value={3}>Bold Italic</option>
                                </select>
                            </div>

                            <div style={{marginBottom: '10px'}}>
                                <label style={{display: 'block', marginBottom: '5px'}}>Font File:</label>
                                <select
                                    value={attr.fontFile}
                                    onChange={(e) => updateGadgetProperty(`attributes${groupIndex}:${attrIndex}.fontFile`, e.target.value)}
                                    style={{width: '100%', padding: '5px'}}
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
