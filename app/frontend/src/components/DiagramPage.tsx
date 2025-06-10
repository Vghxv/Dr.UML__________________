import React, { useState, useEffect } from 'react';
import { SelectDiagram, OpenDiagram, GetDrawData, SaveDiagramFileDialog, GetAvailableDiagramsNames } from '../../wailsjs/go/umlproject/UMLProject';

interface ProjectData {
    ProjectName: string;
    diagrams: string[];
}

interface DiagramPageProps {
    projectData: ProjectData;
    onDiagramSelected: (diagramData: any) => void;
}

const DiagramPage: React.FC<DiagramPageProps> = ({ projectData, onDiagramSelected }) => {
    const [selectedDiagram, setSelectedDiagram] = useState<string | null>(null);
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const handleDiagramClick = async (diagramPath: string) => {
        setIsLoading(true);
        setError(null);
        setSelectedDiagram(diagramPath);

        try {
            // Open the diagram first
            await OpenDiagram(diagramPath);
            
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
        <div className="min-h-screen bg-neutral-700 p-4">
            <div className="max-w-4xl mx-auto">
                <div className="flex items-center justify-between mb-6">
                    <div>
                        <h1 className="text-3xl font-bold text-white">Dr.UML</h1>
                        <h2 className="text-xl text-neutral-300 mt-2">
                            Project: {projectData.ProjectName}
                        </h2>
                    </div>
                </div>

                {error && (
                    <div className="bg-red-600 text-white p-3 rounded mb-4">
                        {error}
                    </div>
                )}

                <div className="bg-neutral-800 rounded-lg p-6">
                    {projectData.diagrams.length === 0 ? (
                        <div className="text-center py-12">
                            you do not have any diagrams in this project. nor can you create a new one.

                        </div>
                    ) : (
                        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
                            {projectData.diagrams.map((diagramPath, index) => (
                                <div
                                    key={index}
                                    className={`bg-neutral-700 rounded-lg p-4 cursor-pointer transition-all hover:bg-neutral-600 ${selectedDiagram === diagramPath && isLoading
                                            ? 'opacity-50 cursor-not-allowed'
                                            : ''
                                        }`}
                                    onClick={() => !isLoading && handleDiagramClick(diagramPath)}
                                >
                                    <div className="flex flex-col">
                                        <h4 className="text-lg font-medium text-white mb-2">
                                            {getBaseName(diagramPath)}
                                        </h4>
                                        <p className="text-sm text-neutral-400 break-all">
                                            {diagramPath}
                                        </p>
                                        {selectedDiagram === diagramPath && isLoading && (
                                            <div className="mt-3 text-blue-400 text-sm">
                                                Loading...
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
