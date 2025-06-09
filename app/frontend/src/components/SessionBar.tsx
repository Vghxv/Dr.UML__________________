import React from "react";

interface SessionBarProps {
    sessionName: string | null;
    isConnected: boolean;
    onJoinSession: () => void;
    onLeaveSession: () => void;
}

const SessionBar: React.FC<SessionBarProps> = ({ sessionName, isConnected, onJoinSession, onLeaveSession }) => {
    return (
        <div className="flex items-center justify-between backdrop-blur-md bg-white/20 border border-white/30 shadow-lg rounded-2xl px-6 py-2 mb-4 mt-2">
            <div className="flex items-center gap-3">
                <span className={`w-3 h-3 rounded-full shadow ${isConnected ? 'bg-green-400 animate-pulse' : 'bg-gray-400'}`}></span>
                <span className="font-bold text-base text-blue-900 flex items-center gap-1">
                    <span className="material-icons text-blue-700 text-lg">group</span>
                    Session:
                </span>
                <span className="text-blue-900 font-semibold text-base select-none">
                    {sessionName ? sessionName : <span className="italic text-gray-500">No session</span>}
                </span>
            </div>
            <div>
                {isConnected ? (
                    <button className="flex items-center gap-1 bg-red-500 hover:bg-red-600 active:bg-red-700 text-white px-4 py-1.5 rounded-lg shadow font-semibold border border-red-400 transition-all duration-150 hover:scale-105" onClick={onLeaveSession}>
                        <span className="material-icons text-base">logout</span>
                        離開 Session
                    </button>
                ) : (
                    <button className="flex items-center gap-1 bg-green-500 hover:bg-green-600 active:bg-green-700 text-white px-4 py-1.5 rounded-lg shadow font-semibold border border-green-400 transition-all duration-150 hover:scale-105" onClick={onJoinSession}>
                        <span className="material-icons text-base">login</span>
                        加入/建立 Session
                    </button>
                )}
            </div>
        </div>
    );
};

export default SessionBar;
