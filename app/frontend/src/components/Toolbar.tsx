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
    <div className="flex items-center gap-4 my-6 py-3.5 px-8 bg-gray-900 rounded-xl shadow-md min-h-[60px]">
        <button
            onClick={onGetDiagramName}
            className="bg-blue-500 hover:bg-blue-600 text-white font-semibold text-sm py-2.5 px-5 rounded-md shadow-sm transition-colors cursor-pointer"
        >
            Get Diagram Name
        </button>
        {diagramName && (
            <span className="text-white font-medium text-sm ml-1 mr-2 tracking-wide">
                Diagram Name: {diagramName}
            </span>
        )}
        <button
            onClick={onShowPopup}
            className="bg-yellow-500 hover:bg-yellow-600 text-gray-900 font-semibold text-sm py-2.5 px-5 rounded-md shadow-sm transition-colors cursor-pointer"
        >
            + Create Gadget (Popup)
        </button>
        <button
            onClick={onAddAss}
            className="bg-red-500 hover:bg-red-600 text-white font-semibold text-sm py-2.5 px-5 rounded-md shadow-sm transition-colors cursor-pointer"
        >
            Add Association
        </button>
    </div>
);

export default Toolbar;
