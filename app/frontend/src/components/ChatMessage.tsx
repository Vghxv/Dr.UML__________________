import React from 'react';

interface ChatMessageProps {
  user: string;
  text: string;
}

const ChatMessage: React.FC<ChatMessageProps> = ({ user, text }) => {
  return (
    <div
      style={{
        marginBottom: '10px',
        padding: '10px',
        backgroundColor: '#f9f9f9', // Light gray background for better readability
        borderRadius: '8px', // Rounded corners
        boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)', // Subtle shadow for depth
      }}
    >
      <strong style={{ color: '#333' }}>{user}:</strong>{' '}
      <span style={{ color: '#555' }}>{text}</span> {/* Subtle text color */}
    </div>
  );
};

export default ChatMessage;
