import React, { useEffect, useState } from 'react';
import ChatMessage from './ChatMessage';

interface Message {
    id: number;
    user: string;
    text: string;
}
// TODO: integrate into App

const ChatRoom: React.FC = () => {
    const [messages, setMessages] = useState<Message[]>([]);
    const [input, setInput] = useState<string>('');
    const [username, setUsername] = useState<string>('User');

    useEffect(() => {
        // Simulate receiving messages from a server
        const interval = setInterval(() => {
            setMessages((prev) => [
                ...prev,
                { id: prev.length + 1, user: 'Server', text: 'Hello from the server!' },
            ]);
        }, 5000);

        return () => clearInterval(interval);
    }, []);

    const handleSend = () => {
        if (input.trim()) {
            setMessages((prev) => [
                ...prev,
                { id: prev.length + 1, user: username, text: input },
            ]);
            setInput('');
        }
    };

    const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key === 'Enter') {
            handleSend();
        }
    };

    return (
        <div className="flex flex-col h-[400px] w-[300px] border border-gray-300 rounded-lg overflow-hidden bg-gray-100">
            {/* Message area */}
            <div className="flex-1 overflow-y-auto p-2.5">
        <div className="flex flex-col h-[400px] w-[300px] border border-gray-300 rounded-lg overflow-hidden bg-gray-100">
            {/* Message area */}
            <div className="flex-1 overflow-y-auto p-2.5">
                {messages.map((message) => (
                    <ChatMessage key={message.id} user={message.user} text={message.text} />
                ))}
            </div>

            {/* Input area */}
            <div className="flex p-2.5 border-t border-gray-300 bg-white">
            {/* Input area */}
            <div className="flex p-2.5 border-t border-gray-300 bg-white">
                <input
                    type="text"
                    value={input}
                    onChange={(e) => setInput(e.target.value)}
                    onKeyDown={handleKeyDown}
                    placeholder="Type a message..."
                    className="flex-1 p-2 border border-gray-300 rounded mr-2.5"
                    className="flex-1 p-2 border border-gray-300 rounded mr-2.5"
                />
                <button
                    onClick={handleSend}
                    className="px-3 py-2 bg-green-600 text-white rounded cursor-pointer hover:bg-green-700"
                    className="px-3 py-2 bg-green-600 text-white rounded cursor-pointer hover:bg-green-700"
                >
                    Send
                </button>
            </div>
        </div>
    );
};

export default ChatRoom;
