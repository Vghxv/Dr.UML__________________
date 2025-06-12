import React, { useEffect } from 'react';

interface ToolbarProps {
    projectData: any
    handleBackToDiagrams: () => void;
    onGetDiagramName: () => void;
    onShowPopup: () => void;
    onAddAss: () => void;
    onSaveProject: () => void;
    onSaveDiagram: () => void;
    onCanvasColorChange: (color: string) => void;
    diagramName?: string | null;
    canvasBackgroundColor: string;
    onUndo: () => void;
    onRedo: () => void;
    onDeleteSelectedComponent: () => void;
}

const TopMenu: React.FC<ToolbarProps> = (
    { 
        projectData, 
        handleBackToDiagrams,
        onGetDiagramName,
        onShowPopup,
        onAddAss,
        onSaveProject,
        onSaveDiagram,
        onCanvasColorChange,
        diagramName,
        canvasBackgroundColor = "#C2C2C2",
        onUndo,
        onRedo,
        onDeleteSelectedComponent
    } = { 
        projectData: null, 
        handleBackToDiagrams: () => {},
        onGetDiagramName: () => {},
        onShowPopup: () => {},
        onAddAss: () => {},
        onSaveProject: () => {},
        onSaveDiagram: () => {},
        onCanvasColorChange: () => {},
        diagramName: null,
        canvasBackgroundColor: "#C2C2C2",
        onUndo: () => {},
        onRedo: () => {},
        onDeleteSelectedComponent: () => {}
    }
) => {
    useEffect(() => {
        if (onGetDiagramName) {
            onGetDiagramName();
        }
    }, [onGetDiagramName]);

    return (
        <header className="w-full sticky top-0 z-50 shadow-2xl border-b-2 border-[#1C1C1C]">
            <div 
                className="relative px-6 py-4"
                style={{
                    background: 'linear-gradient(135deg, #333333 0%, #1C1C1C 50%, #333333 100%)',
                    boxShadow: 'inset 0 1px 0 rgba(242, 242, 240, 0.1), inset 0 -1px 0 rgba(0, 0, 0, 0.3)'
                }}
            >
                {/* Industrial texture overlay */}
                <div 
                    className="absolute inset-0 opacity-10"
                    style={{
                        backgroundImage: `url("data:image/svg+xml,%3Csvg width='40' height='40' viewBox='0 0 40 40' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='%23F2F2F0' fill-opacity='0.1'%3E%3Cpath d='M20 20c0 11.046-8.954 20-20 20v-40c11.046 0 20 8.954 20 20zM0 0h40v40H0V0z'/%3E%3C/g%3E%3C/svg%3E")`,
                    }}
                />
                
                <nav className="relative flex items-center justify-between flex-wrap gap-3">
                    <div className="flex items-center gap-3 flex-wrap">
                        {projectData && (
                            <>
                                <button
                                    onClick={handleBackToDiagrams}
                                    className="flex items-center gap-2 px-4 py-2.5 text-[#F2F2F0] font-bold text-sm tracking-wide uppercase transition-all duration-200 border-2 border-[#918175] hover:border-[#B87333] active:border-[#4682B4] transform hover:translate-y-[-1px] active:translate-y-0"
                                    style={{
                                        background: 'linear-gradient(145deg, #918175, #333333)',
                                        boxShadow: '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                                    }}
                                >
                                    ← DIAGRAMS
                                </button>
                                <div className="w-0.5 h-8 bg-[#4682B4] mx-2 shadow-sm" />
                            </>
                        )}
                        
                        {diagramName && (
                            <div 
                                className="px-3 py-2 text-[#1C1C1C] font-bold text-sm tracking-wider uppercase border-2 border-[#B87333]"
                                style={{
                                    background: 'linear-gradient(145deg, #F2F2F0, #B87333)',
                                    boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.3)'
                                }}
                            >
                                {diagramName}
                            </div>
                        )}
                        
                        <div className="w-0.5 h-8 bg-[#4682B4] mx-2 shadow-sm" />
                        
                        <button
                            onClick={onShowPopup}
                            className="flex items-center gap-2 px-4 py-2.5 text-[#1C1C1C] font-bold text-sm tracking-wide uppercase transition-all duration-200 border-2 border-[#B87333] hover:border-[#B7410E] active:border-[#556B2F]"
                            style={{
                                background: 'linear-gradient(145deg, #B87333, #918175)',
                                boxShadow: '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                            }}
                        >
                            + GADGET
                        </button>
                        
                        <div className="w-0.5 h-8 bg-[#4682B4] mx-2 shadow-sm" />
                        
                        <button
                            onClick={onAddAss}
                            className="flex items-center gap-2 px-4 py-2.5 text-[#F2F2F0] font-bold text-sm tracking-wide uppercase transition-all duration-200 border-2 border-[#B7410E] hover:border-[#B87333] active:border-[#4682B4]"
                            style={{
                                background: 'linear-gradient(145deg, #B7410E, #333333)',
                                boxShadow: '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                            }}
                        >
                            + ASSOCIATION
                        </button>
                        
                        <button
                            onClick={onDeleteSelectedComponent}
                            className="flex items-center gap-2 px-4 py-2.5 text-[#fff] font-bold text-sm tracking-wide uppercase transition-all duration-200 border-2 border-red-600 hover:border-red-400 active:border-red-800 ml-2"
                            style={{
                                background: 'linear-gradient(145deg, #e3342f, #b91c1c)',
                                boxShadow: '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                            }}
                        >
                            Delete component
                        </button>

                        <div className="flex items-center gap-3 ml-4">
                            <label htmlFor="canvas-color" className="text-[#F2F2F0] font-bold text-sm tracking-wide uppercase">
                                CANVAS:
                            </label>
                            <div 
                                className="p-1 border-2 border-[#918175]"
                                style={{
                                    background: 'linear-gradient(145deg, #333333, #1C1C1C)',
                                    boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.5)'
                                }}
                            >
                                <input
                                    id="canvas-color"
                                    type="color"
                                    value={canvasBackgroundColor}
                                    onChange={(e) => onCanvasColorChange(e.target.value)}
                                    className="w-8 h-6 border-0 cursor-pointer"
                                    title="Change canvas background color"
                                />
                            </div>
                        </div>
                    </div>
                    
                    <div className="flex items-center gap-3 flex-wrap">
                        <button
                            onClick={onSaveProject}
                            className="flex items-center gap-2 px-4 py-2.5 text-[#F2F2F0] font-bold text-sm tracking-wide uppercase transition-all duration-200 border-2 border-[#556B2F] hover:border-[#4682B4] active:border-[#B87333]"
                            style={{
                                background: 'linear-gradient(145deg, #556B2F, #333333)',
                                boxShadow: '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                            }}
                        >
                            SAVE PROJECT
                        </button>
                        
                        <button
                            onClick={onSaveDiagram}
                            className="flex items-center gap-2 px-4 py-2.5 text-[#F2F2F0] font-bold text-sm tracking-wide uppercase transition-all duration-200 border-2 border-[#556B2F] hover:border-[#4682B4] active:border-[#B87333]"
                            style={{
                                background: 'linear-gradient(145deg, #556B2F, #1C1C1C)',
                                boxShadow: '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                            }}
                        >
                            SAVE DIAGRAM
                        </button>

                        <button
                            onClick={onUndo}
                            className="flex items-center gap-2 px-4 py-2.5 text-[#F2F2F0] font-bold text-sm tracking-wide uppercase transition-all duration-200 border-2 border-[#4682B4] hover:border-[#B87333] active:border-[#556B2F] rounded shadow-md hover:shadow-lg active:shadow-sm"
                            style={{
                                background: 'linear-gradient(145deg, #4682B4, #1C1C1C)',
                                boxShadow: '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                            }}
                            title="Undo"
                        >
                            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 14H5.41a1 1 0 01-.71-1.71l5.3-5.29a1 1 0 011.42 0l5.3 5.29A1 1 0 0118.59 14H15" />
                            </svg>
                            UNDO
                        </button>
                        <button
                            onClick={onRedo}
                            className="flex items-center gap-2 px-4 py-2.5 text-[#F2F2F0] font-bold text-sm tracking-wide uppercase transition-all duration-200 border-2 border-[#B87333] hover:border-[#4682B4] active:border-[#556B2F] rounded shadow-md hover:shadow-lg active:shadow-sm"
                            style={{
                                background: 'linear-gradient(145deg, #B87333, #1C1C1C)',
                                boxShadow: '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                            }}
                            title="Redo"
                        >
                            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 10h3.59a1 1 0 01.71 1.71l-5.3 5.29a1 1 0 01-1.42 0l-5.3-5.29A1 1 0 015.41 10H9" />
                            </svg>
                            REDO
                        </button>
                        
                        
                    </div>
                </nav>
            </div>
        </header>
    );
};

export default TopMenu;
