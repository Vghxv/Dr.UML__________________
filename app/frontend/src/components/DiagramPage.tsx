import React, { useState } from 'react';
import path from 'path-browserify';

interface ProjectData {
    ProjectName: string;
    diagrams: string[];
}

interface DiagramPageProps {
    projectData: ProjectData;
    onBack: () => void;
}

const DiagramPage: React.FC<DiagramPageProps> = ({ projectData, onBack }) => {
    const [selectedDiagram, setSelectedDiagram] = useState<string | null>(null);
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const handleDiagramClick = async (diagramPath: string) => {
        setIsLoading(true);
        setError(null);
        setSelectedDiagram(diagramPath);

        try {
            // TODO: Call backend API to load the .duml file
            // For now, we'll simulate loading the diagram
            console.log('Loading diagram:', diagramPath);
            
            // Simulate API call delay
            await new Promise(resolve => setTimeout(resolve, 500));
            
            // Mock diagram data in JSON5 format
            const mockDiagramData = {
                diagramName: path.basename(diagramPath, '.duml'),
                diagramType: "ClassDiagram",
                components: [
                    {
                        type: "Class",
                        name: "Example Class",
                        x: 100,
                        y: 100,
                        attributes: ["attribute1: String", "attribute2: int"],
                        methods: ["method1(): void", "method2(): String"]
                    }
                ],
                associations: []
            };
            
            console.log('Loaded diagram data:', mockDiagramData);
            
            // TODO: Here you would typically navigate to the main editor or update the app state
            // with the loaded diagram data
            
        } catch (err) {
            console.error('Error loading diagram:', err);
            setError('Failed to load diagram file.');
        } finally {
            setIsLoading(false);
            setSelectedDiagram(null);
        }
    };

    const getBaseName = (filePath: string): string => {
        return path.basename(filePath, '.duml');
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
                    <button
                        onClick={onBack}
                        className="bg-neutral-600 hover:bg-neutral-500 text-white px-4 py-2 rounded transition-colors"
                    >
                        Back to Load Project
                    </button>
                </div>

                {error && (
                    <div className="bg-red-600 text-white p-3 rounded mb-4">
                        {error}
                    </div>
                )}

                <div className="bg-neutral-800 rounded-lg p-6">
                    <h3 className="text-xl font-semibold text-white mb-4">
                        Available Diagrams ({projectData.diagrams.length})
                    </h3>
                    
                    {projectData.diagrams.length === 0 ? (
                        <div className="text-neutral-400 text-center py-8">
                            No diagrams found in this project.
                        </div>
                    ) : (
                        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
                            {projectData.diagrams.map((diagramPath, index) => (
                                <div
                                    key={index}
                                    className={`bg-neutral-700 rounded-lg p-4 cursor-pointer transition-all hover:bg-neutral-600 ${
                                        selectedDiagram === diagramPath && isLoading
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
