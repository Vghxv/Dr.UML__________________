import { SaveFileDialog, SaveProject, SaveDiagramFileDialog, SaveDiagram, UndoDiagramChange, RedoDiagramChange, RemoveSelectedComponents } from "../../wailsjs/go/umlproject/UMLProject";

export function useDiagramActions(reloadBackendData: () => void) {
    const handleSaveProject = async () => {
        try {
            const filePath = await SaveFileDialog();
            if (filePath) {
                await SaveProject(filePath);
            }
        } catch (error) {
            console.error("Error saving project:", error);
        }
    };

    const handleSaveDiagram = async () => {
        try {
            const filePath = await SaveDiagramFileDialog();
            if (filePath) {
                await SaveDiagram(filePath);
            }
        } catch (error) {
            console.error("Error saving diagram:", error);
        }
    };

    const handleDiagramUndo = async () => {
        try {
            await UndoDiagramChange();
            reloadBackendData();
        } catch (error) {
            console.error("Error performing undo:", error);
        }
    };

    const handleDiagramRedo = async () => {
        try {
            await RedoDiagramChange();
            reloadBackendData();
        } catch (error) {
            console.error("Error performing redo:", error);
        }
    };

    const handleDeleteSelectedComponent = async () => {
        try {
            console.log("deleting selected component ");
            await RemoveSelectedComponents();
            reloadBackendData();
        } catch (error) {
            console.error("Error deleting component:", error);
        }
    };

    return {
        handleSaveProject,
        handleSaveDiagram,
        handleDiagramUndo,
        handleDiagramRedo,
        handleDeleteSelectedComponent
    };
}
