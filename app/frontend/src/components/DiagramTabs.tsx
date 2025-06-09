import React from "react";

interface DiagramTab {
    name: string;
}

interface DiagramTabsProps {
    diagrams: DiagramTab[];
    activeDiagram: string | null;
    onSelect: (name: string) => void;
    onClose: (name: string) => void;
    onAdd: () => void;
}

const DiagramTabs: React.FC<DiagramTabsProps> = ({ diagrams, activeDiagram, onSelect, onClose, onAdd }) => {
    return (
        <div className="flex items-end space-x-2 mb-4 px-2">
            {diagrams.map((tab) => (
                <div
                    key={tab.name}
                    className={`flex items-center px-5 py-2 rounded-t-2xl shadow transition-all duration-150 border-b-4 cursor-pointer select-none
                        ${activeDiagram === tab.name
                            ? 'bg-white text-blue-700 border-blue-700 font-bold scale-105 z-10'
                            : 'bg-blue-100 text-blue-400 border-transparent hover:bg-blue-200 hover:text-blue-700'}
                    `}
                    onClick={() => onSelect(tab.name)}
                    style={{ minWidth: 120 }}
                >
                    <span className="truncate max-w-[80px]">{tab.name}</span>
                    <button
                        className="ml-2 text-gray-400 hover:text-red-500 rounded-full p-1 transition-colors"
                        onClick={e => { e.stopPropagation(); onClose(tab.name); }}
                        title="Close tab"
                    >
                        <span className="material-icons text-base">close</span>
                    </button>
                </div>
            ))}
            <button
                className="ml-2 px-3 py-2 rounded-full bg-blue-600 text-white hover:bg-blue-700 shadow border border-blue-400 transition-all duration-150 flex items-center gap-1"
                onClick={onAdd}
                title="Add new diagram tab"
            >
                <span className="material-icons text-base">add</span>
            </button>
        </div>
    );
};

export default DiagramTabs;
