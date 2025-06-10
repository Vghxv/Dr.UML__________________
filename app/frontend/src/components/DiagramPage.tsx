import React, { useState, useEffect } from 'react';
import { SelectDiagram, OpenDiagram, GetDrawData, SaveDiagramFileDialog, GetAvailableDiagramsNames } from '../../wailsjs/go/umlproject/UMLProject';

interface ProjectData {
    ProjectName: string;
    diagrams: string[];
}

interface DiagramPageProps {
    projectData: ProjectData;
    onBack: () => void;
    onDiagramSelected: (diagramData: any) => void;
}

const DiagramPage: React.FC<DiagramPageProps> = ({ projectData, onBack, onDiagramSelected }) => {
    const [selectedDiagram, setSelectedDiagram] = useState<string | null>(null);
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const handleDiagramClick = async (diagramPath: string) => {
        setIsLoading(true);
        setError(null);
        setSelectedDiagram(diagramPath);

        try {
            const diagramData = await GetDrawData();

            const frontendDiagramData = {
                diagramName: getBaseName(diagramPath),
                diagramType: "ClassDiagram", // This could be extracted from backend data
                components: diagramData.gadgets || [],
                associations: diagramData.associations || []
            };
            // Call the callback to switch to editor view
            onDiagramSelected(frontendDiagramData);
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
