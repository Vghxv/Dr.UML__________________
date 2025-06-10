import React from 'react';

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
        canvasBackgroundColor = "#C2C2C2"
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
        canvasBackgroundColor: "#C2C2C2"
    }
) => {
    const handleExport = () => {
        alert("[TODO] Export Project API");
    };
    const handleValidate = () => {
        alert("[TODO] Validate Project API");
    };
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
                        
                        <button
                            onClick={onGetDiagramName}
                            className="flex items-center gap-2 px-4 py-2.5 text-[#F2F2F0] font-bold text-sm tracking-wide uppercase transition-all duration-200 border-2 border-[#4682B4] hover:border-[#B87333] active:border-[#918175]"
                            style={{
                                background: 'linear-gradient(145deg, #4682B4, #333333)',
                                boxShadow: '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                            }}
                        >
                            DIAGRAM NAME
                        </button>
                        
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
                        
                        <div className="w-0.5 h-8 bg-[#4682B4] mx-2 shadow-sm" />
                        
                        <button
                            onClick={handleExport}
                            className="flex items-center gap-2 px-4 py-2.5 text-[#1C1C1C] font-bold text-sm tracking-wide uppercase transition-all duration-200 border-2 border-[#B87333] hover:border-[#B7410E] active:border-[#556B2F]"
                            style={{
                                background: 'linear-gradient(145deg, #B87333, #918175)',
                                boxShadow: '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                            }}
                        >
                            EXPORT
                        </button>
                        
                        <button
                            onClick={handleValidate}
                            className="flex items-center gap-2 px-4 py-2.5 text-[#F2F2F0] font-bold text-sm tracking-wide uppercase transition-all duration-200 border-2 border-[#B7410E] hover:border-[#B87333] active:border-[#4682B4]"
                            style={{
                                background: 'linear-gradient(145deg, #B7410E, #1C1C1C)',
                                boxShadow: '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                            }}
                        >
                            VALIDATE
                        </button>
                    </div>
                </nav>
            </div>
        </header>
    );
};

export default TopMenu;
