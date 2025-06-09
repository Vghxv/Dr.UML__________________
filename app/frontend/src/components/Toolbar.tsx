import React from "react";

interface ToolbarProps {
    onGetDiagramName: () => void;
    onShowPopup: () => void;
    onAddAss: () => void;
    diagramName?: string | null;
}

const Toolbar: React.FC<ToolbarProps> = ({
    onGetDiagramName,
    onShowPopup,
    onAddAss,
    diagramName
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
    </div>
);

export default Toolbar;
