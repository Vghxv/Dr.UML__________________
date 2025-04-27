import React, { useState } from 'react';

const SessionManager: React.FC = () => {
  const [sessionId, setSessionId] = useState('');
  const [isJoining, setIsJoining] = useState(false);

  const handleCreateSession = () => {
    // Logic to create a new session
    console.log('Creating a new session...');
  };

  const handleJoinSession = () => {
    if (sessionId.trim()) {
      // Logic to join an existing session
      console.log(`Joining session with ID: ${sessionId}`);
    } else {
      alert('Please enter a valid session ID.');
    }
  };

  return (
    <div
      style={{
        padding: '20px',
        backgroundColor: '#f9f9f9',
        border: '1px solid #ddd',
        borderRadius: '8px',
        boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)',
        width: '300px',
        margin: '20px auto',
      }}
    >
      <h3 style={{ marginBottom: '15px', color: '#333' }}>Session Manager</h3>
      {!isJoining ? (
        <>
          <button
            onClick={handleCreateSession}
            style={{
              width: '100%',
              padding: '10px',
              backgroundColor: '#4CAF50',
              color: '#fff',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer',
              marginBottom: '10px',
            }}
          >
            Create New Session
          </button>
          <button
            onClick={() => setIsJoining(true)}
            style={{
              width: '100%',
              padding: '10px',
              backgroundColor: '#2196F3',
              color: '#fff',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer',
            }}
          >
            Join Existing Session
          </button>
        </>
      ) : (
        <>
          <input
            type="text"
            value={sessionId}
            onChange={(e) => setSessionId(e.target.value)}
            placeholder="Enter Session ID"
            style={{
              width: '100%',
              padding: '10px',
              marginBottom: '10px',
              border: '1px solid #ccc',
              borderRadius: '4px',
            }}
          />
          <button
            onClick={handleJoinSession}
            style={{
              width: '100%',
              padding: '10px',
              backgroundColor: '#4CAF50',
              color: '#fff',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer',
              marginBottom: '10px',
            }}
          >
            Join Session
          </button>
          <button
            onClick={() => setIsJoining(false)}
            style={{
              width: '100%',
              padding: '10px',
              backgroundColor: '#f44336',
              color: '#fff',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer',
            }}
          >
            Cancel
          </button>
        </>
      )}
    </div>
  );
};

export default SessionManager;
