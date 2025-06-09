import React from 'react';

interface TopMenuProps {
    onOpenProject: () => void;
    onSave: () => void;
    onExport: () => void;
    onValidate: () => void;
}

const TopMenu: React.FC<TopMenuProps & { sessionName: string | null; isConnected: boolean; onJoinSession: () => void; onLeaveSession: () => void; }> = ({ onOpenProject, onSave, onExport, onValidate, sessionName, isConnected, onJoinSession, onLeaveSession }) => {
    return (
        <header className="w-full shadow-lg bg-gradient-to-r from-blue-900 via-blue-700 to-blue-500 text-white sticky top-0 z-50">
            <div className="flex items-center justify-between px-8 py-3">
                <div className="flex items-center gap-3">
                    {/* <img src="/favicon.ico" alt="Dr.UML Logo" className="w-8 h-8 rounded shadow-md" /> */}
                    <span className="text-2xl font-extrabold tracking-wider drop-shadow-lg select-none">Dr.UML</span>
                </div>
                <nav className="flex items-center gap-2.5">
                    <button
                        onClick={onOpenProject}
                        className="flex items-center gap-1 bg-green-500 hover:bg-green-600 active:bg-green-700 text-white px-4 py-2 rounded-lg shadow transition-all duration-150 font-semibold border border-green-400 hover:scale-105"
                    >
                        Open
                    </button>
                    <span className="w-px h-6 bg-white/30 mx-1" />
                    <button
                        onClick={onSave}
                        className="flex items-center gap-1 bg-blue-500 hover:bg-blue-600 active:bg-blue-700 text-white px-4 py-2 rounded-lg shadow transition-all duration-150 font-semibold border border-blue-400 hover:scale-105"
                    >
                        Save
                    </button>
                    <span className="w-px h-6 bg-white/30 mx-1" />
                    <button
                        onClick={onExport}
                        className="flex items-center gap-1 bg-orange-500 hover:bg-orange-600 active:bg-orange-700 text-white px-4 py-2 rounded-lg shadow transition-all duration-150 font-semibold border border-orange-400 hover:scale-105"
                    >
                        Export
                    </button>
                    <span className="w-px h-6 bg-white/30 mx-1" />
                    <button
                        onClick={onValidate}
                        className="flex items-center gap-1 bg-red-500 hover:bg-red-600 active:bg-red-700 text-white px-4 py-2 rounded-lg shadow transition-all duration-150 font-semibold border border-red-400 hover:scale-105"
                    >
                        Validate
                    </button>
                </nav>
                <div className="ml-8">
                    {/* Inline SessionBar UI here for top menu */}
                    <div className="flex items-center gap-3">
                        <span className={`w-3 h-3 rounded-full shadow ${isConnected ? 'bg-green-400 animate-pulse' : 'bg-gray-400'}`}></span>
                        <span className="font-bold text-base text-white flex items-center gap-1">
                            <span className="material-icons text-white text-lg">group</span>
                            Session:
                        </span>
                        <span className="text-white font-semibold text-base select-none">
                            {sessionName ? sessionName : <span className="italic text-gray-200">No session</span>}
                        </span>
                        {isConnected ? (
                            <button className="flex items-center gap-1 bg-red-500 hover:bg-red-600 active:bg-red-700 text-white px-3 py-1 rounded-lg shadow font-semibold border border-red-400 transition-all duration-150 hover:scale-105 ml-2" onClick={onLeaveSession}>
                                離開
                            </button>
                        ) : (
                            <button className="flex items-center gap-1 bg-green-500 hover:bg-green-600 active:bg-green-700 text-white px-3 py-1 rounded-lg shadow font-semibold border border-green-400 transition-all duration-150 hover:scale-105 ml-2" onClick={onJoinSession}>
                                加入/建立
                            </button>
                        )}
                    </div>
                </div>
            </div>
        </header>
    );
};

export default TopMenu;
