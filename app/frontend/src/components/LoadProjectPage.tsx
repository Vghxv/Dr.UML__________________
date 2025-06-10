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
        <div 
            className="flex flex-col items-center justify-center min-h-screen relative"
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
            
            <div 
                className="relative z-10 p-8 max-w-md w-full border-2 border-[#4682B4]"
                style={{
                    background: 'linear-gradient(145deg, #918175, #333333)',
                    boxShadow: '6px 6px 12px rgba(0,0,0,0.6), inset 2px 2px 4px rgba(242,242,240,0.1)'
                }}
            >
                {/* Industrial Header */}
                <div 
                    className="text-center mb-8 p-4 border-2 border-[#B87333]"
                    style={{
                        background: 'linear-gradient(145deg, #4682B4, #1C1C1C)',
                        boxShadow: '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)'
                    }}
                >
                    <h1 className="text-3xl font-bold text-[#F2F2F0] tracking-wider uppercase mb-2">DR.UML</h1>
                    <h2 className="text-lg text-[#F2F2F0] font-bold tracking-wide uppercase">LOAD PROJECT</h2>
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
                
                <button
                    onClick={handleLoadProject}
                    disabled={isLoading}
                    className={`w-full py-4 px-6 font-bold text-sm tracking-wide uppercase transition-all duration-200 mb-6 border-2 ${
                        isLoading
                            ? 'border-[#918175] cursor-not-allowed opacity-50'
                            : 'border-[#4682B4] hover:border-[#B87333] active:border-[#556B2F]'
                    }`}
                    style={{
                        background: isLoading 
                            ? 'linear-gradient(145deg, #918175, #1C1C1C)' 
                            : 'linear-gradient(145deg, #4682B4, #333333)',
                        boxShadow: isLoading 
                            ? 'inset 2px 2px 4px rgba(0,0,0,0.5)' 
                            : '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)',
                        color: '#F2F2F0'
                    }}
                >
                    {isLoading ? 'LOADING...' : 'SELECT PROJECT FILE (.PUML)'}
                </button>

                <div 
                    className="text-center mb-6 py-2"
                    style={{
                        background: 'linear-gradient(90deg, transparent, #4682B4, transparent)',
                        height: '2px'
                    }}
                >
                    <span 
                        className="px-4 text-sm font-bold text-[#F2F2F0] tracking-wider uppercase"
                        style={{ 
                            background: 'linear-gradient(145deg, #918175, #333333)',
                            lineHeight: '2px',
                            position: 'relative',
                            top: '-8px'
                        }}
                    >
                        OR
                    </span>
                </div>

                <button
                    onClick={handleCreateEmptyProject}
                    disabled={isLoading}
                    className={`w-full py-4 px-6 font-bold text-sm tracking-wide uppercase transition-all duration-200 mb-6 border-2 ${
                        isLoading
                            ? 'border-[#918175] cursor-not-allowed opacity-50'
                            : 'border-[#556B2F] hover:border-[#B87333] active:border-[#4682B4]'
                    }`}
                    style={{
                        background: isLoading 
                            ? 'linear-gradient(145deg, #918175, #1C1C1C)' 
                            : 'linear-gradient(145deg, #556B2F, #333333)',
                        boxShadow: isLoading 
                            ? 'inset 2px 2px 4px rgba(0,0,0,0.5)' 
                            : '3px 3px 6px rgba(0,0,0,0.4), inset 1px 1px 2px rgba(242,242,240,0.1)',
                        color: '#F2F2F0'
                    }}
                >
                    {isLoading ? 'CREATING...' : 'CREATE EMPTY PROJECT'}
                </button>
                
                <div 
                    className="p-4 border-2 border-[#918175]"
                    style={{
                        background: 'linear-gradient(145deg, #333333, #1C1C1C)',
                        boxShadow: 'inset 2px 2px 4px rgba(0,0,0,0.5)'
                    }}
                >
                    <p className="text-center text-xs font-bold text-[#F2F2F0] tracking-wide uppercase mb-2">
                        SELECT A .PUML PROJECT FILE TO LOAD YOUR UML DIAGRAMS,
                    </p>
                    <p className="text-center text-xs font-bold text-[#F2F2F0] tracking-wide uppercase">
                        OR CREATE A NEW EMPTY PROJECT TO GET STARTED.
                    </p>
                </div>
            </div>
        </div>
    );
};

export default LoadProjectPage;
