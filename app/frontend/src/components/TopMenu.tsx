import React from 'react';
// TODO: integrate with to App

const TopMenu: React.FC = () => {
    const handleOpenProject = () => {
        // Logic to open a project
        console.log('Open Project clicked');
    };

    const handleSave = () => {
        // Logic to save the current project
        console.log('Save clicked');
    };

    const handleExport = () => {
        // Logic to export the project
        console.log('Export clicked');
    };

    const handleValidate = () => {
        // Logic to validate the UML diagram
        console.log('Validate clicked');
    };

    return (
        <div className="flex items-center justify-between px-5 py-2.5 bg-gray-800 text-white shadow-md">
            <h1 className="text-lg font-semibold m-0">Dr.UML</h1>
            <div className="flex gap-2.5">
                <button
                    onClick={handleOpenProject}
                    className="bg-green-600 hover:bg-green-700 text-white px-3 py-2 rounded cursor-pointer transition-colors"
                >
                    Open Project
                </button>
                <button
                    onClick={handleSave}
                    className="bg-blue-600 hover:bg-blue-700 text-white px-3 py-2 rounded cursor-pointer transition-colors"
                >
                    Save
                </button>
                <button
                    onClick={handleExport}
                    className="bg-orange-500 hover:bg-orange-600 text-white px-3 py-2 rounded cursor-pointer transition-colors"
                >
                    Export
                </button>
                <button
                    onClick={handleValidate}
                    className="bg-red-600 hover:bg-red-700 text-white px-3 py-2 rounded cursor-pointer transition-colors"
                >
                    Validate
                </button>
            </div>
        </div>
    );
};

export default TopMenu;
