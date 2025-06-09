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
import AssociationPopup from "./components/AssociationPopup";
import SessionBar from "./components/SessionBar";
import DiagramTabs from "./components/DiagramTabs";
import TopMenu from "./components/TopMenu";

const App: React.FC = () => {
    const [diagramName, setDiagramName] = useState<string | null>(null);
    const [showPopup, setShowPopup] = useState(false);
    const [selectedComponent, setSelectedComponent] = useState<GadgetProps | AssociationProps | null>(null);
    const [selectedGadgetCount, setSelectedGadgetCount] = useState<number>(0);
    const [isAddingAssociation, setIsAddingAssociation] = useState(false);
    const [showAssPopup, setShowAssPopup] = useState(false);
    const [assStartPoint, setAssStartPoint] = useState<{ x: number, y: number } | null>(null);
    const [assEndPoint, setAssEndPoint] = useState<{ x: number, y: number } | null>(null);
    const [sessionName, setSessionName] = useState<string | null>(null);
    const [isSessionConnected, setIsSessionConnected] = useState(false);
    const [showSessionModal, setShowSessionModal] = useState(false);
    const [diagramTabs, setDiagramTabs] = useState<string[]>(["DefaultDiagram"]);
    const [activeDiagram, setActiveDiagram] = useState<string>("DefaultDiagram");

    const {backendData, setBackendData, reloadBackendData} = useBackendCanvasData();

    const {handleUpdateGadgetProperty, handleAddAttributeToGadget} = useGadgetUpdater(
        selectedComponent as GadgetProps | null,
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
        // setIsAddingAssociation(true);
        // setAssStartPoint(null);
        // setAssEndPoint(null);
        // setShowAssPopup(false);
        // use mock-data for now

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
        setSelectedGadgetCount(count);
    };

    const handleJoinSession = () => {
        setShowSessionModal(true);
    };

    const handleLeaveSession = () => {
        setIsSessionConnected(false);
        setSessionName(null);
    };

    const handleSessionModalConfirm = (name: string) => {
        setSessionName(name);
        setIsSessionConnected(true);
        setShowSessionModal(false);
    };

    const handleSessionModalClose = () => {
        setShowSessionModal(false);
    };

    const handleAddDiagramTab = () => {
        let newName = "Diagram" + (diagramTabs.length + 1);
        let i = 1;
        while (diagramTabs.includes(newName)) {
            i++;
            newName = `Diagram${diagramTabs.length + i}`;
        }
        setDiagramTabs([...diagramTabs, newName]);
        setActiveDiagram(newName);
        setDiagramName(newName);
    };

    const handleSelectDiagramTab = (name: string) => {
        setActiveDiagram(name);
        setDiagramName(name);
        // 這裡可根據需求載入對應diagram資料
    };

    const handleCloseDiagramTab = (name: string) => {
        const idx = diagramTabs.indexOf(name);
        const newTabs = diagramTabs.filter(tab => tab !== name);
        setDiagramTabs(newTabs);
        if (activeDiagram === name) {
            const nextTab = newTabs[idx - 1] || newTabs[0] || null;
            setActiveDiagram(nextTab || "");
            setDiagramName(nextTab || null);
        }
    };

    // TopMenu handlers
    const handleOpenProject = () => {
        // TODO: 呼叫後端 Open Project API
        alert("[TODO] Open Project API");
    };
    const handleSave = () => {
        // TODO: 呼叫後端 Save API
        alert("[TODO] Save Project API");
    };
    const handleExport = () => {
        // TODO: 呼叫後端 Export API
        alert("[TODO] Export Project API");
    };
    const handleValidate = () => {
        // TODO: 呼叫後端 Validate API
        alert("[TODO] Validate Project API");
    };

    return (
        <div className="h-screen mx-auto px-4 bg-neutral-700">
            <TopMenu
                onOpenProject={handleOpenProject}
                onSave={handleSave}
                onExport={handleExport}
                onValidate={handleValidate}
                sessionName={sessionName}
                isConnected={isSessionConnected}
                onJoinSession={handleJoinSession}
                onLeaveSession={handleLeaveSession}
            />
            <DiagramTabs
                diagrams={diagramTabs.map(name => ({ name }))}
                activeDiagram={activeDiagram}
                onSelect={handleSelectDiagramTab}
                onClose={handleCloseDiagramTab}
                onAdd={handleAddDiagramTab}
            />
            <Toolbar
                onGetDiagramName={handleGetDiagramName}
                onShowPopup={() => setShowPopup(true)}
                onAddAss={handleAddAss}
                diagramName={diagramName}
            />
            {/* <SessionBar
                sessionName={sessionName}
                isConnected={isSessionConnected}
                onJoinSession={handleJoinSession}
                onLeaveSession={handleLeaveSession}
            /> */}
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
            {selectedComponent && (
                <ComponentPropertiesPanel
                    selectedComponent={selectedComponent}
                    updateComponentProperty={handleUpdateGadgetProperty}
                    addAttributeToComponent={handleAddAttributeToGadget}
                />
            )}
            {showSessionModal && (
                <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-40 z-50">
                    <div className="bg-white rounded shadow-lg p-6 w-80">
                        <h2 className="text-lg font-bold mb-4">加入/建立 Session</h2>
                        <input
                            id="session-input"
                            type="text"
                            className="border rounded px-2 py-1 w-full mb-4"
                            placeholder="輸入 Session 名稱"
                            onKeyDown={e => {
                                if (e.key === 'Enter') {
                                    handleSessionModalConfirm((e.target as HTMLInputElement).value);
                                }
                            }}
                        />
                        <div className="flex justify-end gap-2">
                            <button className="px-3 py-1 rounded bg-gray-300" onClick={handleSessionModalClose}>取消</button>
                            <button className="px-3 py-1 rounded bg-blue-600 text-white" onClick={() => {
                                const input = document.querySelector<HTMLInputElement>("#session-input");
                                if (input) handleSessionModalConfirm(input.value);
                            }}>確認</button>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
};

export default App;
