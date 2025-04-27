import React from 'react';

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
    <div
      style={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        padding: '10px 20px',
        backgroundColor: '#333', // Dark background for the menu
        color: '#fff', // White text for contrast
        boxShadow: '0 2px 5px rgba(0, 0, 0, 0.2)', // Subtle shadow for depth
      }}
    >
      <h1 style={{ margin: 0, fontSize: '18px' }}>Dr.UML</h1>
      <div style={{ display: 'flex', gap: '10px' }}>
        <button
          onClick={handleOpenProject}
          style={{
            backgroundColor: '#4CAF50', // Green button
            color: '#fff',
            border: 'none',
            borderRadius: '4px',
            padding: '8px 12px',
            cursor: 'pointer',
          }}
        >
          Open Project
        </button>
        <button
          onClick={handleSave}
          style={{
            backgroundColor: '#2196F3', // Blue button
            color: '#fff',
            border: 'none',
            borderRadius: '4px',
            padding: '8px 12px',
            cursor: 'pointer',
          }}
        >
          Save
        </button>
        <button
          onClick={handleExport}
          style={{
            backgroundColor: '#FF9800', // Orange button
            color: '#fff',
            border: 'none',
            borderRadius: '4px',
            padding: '8px 12px',
            cursor: 'pointer',
          }}
        >
          Export
        </button>
        <button
          onClick={handleValidate}
          style={{
            backgroundColor: '#f44336', // Red button
            color: '#fff',
            border: 'none',
            borderRadius: '4px',
            padding: '8px 12px',
            cursor: 'pointer',
          }}
        >
          Validate
        </button>
      </div>
    </div>
  );
};

export default TopMenu;
