import React, {useEffect, useState} from "react";
import {offBackendEvent, onBackendEvent, ToPoint} from "./utils/wailsBridge";

import {
    AddAttributeToGadget,
    GetCurrentDiagramName,
    GetDrawData,
    SetAttrContentGadget,
    SetAttrSizeGadget,
    SetAttrStyleGadget,
    SetColorGadget,
    SetPointGadget,
    SetSetLayerGadget
} from "../wailsjs/go/umlproject/UMLProject";
import {mockSelfAssociationUp} from "./assets/mock/ass";

import {CanvasProps, GadgetProps} from "./utils/Props";
import DrawingCanvas from "./components/Canvas";
// import mockData from './assets/mock/gadget';
import {GadgetPopup} from "./components/CreateGadgetPopup";
import Toolbar from "./components/Toolbar";
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

    const handleAddAss = async () => {
        setBackendData(prev => {
            if (!prev) return prev;
            const newAss = prev.Association ? [...prev.Association, mockSelfAssociationUp] : [mockSelfAssociationUp];
            return {...prev, Association: newAss};
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
    // Helper function to handle setting a single value with a promise
    const setSingleValue = (
        apiFunction: (value: any) => Promise<any>,
        value: any,
        successMessage: string,
        errorPrefix: string
    ) => {
        apiFunction(value).then(
            () => {
                console.log(successMessage);
                loadCanvasData().then(
                    r => console.log("Loaded canvas data:", r)
                )
            }
        ).catch((error) => {
                console.error(`${errorPrefix}:`, error);
            }
        );
    };

    // Helper function to handle setting a value with three parameters
    const setTripleValue = (
        apiFunction: (param1: any, param2: any, param3: any) => Promise<any>,
        param1: any,
        param2: any,
        param3: any,
        successMessage: string,
        errorPrefix: string
    ) => {
        apiFunction(param1, param2, param3).then(
            () => {
                console.log(successMessage);
                loadCanvasData().then(
                    r => console.log("Loaded canvas data:", r)
                )
            }
        ).catch((error) => {
                console.error(`${errorPrefix}:`, error);
            }
        );
    };

    const handleAddAttributeToGadget = (section: number, content: string) => {
        if (!selectedGadget || !backendData || !backendData.gadgets) return;

        console.log(`Adding attribute to section ${section} with content ${content}`);

        setSingleValue(
            (val) => AddAttributeToGadget(section, val),
            content,
            "Attribute added",
            "Error adding attribute"
        );
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

                    if (childProp === 'content') {
                        setTripleValue(
                            SetAttrContentGadget,
                            i, j, value,
                            "Gadget content changed",
                            "Error editing gadget content"
                        );
                    }
                    if (childProp === 'fontSize') {
                        setTripleValue(
                            SetAttrSizeGadget,
                            i, j, value,
                            "Gadget fontSize changed",
                            "Error editing gadget fontSize"
                        );
                    }
                    if (childProp === 'fontStyle') {
                        setTripleValue(
                            SetAttrStyleGadget,
                            i, j, value,
                            "Gadget fontStyle changed",
                            "Error editing gadget fontStyle"
                        );
                    }
                }
            }
        } else {
            if (property === "x") {
                setSingleValue(
                    (val) => SetPointGadget(ToPoint(val, selectedGadget.y)),
                    value,
                    "Gadget moved",
                    "Error editing gadget"
                );
            }
            if (property === "y") {
                setSingleValue(
                    (val) => SetPointGadget(ToPoint(selectedGadget.x, val)),
                    value,
                    "Gadget moved",
                    "Error editing gadget"
                );
            }
            if (property === "layer") {
                setSingleValue(
                    SetSetLayerGadget,
                    value,
                    "layer changed",
                    "Error editing gadget"
                );
            }
            if (property === "color") {
                setSingleValue(
                    SetColorGadget,
                    value,
                    "color changed",
                    "Error editing gadget"
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
        <div className="h-screen mx-auto px-4 bg-neutral-700">
            <h1 className="text-3xl text-center font-bold text-white mb-4">Dr.UML</h1>
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
            />
            {selectedGadgetCount === 1 && (
                <GadgetPropertiesPanel
                    selectedGadget={selectedGadget}
                    updateGadgetProperty={handleUpdateGadgetProperty}
                    addAttributeToGadget={handleAddAttributeToGadget}
                />
            )}
        </div>
    );
};

export default App;
