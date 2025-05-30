import React, {useEffect, useState} from "react";
import {offBackendEvent, onBackendEvent, ToPoint} from "./utils/wailsBridge";

import {
    AddAttributeToGadget,
    GetCurrentDiagramName,
    GetDrawData,
    SetAttrContentComp,
    SetAttrFontComp,
    SetAttrSizeComp,
    SetAttrStyleComp,
    SetColorComp,
    SetPointComp,
    SetSetLayerComp
} from "../wailsjs/go/umlproject/UMLProject";
import {mockSelfAssociationUp} from "./assets/mock/ass";

import {CanvasProps, GadgetProps, AssociationProps} from "./utils/Props";
import DrawingCanvas from "./components/Canvas";
import {GadgetPopup} from "./components/CreateGadgetPopup";
import Toolbar from "./components/Toolbar";
import GadgetPropertiesPanel from "./components/GadgetPropertiesPanel";
import {GetCurrentDiagramName} from "../wailsjs/go/umlproject/UMLProject";
import {useBackendCanvasData} from "./hooks/useBackendCanvasData";
import {useGadgetUpdater} from "./hooks/useGadgetUpdater";
import {StartAddAssociation, EndAddAssociation} from "../wailsjs/go/umlproject/UMLProject";
import AssociationPopup from "./components/AssociationPopup";

const App: React.FC = () => {
    const [diagramName, setDiagramName] = useState<string | null>(null);
    const [showPopup, setShowPopup] = useState(false);
    const [selectedGadget, setSelectedGadget] = useState<GadgetProps | null>(null);
    const [selectedComponent, setSelectedComponent] = useState<GadgetProps | AssociationProps | null>(null);
    const [selectedComponentType, setSelectedComponentType] = useState<"gadget" | "association" | null>(null);
    const [selectedGadgetCount, setSelectedGadgetCount] = useState<number>(0);
    const [isAddingAssociation, setIsAddingAssociation] = useState(false);
    const [showAssPopup, setShowAssPopup] = useState(false);
    const [assStartPoint, setAssStartPoint] = useState<{x: number, y: number} | null>(null);
    const [assEndPoint, setAssEndPoint] = useState<{x: number, y: number} | null>(null);

    const {backendData, setBackendData, reloadBackendData} = useBackendCanvasData();

    const {handleUpdateGadgetProperty, handleAddAttributeToGadget} = useGadgetUpdater(
        selectedGadget,
        backendData,
        reloadBackendData
    );

    const handleGetDiagramName = async () => {
        try {
            const name = await GetCurrentDiagramName();
            setDiagramName(name);
        } catch (error) {
            console.error("Error fetching diagram name:", error);
        }
    };

    const handleAddAss = () => {
        setIsAddingAssociation(true);
        setAssStartPoint(null);
        setAssEndPoint(null);
        setShowAssPopup(false);
    };
    // new stuff
    // Canvas 點擊事件的 callback
    const handleCanvasClick = async (point: {x: number, y: number}) => {
        if (isAddingAssociation) {
            if (!assStartPoint) {
                setAssStartPoint(point);
            } else if (!assEndPoint) {
                setAssEndPoint(point);
                setShowAssPopup(true);
            }
        }
    };

    const handleAssPopupAdd = async (assType: number) => {
        if (assStartPoint && assEndPoint) {
            await StartAddAssociation(ToPoint(assStartPoint.x, assStartPoint.y));
            await EndAddAssociation(assType, ToPoint(assEndPoint.x, assEndPoint.y));
            setIsAddingAssociation(false);
            setAssStartPoint(null);
            setAssEndPoint(null);
            setShowAssPopup(false);
            reloadBackendData();
        }
    };

    // deleted
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

    const handleAssPopupClose = () => {
        setIsAddingAssociation(false);
        setAssStartPoint(null);
        setAssEndPoint(null);
        setShowAssPopup(false);
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
                            SetAttrContentComp,
                            i, j, value,
                            "Gadget content changed",
                            "Error editing gadget content"
                        );
                    }
                    if (childProp === 'fontSize') {
                        setTripleValue(
                            SetAttrSizeComp,
                            i, j, value,
                            "Gadget fontSize changed",
                            "Error editing gadget fontSize"
                        );
                    }
                    // new stuff
                    if (childProp === 'fontStyleB') {
                        setTripleValue(
                            SetAttrStyleComp,
                            i, j, value,
                            "Gadget font style (bold) changed",
                            "Error editing gadget font style (bold)"
                        );
                    }
                    if (childProp === 'fontStyleI') {
                        setTripleValue(
                            SetAttrStyleComp,
                            i, j, value,
                            "Gadget font style (italic) changed",
                            "Error editing gadget font style (italic)"
                        );
                    }
                    if (childProp === 'fontStyleU') {
                        setTripleValue(
                            SetAttrStyleComp,
                            i, j, value,
                            "Gadget font style (underline) changed",
                            "Error editing gadget font style (underline)"
                        );
                    }
                    // new stuff ends
                }
            }
        } else {
            if (property === "x") {
                setSingleValue(
                    (val) => SetPointComp(ToPoint(val, selectedGadget.y)),
                    value,
                    "Gadget moved",
                    "Error editing gadget"
                );
            }
            if (property === "y") {
                setSingleValue(
                    (val) => SetPointComp(ToPoint(selectedGadget.x, val)),
                    value,
                    "Gadget moved",
                    "Error editing gadget"
                );
            }
            if (property === "layer") {
                setSingleValue(
                    SetSetLayerComp,
                    value,
                    "layer changed",
                    "Error editing gadget"
                );
            }
            if (property === "color") {
                setSingleValue(
                    SetColorComp,
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
                reloadBackendData={reloadBackendData}
                onSelectionChange={handleSelectionChange}
                onCanvasClick={handleCanvasClick}
                isAddingAssociation={isAddingAssociation}
            />
            {showAssPopup && assStartPoint && assEndPoint && (
                <AssociationPopup
                    isOpen={showAssPopup}
                    startPoint={assStartPoint}
                    endPoint={assEndPoint}
                    onAdd={handleAssPopupAdd}
                    onClose={handleAssPopupClose}
                />
            )}
            {selectedGadgetCount === 1 && (
                <GadgetPropertiesPanel
                    selectedGadget={selectedGadget}
                    updateGadgetProperty={handleUpdateGadgetProperty}
                    addAttributeToGadget={handleAddAttributeToGadget}
                    // 可加上 updateAssociationProperty, addAttributeToAssociation
                />
            )}
        </div>
    );
};

export default App;
