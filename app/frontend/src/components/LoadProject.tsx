import React, { useState } from 'react';

interface ProjectData {
    ProjectName: string;
    diagrams: string[];
}

interface LoadProjectProps {
    onProjectLoaded: (projectData: ProjectData) => void;
}

const LoadProject: React.FC<LoadProjectProps> = ({ onProjectLoaded }) => {
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const handleLoadProject = async () => {
        setIsLoading(true);
        setError(null);

        try {
            // Create file input element
            const input = document.createElement('input');
            input.type = 'file';
            input.accept = '.puml';
            input.style.display = 'none';

            // Handle file selection
            input.onchange = async (event) => {
                const file = (event.target as HTMLInputElement).files?.[0];
                if (!file) {
                    setIsLoading(false);
                    return;
                }                try {
                    // Read file content
                    const fileContent = await file.text();
                    
                    // Parse JSON5 format (treating as JSON for now)
                    // Remove any trailing commas and clean up the format
                    const cleanedContent = fileContent
                        .replace(/,(\s*[}\]])/g, '$1') // Remove trailing commas
                        .replace(/([{,]\s*)(\w+):/g, '$1"$2":'); // Quote unquoted keys
                    
                    const projectData: ProjectData = JSON.parse(cleanedContent);
                    
                    // Validate the project data structure
                    if (!projectData.ProjectName || !Array.isArray(projectData.diagrams)) {
                        throw new Error('Invalid project file format');
                    }
                    
                    // TODO: Call backend API to load the project
                    // For now, we'll simulate a successful load
                    console.log('Loaded project data:', projectData);
                    
                    // Call the callback with the loaded data
                    onProjectLoaded(projectData);
                } catch (parseError) {
                    console.error('Error parsing project file:', parseError);
                    setError('Failed to parse project file. Please ensure it is a valid .puml file.');
                } finally {
                    setIsLoading(false);
                }
            };

            // Trigger file dialog
            input.click();
        } catch (err) {
            console.error('Error loading project:', err);
            setError('Failed to load project file.');
            setIsLoading(false);
        }
    };

    return (
        <div className="flex flex-col items-center justify-center min-h-screen bg-neutral-700">
            <div className="bg-neutral-800 p-8 rounded-lg shadow-lg max-w-md w-full">
                <h1 className="text-3xl font-bold text-white text-center mb-6">Dr.UML</h1>
                <h2 className="text-xl text-white text-center mb-8">Load Project</h2>
                
                {error && (
                    <div className="bg-red-600 text-white p-3 rounded mb-4">
                        {error}
                    </div>
                )}
                
                <button
                    onClick={handleLoadProject}
                    disabled={isLoading}
                    className={`w-full py-3 px-4 rounded font-medium transition-colors ${
                        isLoading
                            ? 'bg-neutral-600 text-neutral-400 cursor-not-allowed'
                            : 'bg-blue-600 hover:bg-blue-700 text-white'
                    }`}
                >
                    {isLoading ? 'Loading...' : 'Select Project File (.puml)'}
                </button>
                
                <div className="mt-6 text-sm text-neutral-400">
                    <p className="text-center">Select a .puml project file to load your UML diagrams.</p>
                    <p className="text-center mt-2">Project files should contain project metadata in JSON5 format.</p>
                </div>
            </div>
        </div>
    );
};

export default LoadProject;
