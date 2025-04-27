import React, { useState } from 'react';
import { useDrag } from 'react-dnd';
import { ItemTypes } from '../types';

const PaletteItem: React.FC<{ type: string }> = ({ type }) => {
  const [{ isDragging }, drag] = useDrag({
    type: ItemTypes.SHAPE,
    item: { type },
    collect: (monitor) => ({
      isDragging: monitor.isDragging(),
    }),
  });

  return (
    <div
      ref={drag}
      className="palette-item"
      style={{
        opacity: isDragging ? 0.5 : 1,
        cursor: 'move',
        margin: '10px',
        padding: '10px',
        border: '1px solid #888', // Darker border for better contrast
        backgroundColor: '#f0f8ff', // Light blue background
        borderRadius: type === 'Circle' ? '50%' : '4px', // Circle for Circle type
        width: '60px', // Fixed width for shape preview
        height: '60px', // Fixed height for shape preview
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
      }}
    >
      {type}
    </div>
  );
};

const SidePalette: React.FC<{ graph: any }> = () => {
  const [isCollapsed, setIsCollapsed] = useState(false);

  return (
    <div
      className="side-palette"
      style={{
        width: isCollapsed ? '50px' : '200px', // Adjust width based on state
        backgroundColor: '#e6e6fa', // Lavender background
        padding: '10px',
        boxShadow: '2px 0 5px rgba(0,0,0,0.2)', // Slightly darker shadow
        borderRight: '2px solid #ccc', // Add a border to separate from Canvas
        transition: 'width 0.3s ease', // Smooth transition for collapse/expand
        overflow: 'hidden', // Hide content when collapsed
      }}
    >
      <button
        onClick={() => setIsCollapsed(!isCollapsed)}
        style={{
          marginBottom: '10px',
          padding: '5px',
          cursor: 'pointer',
          backgroundColor: '#888',
          color: '#fff',
          border: 'none',
          borderRadius: '4px',
          width: '100%',
        }}
      >
        {isCollapsed ? 'Expand' : 'Collapse'}
      </button>
      {!isCollapsed && (
        <>
          <h3>Shapes</h3>
          <PaletteItem type="Rectangle" />
          <PaletteItem type="Circle" />
        </>
      )}
    </div>
  );
};

export default SidePalette;
