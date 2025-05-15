import React, {useEffect, useRef, useState} from 'react';
import {CanvasProps, GadgetProps} from '../utils/Props';
import {createGadget} from '../utils/createGadget';
import {ToPoint} from '../utils/wailsBridge'
import {
    SelectComponent,
    SetColorGadget,
    SetPointGadget,
    SetSetLayerGadget,
    SetAttrContentGadget,
    SetAttrSizeGadget,
    SetAttrStyleGadget
} from "../../wailsjs/go/umlproject/UMLProject";

const DrawingCanvas: React.FC<{ backendData: CanvasProps | null }> = ({backendData}) => {
    const canvasRef = useRef<HTMLCanvasElement>(null);
    const [selectedGadgetCount, setSelectedGadgetCount] = useState<number>(0);
    const [selectedGadget, setSelectedGadget] = useState<GadgetProps | null>(null);

    const redrawCanvas = () => {
        const canvas = canvasRef.current;
        if (canvas) {
            const ctx = canvas.getContext('2d');
            if (ctx) {
                ctx.clearRect(0, 0, canvas.width, canvas.height);
                let selectedCount = 0;
                let selectedGad: GadgetProps | null = null;

                backendData?.gadgets?.forEach((gadget: GadgetProps) => {
                    const gad = createGadget("Class", gadget, backendData.margin);
                    gad.draw(ctx, backendData.margin, backendData.lineWidth);

                    if (gadget.isSelected) {
                        selectedCount++;
                        selectedGad = gadget;
                    }
                });

                setSelectedGadgetCount(selectedCount);
                setSelectedGadget(selectedCount === 1 ? selectedGad : null);
            }
        }
    }

    useEffect(() => {
        redrawCanvas();
    }, [backendData]);

    const handleMouseDown = (event: React.MouseEvent<HTMLCanvasElement>) => {
        console.log("handleMouseDown");
        const rect = canvasRef.current?.getBoundingClientRect();
        const x = Math.round(event.clientX - (rect?.left || 0));
        const y = Math.round(event.clientY - (rect?.top || 0));
        console.log("Mouse down at:", x, y);
        SelectComponent(ToPoint(
            x,
            y
        )).then(() => {
                console.log("handleMouseDown");
            }
        ).catch((error) => {
                console.error("Error selecting component:", error);
            }
        );
    };

    const handleMouseMove = (event: React.MouseEvent<HTMLCanvasElement>) => {
        // TODO: add some hover things
    };

    const handleMouseUp = () => {
    };

    const updateGadgetProperty = (property: string, value: any) => {
        if (!selectedGadget || !backendData || !backendData.gadgets) return;

        // Handle nested properties like attributes[0][0].content
        if (property.includes('.')) {
            const [parentProp, childProp] = property.split('.');
            if (parentProp.startsWith('attributes')) {
                // Parse indices from string like 'attributes[0][0]'
                const matches = parentProp.match(/attributes(\d+):(\d+)/);
                if (matches && matches.length === 3) {
                    const i = parseInt(matches[1]);
                    const j = parseInt(matches[2]);

                    // console.log(i, j, childProp);
                    if (childProp === 'content') {
                        SetAttrContentGadget(i, j, value).then(
                            () => {
                                console.log("Gadget content changed");
                            }
                        ).catch((error) => {
                                console.error("Error changing gadget content:", error);
                            }
                        );
                    }
                    if(childProp === 'fontSize') {
                        SetAttrSizeGadget(i, j, value).then(
                            () => {
                                console.log("Gadget fontSize changed");
                            }
                        ).catch((error) => {
                                console.error("Error changing gadget fontSize:", error);
                            }
                        );
                    }
                    if(childProp === 'fontStyle') {
                        SetAttrStyleGadget(i, j, value).then(
                            () => {
                                console.log("Gadget fontStyle changed");
                            }
                        ).catch((error) => {
                                console.error("Error changing gadget fontStyle:", error);
                            }
                        );
                    }
                }
            }
        } else {
            if (property === "x") {
                SetPointGadget(ToPoint(value, selectedGadget.y)).then(
                    () => {
                        console.log("Gadget moved");
                    }
                ).catch((error) => {
                        console.error("Error moving gadget:", error);
                    }
                );
            }
            if (property === "y") {
                SetPointGadget(ToPoint(selectedGadget.x, value)).then(
                    () => {
                        console.log("Gadget moved");
                    }
                ).catch((error) => {
                        console.error("Error moving gadget:", error);
                    }
                );
            }
            if (property === "layer") {
                SetSetLayerGadget(value).then(
                    () => {
                        console.log("layer changed");
                    }
                ).catch((error) => {
                        console.error("Error moving gadget:", error);
                    }
                );
            }
            if (property === "color") {
                SetColorGadget(value).then(
                    () => {
                        console.log("color changed");
                    }
                ).catch((error) => {
                        console.error("Error moving gadget:", error);
                    }
                );
            }
        }

        // updatedGadgets[gadgetIndex] = updatedGadget;
        // updatedBackendData.gadgets = updatedGadgets;

        // Here you would typically call an API to update the backend
        // For now, we'll just redraw the canvas with the updated data
        // This is a placeholder for actual backend update logic
        // console.log('Updated gadget:', updatedGadget);

        // Update the selected gadget state
        // setSelectedGadget(updatedGadget);

        // Redraw canvas with updated data
        // In a real implementation, you would update the backendData through proper channels
        // This is just for demonstration
        redrawCanvas();
    };

    // Side panel component for editing gadget properties
    const GadgetPropertiesPanel = () => {
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
                                        {/*TODO: make this part work */}
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

    return (
        <div style={{position: 'relative', display: 'flex'}}>
            <canvas
                ref={canvasRef}
                style={{
                    border: '2px solid #444',
                    borderRadius: '8px',
                    backgroundColor: '#1e1e1e',
                    boxShadow: '0 4px 8px rgba(0, 0, 0, 0.5)',
                    margin: '20px auto',
                    position: 'relative',
                }}
                width="1024"
                height="768"
                onMouseDown={handleMouseDown}
                onMouseMove={handleMouseMove}
                onMouseUp={handleMouseUp}
            />
            {selectedGadgetCount === 1 && <GadgetPropertiesPanel/>}
        </div>
    );
};

export default DrawingCanvas;
