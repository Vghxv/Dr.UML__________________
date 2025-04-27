import React, { useState } from 'react';
import Canvas from './Canvas';
import { dia } from '@joint/core';

interface DiagramTab {
  id: string;
  name: string;
  graph: dia.Graph;
}

const DiagramTabs: React.FC = () => {
  const [tabs, setTabs] = useState<DiagramTab[]>([
    { id: '1', name: 'Diagram 1', graph: new dia.Graph() },
  ]);
  const [activeTabId, setActiveTabId] = useState<string>('1');

  const addTab = () => {
    const newTabId = (tabs.length + 1).toString();
    setTabs([
      ...tabs,
      { id: newTabId, name: `Diagram ${newTabId}`, graph: new dia.Graph() },
    ]);
    setActiveTabId(newTabId);
  };

  const closeTab = (id: string) => {
    const updatedTabs = tabs.filter((tab) => tab.id !== id);
    setTabs(updatedTabs);
    if (id === activeTabId && updatedTabs.length > 0) {
      setActiveTabId(updatedTabs[0].id);
    }
  };

  return (
    <div style={{ display: 'flex', flexDirection: 'column', height: '100%' }}>
      <div
        style={{
          display: 'flex',
          borderBottom: '1px solid #ddd',
          backgroundColor: '#f5f5f5',
        }}
      >
        {tabs.map((tab) => (
          <div
            key={tab.id}
            style={{
              padding: '10px 15px',
              cursor: 'pointer',
              backgroundColor: tab.id === activeTabId ? '#ffffff' : '#e6e6e6',
              border: tab.id === activeTabId ? '1px solid #ddd' : 'none',
              borderBottom: tab.id === activeTabId ? 'none' : '1px solid #ddd',
            }}
            onClick={() => setActiveTabId(tab.id)}
          >
            {tab.name}
            <button
              onClick={(e) => {
                e.stopPropagation();
                closeTab(tab.id);
              }}
              style={{
                marginLeft: '10px',
                backgroundColor: 'transparent',
                border: 'none',
                cursor: 'pointer',
                color: '#888',
              }}
            >
              ✕
            </button>
          </div>
        ))}
        <button
          onClick={addTab}
          style={{
            marginLeft: 'auto',
            padding: '10px 15px',
            backgroundColor: '#4CAF50',
            color: '#fff',
            border: 'none',
            cursor: 'pointer',
          }}
        >
          + Add Tab
        </button>
      </div>
      <div style={{ flex: 1, position: 'relative' }}>
        {tabs.map(
          (tab) =>
            tab.id === activeTabId && (
              <Canvas key={tab.id} graph={tab.graph} />
            )
        )}
      </div>
    </div>
  );
};

export default DiagramTabs;
