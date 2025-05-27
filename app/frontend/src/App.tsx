import React, {useEffect, useState} from "react";
import {offBackendEvent, onBackendEvent, ToPoint} from "./utils/wailsBridge";

import {
    AddGadget,
    GetCurrentDiagramName,
    GetDrawData,
    SetColorGadget,
    SetPointGadget,
    SetSetLayerGadget,
    SetAttrContentGadget,
    SetAttrSizeGadget,
    SetAttrStyleGadget
} from "../wailsjs/go/umlproject/UMLProject";
import { mockAssociation, mockSelfAssociation, mockHorizontalAssociation, mockVerticalAssociation, mockSelfAssociationLeft, mockSelfAssociationUp} from "./assets/mock/ass";

import {CanvasProps, GadgetProps} from "./utils/Props";
import DrawingCanvas from "./components/Canvas";
// import mockData from './assets/mock/gadget';
import {GadgetPopup} from "./components/CreateGadgetPopup";
import Toolbar from "./components/Toolbar";
import { createAss } from "./utils/createAssociation";
import GadgetPropertiesPanel from "./components/GadgetPropertiesPanel";

const App: React.FC = () => {
    const [diagramName, setDiagramName] = useState<string | null>(null);
    const [backendData, setBackendData] = useState<CanvasProps | null>(null);
    const [showPopup, setShowPopup] = useState(false);
    const [selectedGadget, setSelectedGadget] = useState<GadgetProps | null>(null);
    const [selectedGadgetCount, setSelectedGadgetCount] = useState<number>(0);

    const handleGetDiagramName = async () => {
        try {
            const name = await GetCurrentDiagramName();
            setDiagramName(name);
        } catch (error) {
            console.error("Error fetching diagram name:", error);
        }
    };

    // 新增 addAss handler
    const handleAddAss = async () => {
        // 直接將 mockAssociation 加入 backendData.Association
        setBackendData(prev => {
            if (!prev) return prev;
            const newAss = prev.Association ? [...prev.Association, mockSelfAssociationUp] : [mockSelfAssociationUp];
            return { ...prev, Association: newAss };
        });
    };

    const loadCanvasData = async () => {
        try {
            const diagram = await GetDrawData();
            const canvasData: CanvasProps = {
                margin: diagram.margin,
                color: diagram.color,
                lineWidth: diagram.lineWidth,
                gadgets: diagram.gadgets.map(gadget => ({
                    gadgetType: gadget.gadgetType.toString(),
                    x: gadget.x,
                    y: gadget.y,
                    layer: gadget.layer,
                    height: gadget.height,
                    width: gadget.width,
                    color: gadget.color,
                    isSelected: gadget.isSelected,
                    attributes: gadget.attributes
                }))
            };
            setBackendData(canvasData);
        } catch (error) {
            console.error("Error loading canvas data:", error);
        }
    };

    const handleSelectionChange = (gadget: GadgetProps | null, count: number) => {
        setSelectedGadget(gadget);
        setSelectedGadgetCount(count);
    };

    const handleUpdateGadgetProperty = (property: string, value: any) => {
        if (!selectedGadget || !backendData || !backendData.gadgets) return;

        console.log(`Updating property ${property} to ${value}`);

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
                                loadCanvasData();
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
                                loadCanvasData();
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
                                loadCanvasData();
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
                        loadCanvasData();
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
                        loadCanvasData();
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
                        loadCanvasData();
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
                        loadCanvasData();
                    }
                ).catch((error) => {
                        console.error("Error moving gadget:", error);
                    }
                );
            }
        }
    };

    useEffect(() => {
        // Load canvas data when the component mounts
        loadCanvasData().then(r => console.log("Loaded canvas data:", r));

        onBackendEvent("backend-event", (result) => {
            console.log("Received data from backend:", result);
            // Convert the received data to CanvasProps format
            if (result) {
                const canvasData: CanvasProps = {
                    margin: result.margin,
                    color: result.color,
                    lineWidth: result.lineWidth,
                    gadgets: result.gadgets?.map((gadget: {
                        gadgetType: { toString: () => any; };
                        x: any;
                        y: any;
                        layer: any;
                        height: any;
                        width: any;
                        color: any;
                        isSelected: any;
                        attributes: any;
                    }) => ({
                        gadgetType: gadget.gadgetType.toString(),
                        x: gadget.x,
                        y: gadget.y,
                        layer: gadget.layer,
                        height: gadget.height,
                        width: gadget.width,
                        color: gadget.color,
                        isSelected: gadget.isSelected,
                        attributes: gadget.attributes
                    }))
                };
                setBackendData(canvasData);
            }
        });

        return () => {
            offBackendEvent("backend-event");
        };
    }, []);

    return (
        <div>
            <h1 style={{fontFamily: "Inkfree"}}>Dr.UML</h1>
            <Toolbar
                onGetDiagramName={handleGetDiagramName}
                onShowPopup={() => setShowPopup(true)}
                onAddAss={handleAddAss}
                diagramName={diagramName}
            />
            {showPopup && (
                <GadgetPopup
                    isOpen={showPopup}
                    onCreate={(gadget) => {
                        console.log("New Gadget Created:", gadget);
                        setShowPopup(false);
                    }}
                    onClose={() => setShowPopup(false)}
                />
            )}
            <DrawingCanvas
                backendData={backendData}
                reloadBackendData={loadCanvasData}
                onSelectionChange={handleSelectionChange}
                onUpdateGadgetProperty={handleUpdateGadgetProperty}
            />
            {selectedGadgetCount === 1 && (
                <GadgetPropertiesPanel
                    selectedGadget={selectedGadget}
                    updateGadgetProperty={handleUpdateGadgetProperty}
                />
            )}
        </div>
    );
};

export default App;
