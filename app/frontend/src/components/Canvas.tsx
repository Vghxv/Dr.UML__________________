import React, { useEffect, useRef } from 'react';
import { CanvasProps, GadgetProps } from '../utils/Props';
import { createGadget } from '../utils/createGadget';
import { createAss } from '../utils/createAssociation';
import { useCanvasMouseEvents } from '../hooks/useCanvasMouseEvents';
import { useSelection } from '../hooks/useSelection';
import GadgetPropertiesPanel from './GadgetPropertiesPanel';
import {ToPoint} from '../utils/wailsBridge'
import {
    SelectComponent,
    SetColorGadget,
    SetPointGadget,
    SetSetLayerGadget,
    SetAttrContentGadget,
    SetAttrSizeGadget,
    SetAttrStyleGadget,
    GetDrawData
} from "../../wailsjs/go/umlproject/UMLProject";

const DrawingCanvas: React.FC<{ backendData: CanvasProps | null, reloadBackendData?: () => void }> = ({backendData, reloadBackendData}) => {
    const canvasRef = useRef<HTMLCanvasElement>(null);
    const { selectedGadgetCount, selectedGadget } = useSelection(backendData?.gadgets);

    const redrawCanvas = () => {
        const canvas = canvasRef.current;
        if (canvas) {
            const ctx = canvas.getContext('2d');
            if (ctx) {
                ctx.clearRect(0, 0, canvas.width, canvas.height);

                backendData?.gadgets?.forEach((gadget: GadgetProps) => {
                    const gad = createGadget("Class", gadget, backendData.margin);
                    gad.draw(ctx, backendData.margin, backendData.lineWidth);
                });

                backendData?.Association?.forEach((association) => {
                    const ass = createAss("Association", association, backendData.margin);
                    ass.draw(ctx, backendData.margin, backendData.lineWidth);
                });

                // 更新選取狀態（確保每次都用最新 gadgets）
                // updateSelection(backendData?.gadgets);
            }
        }
    };

    useEffect(() => {
        redrawCanvas();
    }, [backendData]);

    const { handleMouseDown, handleMouseMove, handleMouseUp } = useCanvasMouseEvents(
        canvasRef,
        () => {
            if (reloadBackendData) {
                reloadBackendData();
            }
        }
    );

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
            {selectedGadgetCount === 1 && (
                <GadgetPropertiesPanel
                    selectedGadget={selectedGadget}
                    updateGadgetProperty={updateGadgetProperty}
                />
            )}
        </div>
    );
};

export default DrawingCanvas;
