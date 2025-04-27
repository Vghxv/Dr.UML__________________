import React from 'react';
import { DndProvider } from 'react-dnd';
import { HTML5Backend } from 'react-dnd-html5-backend';
import './App.css';
import SidePalette from './components/SidePalette';
import TopMenu from './components/TopMenu'; // Import TopMenu
import ChatRoom from './components/ChatRoom'; // Import ChatRoom
import DiagramTabs from './components/DiagramTabs'; // Import DiagramTabs

const App: React.FC = () => {
  return (
    <DndProvider backend={HTML5Backend}>
      <div className="app-container" style={{ display: 'flex', flexDirection: 'column', height: '100vh' }}>
        <TopMenu /> {/* Add TopMenu */}
        <div style={{ display: 'flex', flex: 1 }}>
          <SidePalette graph={null} />
          <DiagramTabs /> {/* Replace Canvas with DiagramTabs */}
          <div style={{ width: '300px', marginLeft: '10px' }}> {/* Add ChatRoom */}
            <ChatRoom />
          </div>
        </div>
      </div>
    </DndProvider>
  );
};

export default App;
