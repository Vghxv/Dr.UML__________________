import React, {useEffect, useState} from "react";
import {offBackendEvent, onBackendEvent, ToPoint} from "./utils/wailsBridge";

import {
    EndAddAssociation,
    GetCurrentDiagramName,
    GetDrawData,
    StartAddAssociation
} from "../wailsjs/go/umlproject/UMLProject";

import {AssociationProps, CanvasProps, GadgetProps} from "./utils/Props";
import DrawingCanvas from "./components/Canvas";
import {GadgetPopup} from "./components/CreateGadgetPopup";
import Toolbar from "./components/Toolbar";
import ComponentPropertiesPanel from "./components/ComponentPropertiesPanel";
import {useBackendCanvasData} from "./hooks/useBackendCanvasData";
import {useGadgetUpdater} from "./hooks/useGadgetUpdater";
import {useAssociationUpdater} from "./hooks/useAssociationUpdater";
import AssociationPopup from "./components/AssociationPopup";

const App: React.FC = () => {
    const [diagramName, setDiagramName] = useState<string | null>(null);
    const [showPopup, setShowPopup] = useState(false);
    const [selectedComponent, setSelectedComponent] = useState<GadgetProps | AssociationProps | null>(null);
    // const [selectedGadgetCount, setSelectedGadgetCount] = useState<number>(0);
    const [isAddingAssociation, setIsAddingAssociation] = useState(false);
    const [showAssPopup, setShowAssPopup] = useState(false);
    const [assStartPoint, setAssStartPoint] = useState<{ x: number, y: number } | null>(null);
    const [assEndPoint, setAssEndPoint] = useState<{ x: number, y: number } | null>(null);

    const {backendData, reloadBackendData} = useBackendCanvasData();

    const {handleUpdateGadgetProperty, handleAddAttributeToGadget} = useGadgetUpdater(
        selectedComponent as GadgetProps | null,
        backendData,
        reloadBackendData
    );

    const {handleUpdateAssociationProperty, handleAddAttributeToAssociation} = useAssociationUpdater(
        selectedComponent as AssociationProps | null,
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

    const handleSelectionChange = (component: GadgetProps| AssociationProps | null, count: number) => {
        setSelectedComponent(component);
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
            {/* TODO generalize updateProperty and addAttributeToXXX */}
            {selectedComponent && (
                <ComponentPropertiesPanel
                    selectedComponent={selectedComponent}
                    updateGadgetProperty={handleUpdateGadgetProperty}
                    updateAssociationProperty={handleUpdateAssociationProperty}
                    addAttributeToGadget={handleAddAttributeToGadget}
                    addAttributeToAssociation={handleAddAttributeToAssociation}
                />
            )}
        </div>
    );
};

export default App;
