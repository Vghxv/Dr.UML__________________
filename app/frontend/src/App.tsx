import React, {useState} from "react";
import {offBackendEvent, onBackendEvent, ToPoint} from "./utils/wailsBridge";
import {mockSelfAssociationUp} from "./assets/mock/ass";

import {CanvasProps, GadgetProps} from "./utils/Props";
import DrawingCanvas from "./components/Canvas";
import {GadgetPopup} from "./components/CreateGadgetPopup";
import Toolbar from "./components/Toolbar";
import GadgetPropertiesPanel from "./components/GadgetPropertiesPanel";
import {GetCurrentDiagramName} from "../wailsjs/go/umlproject/UMLProject";
import {useBackendCanvasData} from "./hooks/useBackendCanvasData";
import {useGadgetUpdater} from "./hooks/useGadgetUpdater";
import {StartAddAssociation, EndAddAssociation} from "../wailsjs/go/umlproject/UMLProject";
import { component } from "../wailsjs/go/models";

const App: React.FC = () => {
    const [diagramName, setDiagramName] = useState<string | null>(null);
    const [showPopup, setShowPopup] = useState(false);
    const [selectedGadget, setSelectedGadget] = useState<GadgetProps | null>(null);
    const [selectedGadgetCount, setSelectedGadgetCount] = useState<number>(0);
    const [isAddingAssociation, setIsAddingAssociation] = useState(false);
    const [assStartPoint, setAssStartPoint] = useState<{x: number, y: number} | null>(null);

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
    };

    // Canvas 點擊事件的 callback
    const handleCanvasClick = async (point: {x: number, y: number}) => {
        if (isAddingAssociation) {
            if (!assStartPoint) {
                // 第一次點擊，設為起點
                setAssStartPoint(point);
                await StartAddAssociation(ToPoint(point.x, point.y));
            } else {
                // 第二次點擊，設為終點並創建 Association
                await EndAddAssociation(component.AssociationType.Dependency, ToPoint(point.x, point.y));
                setIsAddingAssociation(false);
                setAssStartPoint(null);
                reloadBackendData();
            }
        }
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
