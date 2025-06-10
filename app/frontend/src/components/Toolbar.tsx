import React from "react";

interface ToolbarProps {
    onGetDiagramName: () => void;
    onShowPopup: () => void;
    onAddAss: () => void;
    onSaveProject: () => void;
    onSaveDiagram: () => void;
    onCanvasColorChange: (color: string) => void;
    diagramName?: string | null;
    canvasBackgroundColor: string;

}

const Toolbar: React.FC<ToolbarProps> = ({
    onGetDiagramName,
    onShowPopup,
    onAddAss,
    onSaveProject,
    onSaveDiagram,
    onCanvasColorChange,
    diagramName,
    canvasBackgroundColor = "#C2C2C2"
}) => (
    <div className="flex items-center gap-4 py-4 px-8 bg-gradient-to-r from-blue-100 via-blue-50 to-white rounded-2xl shadow-lg min-h-[60px] border border-blue-200 mb-4">
        <button
            onClick={onGetDiagramName}
            className="flex items-center gap-1 bg-blue-500 hover:bg-blue-600 active:bg-blue-700 text-white font-semibold text-sm py-2.5 px-5 rounded-lg shadow border border-blue-400 transition-all duration-150 hover:scale-105"
        >
            Get Diagram Name
        </button>
        {diagramName && (
            <span className="text-blue-900 font-medium text-base ml-1 mr-2 tracking-wide bg-blue-100 rounded px-2 py-1 border border-blue-200">
                {diagramName}
            </span>
        )}
        <span className="w-px h-6 bg-blue-200 mx-1" />
        <button
            onClick={onShowPopup}
            className="flex items-center gap-1 bg-yellow-400 hover:bg-yellow-500 active:bg-yellow-600 text-gray-900 font-semibold text-sm py-2.5 px-5 rounded-lg shadow border border-yellow-300 transition-all duration-150 hover:scale-105"
        >
            + Create Gadget
        </button>
        <span className="w-px h-6 bg-blue-200 mx-1" />
        <button
            onClick={onAddAss}
            className="flex items-center gap-1 bg-red-500 hover:bg-red-600 active:bg-red-700 text-white font-semibold text-sm py-2.5 px-5 rounded-lg shadow border border-red-400 transition-all duration-150 hover:scale-105"
        >
            Add Association
        </button>
        <div className="flex items-center gap-2 ml-4">
            <label htmlFor="canvas-color" className="text-black font-medium text-sm">
                Canvas Color:
            </label>
            <input
                id="canvas-color"
                type="color"
                value={canvasBackgroundColor}
                onChange={(e) => onCanvasColorChange(e.target.value)}
                className="w-10 h-8 rounded border border-gray-300 cursor-pointer"
                title="Change canvas background color"
            />
        </div>
        <button
            onClick={onSaveProject}
            className="bg-green-500 hover:bg-green-600 text-white font-semibold text-sm py-2.5 px-5 rounded-md shadow-sm transition-colors cursor-pointer"
        >
            Save Project
        </button>
        <button
            onClick={onSaveDiagram}
            className="bg-emerald-500 hover:bg-emerald-600 text-white font-semibold text-sm py-2.5 px-5 rounded-md shadow-sm transition-colors cursor-pointer"
        >
            Save Diagram
        </button>
    </div>
);

export default Toolbar;
