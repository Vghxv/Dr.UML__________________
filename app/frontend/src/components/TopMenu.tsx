import React from 'react';



const TopMenu: React.FC<{}> = () => {
    const handleOpenProject = () => {
        alert("[TODO] Open Project API");
    };
    const handleSave = () => {
        alert("[TODO] Save Project API");
    };
    const handleExport = () => {
        alert("[TODO] Export Project API");
    };
    const handleValidate = () => {
        alert("[TODO] Validate Project API");
    };
    return (
        <header className="w-full shadow-lg bg-gradient-to-r from-blue-900 via-blue-700 to-blue-500 text-white sticky top-0 z-50">
            <div className="flex items-center justify-between px-8 py-3">
                <div className="flex items-center gap-3">
                    {/* <img src="/favicon.ico" alt="Dr.UML Logo" className="w-8 h-8 rounded shadow-md" /> */}
                    <span className="text-2xl font-extrabold tracking-wider drop-shadow-lg select-none">Dr.UML</span>
                </div>
                <nav className="flex items-center gap-2.5">
                    <button
                        onClick={handleOpenProject}
                        className="flex items-center gap-1 bg-green-500 hover:bg-green-600 active:bg-green-700 text-white px-4 py-2 rounded-lg shadow transition-all duration-150 font-semibold border border-green-400 hover:scale-105"
                    >
                        Open
                    </button>
                    <span className="w-px h-6 bg-white/30 mx-1" />
                    <button
                        onClick={handleSave}
                        className="flex items-center gap-1 bg-blue-500 hover:bg-blue-600 active:bg-blue-700 text-white px-4 py-2 rounded-lg shadow transition-all duration-150 font-semibold border border-blue-400 hover:scale-105"
                    >
                        Save
                    </button>
                    <span className="w-px h-6 bg-white/30 mx-1" />
                    <button
                        onClick={handleExport}
                        className="flex items-center gap-1 bg-orange-500 hover:bg-orange-600 active:bg-orange-700 text-white px-4 py-2 rounded-lg shadow transition-all duration-150 font-semibold border border-orange-400 hover:scale-105"
                    >
                        Export
                    </button>
                    <span className="w-px h-6 bg-white/30 mx-1" />
                    <button
                        onClick={handleValidate}
                        className="flex items-center gap-1 bg-red-500 hover:bg-red-600 active:bg-red-700 text-white px-4 py-2 rounded-lg shadow transition-all duration-150 font-semibold border border-red-400 hover:scale-105"
                    >
                        Validate
                    </button>
                </nav>
            </div>
        </header>
    );
};

export default TopMenu;
