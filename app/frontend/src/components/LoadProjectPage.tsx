import React, { useState } from 'react';
import { GetAvailableDiagramsNames, GetName, OpenFileDialog, SaveFileDialog, LoadProject } from '../../wailsjs/go/umlproject/UMLProject';

interface ProjectData {
    ProjectName: string;
    diagrams: string[];
    isNewEmptyProject?: boolean;
}

interface LoadProjectProps {
    onProjectLoaded: (projectData: ProjectData) => void;
}

const LoadProjectPage: React.FC<LoadProjectProps> = ({ onProjectLoaded }) => {
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const loadProjectFromPath = async (filePath: string, isNewProject: boolean = false) => {
        setIsLoading(true);
        setError(null);

        try {
            console.log('Loading project from path:', filePath);
            
            // Call backend API to load the project
            const loadResult = await LoadProject(filePath);
            if (loadResult && loadResult.Error) {
                throw new Error(loadResult.Error);
            }

            console.log('Project loaded successfully:', loadResult);

            // Get the project name and available diagrams from backend
            const projectName = await GetName();
            const availableDiagrams = await GetAvailableDiagramsNames();
            
            const projectData: ProjectData = {
                ProjectName: projectName,
                diagrams: availableDiagrams || [], // Ensure diagrams is always an array
                isNewEmptyProject: isNewProject
            };
            
            console.log('Successfully loaded project from backend:', projectData);
            onProjectLoaded(projectData);
            
        } catch (err) {
            console.error('Error loading project:', err);
            setError(`Failed to load project: ${err instanceof Error ? err.message : 'Unknown error'}`);
        } finally {
            setIsLoading(false);
        }
    };

    const handleLoadProject = async () => {
        setIsLoading(true);
        setError(null);

        try {
            // Use native file dialog to get absolute path
            const selectedFilePath = await OpenFileDialog();
            
            if (!selectedFilePath) {
                // User cancelled the dialog
                setIsLoading(false);
                return;
            }

            console.log('Selected file path:', selectedFilePath);
            
            // Load the project using the absolute path
            await loadProjectFromPath(selectedFilePath);
            
        } catch (err) {
            console.error('Error loading project:', err);
            setError(`Failed to load project: ${err instanceof Error ? err.message : 'Unknown error'}`);
            setIsLoading(false);
        }
    };

    const handleCreateEmptyProject = async () => {
        setIsLoading(true);
        setError(null);

        try {
            // Use native save file dialog to get the path for the new project
            const selectedFilePath = await SaveFileDialog();
            
            if (!selectedFilePath) {
                // User cancelled the dialog
                setIsLoading(false);
                return;
            }

            console.log('Creating new project at path:', selectedFilePath);
            
            // Create and save the empty project
            await loadProjectFromPath(selectedFilePath, true);
 
            console.log('Empty project created successfully');

            // Get the project name and available diagrams from backend (should be empty)
            const projectName = await GetName();
            const availableDiagrams = await GetAvailableDiagramsNames();
            
            const projectData: ProjectData = {
                ProjectName: projectName,
                diagrams: availableDiagrams || [] // Ensure diagrams is always an array
            };
            
            console.log('Successfully created empty project:', projectData);
            onProjectLoaded(projectData);
            
        } catch (err) {
            console.error('Error creating empty project:', err);
            setError(`Failed to create empty project: ${err instanceof Error ? err.message : 'Unknown error'}`);
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
                    className={`w-full py-3 px-4 rounded font-medium transition-colors mb-4 ${
                        isLoading
                            ? 'bg-neutral-600 text-neutral-400 cursor-not-allowed'
                            : 'bg-blue-600 hover:bg-blue-700 text-white'
                    }`}
                >
                    {isLoading ? 'Loading...' : 'Select Project File (.puml)'}
                </button>

                <div className="text-center text-neutral-400 text-sm mb-4">
                    or
                </div>

                <button
                    onClick={handleCreateEmptyProject}
                    disabled={isLoading}
                    className={`w-full py-3 px-4 rounded font-medium transition-colors ${
                        isLoading
                            ? 'bg-neutral-600 text-neutral-400 cursor-not-allowed'
                            : 'bg-green-600 hover:bg-green-700 text-white'
                    }`}
                >
                    {isLoading ? 'Creating...' : 'Create Empty Project'}
                </button>
                
                <div className="mt-6 text-sm text-neutral-400">
                    <p className="text-center">Select a .puml project file to load your UML diagrams,</p>
                    <p className="text-center">or create a new empty project to get started.</p>
                </div>
            </div>
        </div>
    );
};

export default LoadProjectPage;
