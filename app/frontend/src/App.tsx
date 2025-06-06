import React, {useEffect, useState} from "react";
import {offBackendEvent, onBackendEvent, ToPoint} from "./utils/wailsBridge";

import {
    EndAddAssociation,
    GetCurrentDiagramName,
    GetDrawData,
    StartAddAssociation
} from "../wailsjs/go/umlproject/UMLProject";

import {CanvasProps, GadgetProps} from "./utils/Props";
import DrawingCanvas from "./components/Canvas";
import {GadgetPopup} from "./components/CreateGadgetPopup";
import Toolbar from "./components/Toolbar";
import GadgetPropertiesPanel from "./components/GadgetPropertiesPanel";
import {useBackendCanvasData} from "./hooks/useBackendCanvasData";
import {useGadgetUpdater} from "./hooks/useGadgetUpdater";
import AssociationPopup from "./components/AssociationPopup";

const App: React.FC = () => {
    const [diagramName, setDiagramName] = useState<string | null>(null);
    const [showPopup, setShowPopup] = useState(false);
    const [selectedGadget, setSelectedGadget] = useState<GadgetProps | null>(null);
    const [selectedGadgetCount, setSelectedGadgetCount] = useState<number>(0);
    const [isAddingAssociation, setIsAddingAssociation] = useState(false);
    const [showAssPopup, setShowAssPopup] = useState(false);
    const [assStartPoint, setAssStartPoint] = useState<{ x: number, y: number } | null>(null);
    const [assEndPoint, setAssEndPoint] = useState<{ x: number, y: number } | null>(null);

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

    const handleCanvasClick = async (point: { x: number, y: number }) => {
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
