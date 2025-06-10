import React, { useState, useEffect } from 'react';
import { SelectDiagram, OpenDiagram, GetDrawData, SaveDiagramFileDialog, GetAvailableDiagramsNames } from '../../wailsjs/go/umlproject/UMLProject';

interface ProjectData {
    ProjectName: string;
    diagrams: string[];
}

interface DiagramPageProps {
    projectData: ProjectData;
    onDiagramSelected: (diagramData: any) => void;
    isNewEmptyProject?: boolean;
}

const DiagramPage: React.FC<DiagramPageProps> = ({ projectData, onDiagramSelected, isNewEmptyProject = false }) => {
    const [selectedDiagram, setSelectedDiagram] = useState<string | null>(null);
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const handleDiagramClick = async (diagramPath: string) => {
        setIsLoading(true);
        setError(null);
        setSelectedDiagram(diagramPath);

        try {
            // For new empty projects, skip trying to open the diagram file
            if (!isNewEmptyProject) {
                // Check if file exists and open the diagram only if it exists
                try {
                    await OpenDiagram(diagramPath);
                } catch (openError) {
                    // This is normal behavior when file doesn't exist
                }
            }
            
            // Then get the draw data
            const diagramData = await GetDrawData();
            console.log('Diagram data received:', diagramData);
            // Transform the backend data into the proper CanvasProps format
            const canvasData = {
                margin: diagramData.margin,
                color: diagramData.color,
                lineWidth: diagramData.lineWidth,
                gadgets: diagramData.gadgets?.map((gadget: any) => ({
                    gadgetType: gadget.gadgetType.toString(),
                    x: gadget.x,
                    y: gadget.y,
                    layer: gadget.layer,
                    height: gadget.height,
                    width: gadget.width,
                    color: gadget.color,
                    isSelected: gadget.isSelected,
                    attributes: gadget.attributes
                })) || [],
                associations: diagramData.associations?.map((association: any) => ({
                    assType: association.assType,
                    layer: association.layer,
                    startX: association.startX,
                    startY: association.startY,
                    endX: association.endX,
                    endY: association.endY,
                    deltaX: association.deltaX,
                    deltaY: association.deltaY,
                    attributes: association.attributes?.map((attr: any) => ({
                        content: attr.content,
                        fontSize: attr.fontSize,
                        fontStyle: attr.fontStyle,
                        fontFile: attr.fontFile,
                        ratio: attr.ratio
                    })) || []
                })) || []
            };

            // Call the callback to switch to editor view with the properly formatted data
            onDiagramSelected(canvasData);
        } catch (err) {
            console.error('Error loading diagram:', err);
            setError(`Failed to load diagram: ${err instanceof Error ? err.message : 'Unknown error'}`);
        } finally {
            setIsLoading(false);
            setSelectedDiagram(null);
        }
    };

    const getBaseName = (filePath: string): string => {
        const parts = filePath.replace(/\\/g, '/').split('/');
        const fileName = parts[parts.length - 1];
        return fileName.replace(/\.duml$/, '');
    };

    return (
        <div 
            className="min-h-screen p-6 relative"
            style={{
                background: 'linear-gradient(135deg, #333333 0%, #1C1C1C 50%, #333333 100%)',
            }}
        >
            {/* Industrial texture overlay */}
            <div 
                className="absolute inset-0 opacity-10"
                style={{
                    backgroundImage: `url("data:image/svg+xml,%3Csvg width='40' height='40' viewBox='0 0 40 40' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='%23F2F2F0' fill-opacity='0.1'%3E%3Cpath d='M20 20c0 11.046-8.954 20-20 20v-40c11.046 0 20 8.954 20 20zM0 0h40v40H0V0z'/%3E%3C/g%3E%3C/svg%3E")`,
                }}
            />
            
            <div className="max-w-6xl mx-auto relative z-10">
                {/* Header Section */}
                <div 
                    className="flex items-center justify-between mb-8 p-6 border-2 border-[#4682B4]"
                    style={{
                        background: 'linear-gradient(145deg, #4682B4, #333333)',
                        boxShadow: '6px 6px 12px rgba(0,0,0,0.6), inset 2px 2px 4px rgba(242,242,240,0.1)'
                    }}
                >
                    <div>
                        <h1 className="text-4xl font-bold text-[#F2F2F0] tracking-wider uppercase mb-2">DR.UML</h1>
                        <div 
                            className="px-4 py-2 border-2 border-[#B87333] inline-block"
                            style={{
                                background: 'linear-gradient(145deg, #B87333, #1C1C1C)',
                                boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.3)'
                            }}
                        >
                            <h2 className="text-lg text-[#F2F2F0] font-bold tracking-wide uppercase">
                                PROJECT: {projectData.ProjectName}
                            </h2>
                        </div>
                    </div>
                </div>

                {error && (
                    <div 
                        className="p-4 mb-6 border-2 border-[#B7410E]"
                        style={{
                            background: 'linear-gradient(145deg, #B7410E, #1C1C1C)',
                            boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.3)'
                        }}
                    >
                        <p className="text-[#F2F2F0] font-bold text-sm tracking-wide">{error}</p>
                    </div>
                )}

                <div 
                    className="p-8 border-2 border-[#918175]"
                    style={{
                        background: 'linear-gradient(145deg, #918175, #333333)',
                        boxShadow: '6px 6px 12px rgba(0,0,0,0.6), inset 2px 2px 4px rgba(242,242,240,0.1)'
                    }}
                >
                    {projectData.diagrams.length === 0 ? (
                        <div 
                            className="text-center py-16 border-2 border-[#556B2F]"
                            style={{
                                background: 'linear-gradient(145deg, #333333, #1C1C1C)',
                                boxShadow: 'inset 4px 4px 8px rgba(0,0,0,0.5)'
                            }}
                        >
                            <p className="text-[#F2F2F0] font-bold text-lg tracking-wide uppercase">
                                YOU DO NOT HAVE ANY DIAGRAMS IN THIS PROJECT.
                            </p>
                            <p className="text-[#918175] font-bold text-sm tracking-wide uppercase mt-2">
                                NOR CAN YOU CREATE A NEW ONE.
                            </p>
                        </div>
                    ) : (
                        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
                            {projectData.diagrams.map((diagramPath, index) => (
                                <div
                                    key={index}
                                    className={`p-6 cursor-pointer transition-all duration-200 border-2 border-[#4682B4] hover:border-[#B87333] ${
                                        selectedDiagram === diagramPath && isLoading
                                            ? 'opacity-50 cursor-not-allowed'
                                            : 'hover:scale-105'
                                    }`}
                                    style={{
                                        background: selectedDiagram === diagramPath && isLoading
                                            ? 'linear-gradient(145deg, #918175, #1C1C1C)'
                                            : 'linear-gradient(145deg, #556B2F, #333333)',
                                        boxShadow: selectedDiagram === diagramPath && isLoading
                                            ? 'inset 4px 4px 8px rgba(0,0,0,0.5)'
                                            : '4px 4px 8px rgba(0,0,0,0.5), inset 1px 1px 2px rgba(242,242,240,0.1)'
                                    }}
                                    onClick={() => !isLoading && handleDiagramClick(diagramPath)}
                                >
                                    <div className="flex flex-col">
                                        {/* Diagram Title */}
                                        <div 
                                            className="mb-4 p-3 border-2 border-[#B87333]"
                                            style={{
                                                background: 'linear-gradient(145deg, #B87333, #1C1C1C)',
                                                boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.3)'
                                            }}
                                        >
                                            <h4 className="text-lg font-bold text-[#F2F2F0] tracking-wide uppercase">
                                                {getBaseName(diagramPath)}
                                            </h4>
                                        </div>
                                        
                                        {/* Diagram Path */}
                                        <div 
                                            className="p-3 border-2 border-[#918175] mb-3"
                                            style={{
                                                background: 'linear-gradient(145deg, #333333, #1C1C1C)',
                                                boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.5)'
                                            }}
                                        >
                                            <p className="text-xs font-bold text-[#918175] tracking-wide uppercase break-all">
                                                {diagramPath}
                                            </p>
                                        </div>
                                        
                                        {/* Loading State */}
                                        {selectedDiagram === diagramPath && isLoading && (
                                            <div 
                                                className="p-2 border-2 border-[#4682B4]"
                                                style={{
                                                    background: 'linear-gradient(145deg, #4682B4, #1C1C1C)',
                                                    boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.3)'
                                                }}
                                            >
                                                <p className="text-[#F2F2F0] font-bold text-sm tracking-wide uppercase text-center">
                                                    LOADING...
                                                </p>
                                            </div>
                                        )}
                                    </div>
                                </div>
                            ))}
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
};

export default DiagramPage;
