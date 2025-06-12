import React, { useEffect, useState, useCallback } from "react";
import { offBackendEvent, onBackendEvent, ToPoint } from "./utils/wailsBridge";
import { AssociationProps, GadgetProps } from "./utils/Props";
import { GadgetPopup } from "./components/CreateGadgetPopup";
import { useBackendCanvasData } from "./hooks/useBackendCanvasData";
import { useGadgetUpdater } from "./hooks/useGadgetUpdater";
import { useAssociationUpdater } from "./hooks/useAssociationUpdater";
import AssociationPopup from "./components/AssociationPopup";
import LoadProjectPage from "./components/LoadProjectPage";
import DiagramPage from "./components/DiagramPage";
import DrawingCanvas from "./components/Canvas";
import ComponentPropertiesPanel from "./components/ComponentPropertiesPanel";
import TopMenu from "./components/TopMenu";
import { usePopupState, useAssPopupState } from "./hooks/usePopupState";
import { useDiagramActions } from "./hooks/useDiagramActions";

interface ProjectData {
    ProjectName: string;
    diagrams: string[];
    isNewEmptyProject?: boolean;
}

const App: React.FC = () => {
    const [currentView, setCurrentView] = useState<'load' | 'diagrams' | 'editor'>('load');
    const [projectData, setProjectData] = useState<ProjectData | null>(null);
    const [diagramName, setDiagramName] = useState<string | null>(null);
    const [selectedComponent, setSelectedComponent] = useState<GadgetProps | AssociationProps | null>(null);
    const [isAddingAssociation, setIsAddingAssociation] = useState(false);
    const [canvasBackgroundColor, setCanvasBackgroundColor] = useState<string>("#C2C2C2");

    // popup hooks
    const { showPopup, open: openGadgetPopup, close: closeGadgetPopup } = usePopupState();
    const { showAssPopup, open: openAssPopup, close: closeAssPopup, assStartPoint, setAssStartPoint, assEndPoint, setAssEndPoint } = useAssPopupState();

    // backend data
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
    // actions
    const { handleSaveProject, handleSaveDiagram, handleDiagramUndo, handleDiagramRedo, handleDeleteSelectedComponent} = useDiagramActions(reloadBackendData);

    // handler: project/diagram
    const handleProjectLoaded = (loadedProjectData: ProjectData) => {
        setProjectData(loadedProjectData);
        setCurrentView('diagrams');
    };
    const handleDiagramSelected = (diagramData: any) => {
        setCurrentView('editor');
        reloadBackendData();
    };
    const handleBackToDiagrams = () => setCurrentView('diagrams');
    const handleGetDiagramName = async () => {
        try {
            const { GetCurrentDiagramName } = await import("../wailsjs/go/umlproject/UMLProject");
            const name = await GetCurrentDiagramName();
            setDiagramName(name);
        } catch (error) {
            console.error("Error fetching diagram name:", error);
        }
    };

    // handler: association
    const handleAddAss = () => {
        setIsAddingAssociation(true);
        setAssStartPoint(null);
        setAssEndPoint(null);
        closeAssPopup();
    };
    const handleCanvasClick = async (point: { x: number, y: number }) => {
        if (isAddingAssociation) {
            if (!assStartPoint) {
                setAssStartPoint(point);
            } else if (!assEndPoint) {
                setAssEndPoint(point);
                openAssPopup();
            }
        }
    };
    const handleAssPopupAdd = async (assType: number) => {
        if (assStartPoint && assEndPoint) {
            const { StartAddAssociation, EndAddAssociation } = await import("../wailsjs/go/umlproject/UMLProject");
            await StartAddAssociation(ToPoint(assStartPoint.x, assStartPoint.y));
            await EndAddAssociation(assType, ToPoint(assEndPoint.x, assEndPoint.y));
            setIsAddingAssociation(false);
            setAssStartPoint(null);
            setAssEndPoint(null);
            closeAssPopup();
            reloadBackendData();
        }
    };
    const handleAssPopupClose = () => {
        setIsAddingAssociation(false);
        setAssStartPoint(null);
        setAssEndPoint(null);
        closeAssPopup();
    };

    // handler: selection/canvas
    const handleSelectionChange = useCallback((component: GadgetProps | AssociationProps | null) => {
        setSelectedComponent(prev => (prev !== component ? component : prev));
    }, []);
    const handleCanvasColorChange = (color: string) => setCanvasBackgroundColor(color);
    

    // Render
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
    // Editor view
    return (
        <div className="h-screen mx-auto px-4 bg-neutral-700">
            <TopMenu
                projectData={projectData}
                handleBackToDiagrams={handleBackToDiagrams}
                onGetDiagramName={handleGetDiagramName}
                onShowPopup={openGadgetPopup}
                onAddAss={handleAddAss}
                onCanvasColorChange={handleCanvasColorChange}
                onSaveProject={handleSaveProject}
                onSaveDiagram={handleSaveDiagram}
                diagramName={diagramName}
                canvasBackgroundColor={canvasBackgroundColor}
                onUndo={handleDiagramUndo}
                onRedo={handleDiagramRedo}
                onDeleteSelectedComponent={handleDeleteSelectedComponent}
            />
            {showPopup && (
                <GadgetPopup
                    isOpen={showPopup}
                    onCreate={() => closeGadgetPopup()}
                    onClose={closeGadgetPopup}
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
