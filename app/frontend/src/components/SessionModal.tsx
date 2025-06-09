import React from "react";

interface SessionModalProps {
    open: boolean;
    onConfirm: (name: string) => void;
    onClose: () => void;
}

const SessionModal: React.FC<SessionModalProps> = ({ open, onConfirm, onClose }) => {
    if (!open) return null;
    return (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-40 z-50">
            <div className="bg-white rounded shadow-lg p-6 w-80">
                <h2 className="text-lg font-bold mb-4">加入/建立 Session</h2>
                <input
                    id="session-input"
                    type="text"
                    className="border rounded px-2 py-1 w-full mb-4"
                    placeholder="輸入 Session 名稱"
                    onKeyDown={e => {
                        if (e.key === 'Enter') {
                            onConfirm((e.target as HTMLInputElement).value);
                        }
                    }}
                />
                <div className="flex justify-end gap-2">
                    <button className="px-3 py-1 rounded bg-gray-300" onClick={onClose}>取消</button>
                    <button className="px-3 py-1 rounded bg-blue-600 text-white" onClick={() => {
                        const input = document.querySelector<HTMLInputElement>("#session-input");
                        if (input) onConfirm(input.value);
                    }}>確認</button>
                </div>
            </div>
        </div>
    );
};

export default SessionModal;
