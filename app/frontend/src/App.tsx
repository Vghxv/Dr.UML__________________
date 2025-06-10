import React, { useEffect, useState, useCallback } from "react";
import { offBackendEvent, onBackendEvent, ToPoint } from "./utils/wailsBridge";

import {
    EndAddAssociation,
    GetCurrentDiagramName,
    SaveDiagram,
    SaveDiagramFileDialog,
    SaveFileDialog,
    SaveProject,
    StartAddAssociation
} from "../wailsjs/go/umlproject/UMLProject";

import { AssociationProps, CanvasProps, GadgetProps } from "./utils/Props";
import { GadgetPopup } from "./components/CreateGadgetPopup";
import { useBackendCanvasData } from "./hooks/useBackendCanvasData";
import { useGadgetUpdater } from "./hooks/useGadgetUpdater";
import { useAssociationUpdater } from "./hooks/useAssociationUpdater";
import AssociationPopup from "./components/AssociationPopup";
import LoadProjectPage from "./components/LoadProjectPage";
import DiagramPage from "./components/DiagramPage";

interface ProjectData {
    ProjectName: string;
    diagrams: string[];
    isNewEmptyProject?: boolean;
}
import DrawingCanvas from "./components/Canvas";
import Toolbar from "./components/Toolbar";
import ComponentPropertiesPanel from "./components/ComponentPropertiesPanel";
import TopMenu from "./components/TopMenu";


const App: React.FC = () => {
    const [currentView, setCurrentView] = useState<'load' | 'diagrams' | 'editor'>('load');
    const [projectData, setProjectData] = useState<ProjectData | null>(null);
    const [diagramName, setDiagramName] = useState<string | null>(null);
    const [showPopup, setShowPopup] = useState(false);
    const [selectedComponent, setSelectedComponent] = useState<GadgetProps | AssociationProps | null>(null);
    const [isAddingAssociation, setIsAddingAssociation] = useState(false);
    const [showAssPopup, setShowAssPopup] = useState(false);
    const [assStartPoint, setAssStartPoint] = useState<{ x: number, y: number } | null>(null);
    const [assEndPoint, setAssEndPoint] = useState<{ x: number, y: number } | null>(null);
    const [canvasBackgroundColor, setCanvasBackgroundColor] = useState<string>("#C2C2C2");

    const { backendData, reloadBackendData } = useBackendCanvasData();

    const { handleUpdateGadgetProperty, handleAddAttributeToGadget } = useGadgetUpdater(
        selectedComponent as GadgetProps | null,
        backendData,
        reloadBackendData
    );

    const { handleUpdateAssociationProperty, handleAddAttributeToAssociation } = useAssociationUpdater(
        selectedComponent as AssociationProps | null,
        backendData,
        reloadBackendData
    );

    // Handle project loading from LoadProject component
    const handleProjectLoaded = (loadedProjectData: ProjectData) => {
        setProjectData(loadedProjectData);
        setCurrentView('diagrams');
    };

    // Handle diagram selection from DiagramPage component
    const handleDiagramSelected = (diagramData: any) => {
        console.log('Diagram data received:', diagramData);
        setCurrentView('editor');
        reloadBackendData();
    };

    // Handle going back to DiagramPage
    const handleBackToDiagrams = () => {
        setCurrentView('diagrams');
    };

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

    const handleSelectionChange = useCallback((component: GadgetProps | AssociationProps | null) => {
        setSelectedComponent(prev => {
            // Only update if the selection actually changed
            if (prev !== component) {
                return component;
            }
            return prev;
        });
    }, []);

    const handleCanvasColorChange = (color: string) => {
        setCanvasBackgroundColor(color);
    };

    const handleSaveProject = async () => {
        try {
            // Open file dialog to get save location
            const filePath = await SaveFileDialog();
            if (filePath) {
                // Save the project to the selected file path
                await SaveProject(filePath);
                console.log("Project saved successfully to:", filePath);
                // You could add a success notification here
            }
        } catch (error) {
            console.error("Error saving project:", error);
            // You could add an error notification here
        }
    };

    const handleSaveDiagram = async () => {
        try {
            // Open file dialog to get save location for diagram
            const filePath = await SaveDiagramFileDialog();
            if (filePath) {
                // Save the current diagram to the selected file path
                await SaveDiagram(filePath);
                console.log("Diagram saved successfully to:", filePath);
                // You could add a success notification here
            }
        } catch (error) {
            console.error("Error saving diagram:", error);
            // You could add an error notification here
        }
    };

    // Render different views based on current state
    if (currentView === 'load') {
        return <LoadProjectPage onProjectLoaded={handleProjectLoaded} />;
    }

    if (currentView === 'diagrams' && projectData) {
        return (
            <DiagramPage 
                projectData={projectData}
                onDiagramSelected={handleDiagramSelected}
                isNewEmptyProject={projectData.isNewEmptyProject}
            />
        );
    }

    // Editor view (current main application)
    return (
        <div className="h-screen mx-auto px-4 bg-neutral-700">
            <div className="flex items-center justify-between mb-4">
                <h1 className="text-3xl text-center font-bold text-white">Dr.UML</h1>
                {projectData && (
                    <div className="flex gap-2">
                        <button
                            onClick={handleBackToDiagrams}
                            className="bg-neutral-600 hover:bg-neutral-500 text-white px-3 py-1 rounded text-sm transition-colors"
                        >
                            Back to Diagrams
                        </button>
                    </div>
                )}
            </div>
            <TopMenu />
            <Toolbar
                onGetDiagramName={handleGetDiagramName}
                onShowPopup={() => setShowPopup(true)}
                onAddAss={handleAddAss}
                onCanvasColorChange={handleCanvasColorChange}
                onSaveProject={handleSaveProject}
                onSaveDiagram={handleSaveDiagram}
                diagramName={diagramName}
                canvasBackgroundColor={canvasBackgroundColor}
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
            {showAssPopup && assStartPoint && assEndPoint && (
                <AssociationPopup
                    isOpen={showAssPopup}
                    startPoint={assStartPoint}
                    endPoint={assEndPoint}
                    onAdd={handleAssPopupAdd}
                onClose={handleAssPopupClose}
                />
            )}
            <DrawingCanvas
                backendData={backendData}
                reloadBackendData={reloadBackendData}
                onSelectionChange={handleSelectionChange}
                onCanvasClick={handleCanvasClick}
                isAddingAssociation={isAddingAssociation}
                canvasBackgroundColor={canvasBackgroundColor}
            />

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
