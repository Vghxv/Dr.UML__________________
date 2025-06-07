import React from 'react';

interface ChatMessageProps {
    user: string;
    text: string;
}
// TODO: integrate with to App
const ChatMessage: React.FC<ChatMessageProps> = ({user, text}) => {
    return (
        <div className="mb-2.5 p-2.5 bg-gray-100 rounded-lg shadow-sm">
            <strong className="text-gray-800">{user}:</strong>{' '}
            <span className="text-gray-700">{text}</span>
            <div className="text-xs text-gray-500 mt-1.5">
                {new Date().toLocaleTimeString([], {hour: '2-digit', minute: '2-digit'})}
            </div>
        </div>
    );
};

export default ChatMessage;
