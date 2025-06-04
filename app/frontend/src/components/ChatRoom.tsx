import React, { useEffect, useState } from 'react';
import ChatMessage from './ChatMessage';

interface Message {
    id: number;
    user: string;
    text: string;
}

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
        <div
            style={{
                display: 'flex',
                flexDirection: 'column',
                height: '400px', // 固定高度
                width: '300px', // 可選：給個寬度比較好看
                border: '1px solid #ddd',
                borderRadius: '8px',
                overflow: 'hidden',
                backgroundColor: '#f9f9f9',
            }}
        >
            {/* 訊息區 */}
            <div
                style={{
                    flex: 1,
                    overflowY: 'auto', // 超出就出現自己的scroll bar
                    padding: '10px',
                }}
            >
                {messages.map((message) => (
                    <ChatMessage key={message.id} user={message.user} text={message.text} />
                ))}
            </div>

            {/* 輸入區 */}
            <div
                style={{
                    display: 'flex',
                    padding: '10px',
                    borderTop: '1px solid #ddd',
                    backgroundColor: '#fff',
                }}
            >
                <input
                    type="text"
                    value={input}
                    onChange={(e) => setInput(e.target.value)}
                    onKeyDown={handleKeyDown}
                    placeholder="Type a message..."
                    style={{
                        flex: 1,
                        padding: '8px',
                        border: '1px solid #ccc',
                        borderRadius: '4px',
                        marginRight: '10px',
                    }}
                />
                <button
                    onClick={handleSend}
                    style={{
                        padding: '8px 12px',
                        backgroundColor: '#4CAF50',
                        color: '#fff',
                        border: 'none',
                        borderRadius: '4px',
                        cursor: 'pointer',
                    }}
                >
                    Send
                </button>
            </div>
        </div>
    );
};

export default ChatRoom;
