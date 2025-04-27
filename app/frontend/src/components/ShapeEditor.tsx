import React, { useState, useEffect, useRef } from 'react';
import { dia } from '@joint/core';

interface ShapeEditorProps {
  shape: dia.Element;
  onClose: () => void;
}

const ShapeEditor: React.FC<ShapeEditorProps> = ({ shape, onClose }) => {
  const [fillColor, setFillColor] = useState<string>(shape.attr('body/fill') || '#ffffff');
  const [label, setLabel] = useState<string>(shape.attr('label/text') || '');
  const [width, setWidth] = useState<number>(shape.size().width);
  const [height, setHeight] = useState<number>(shape.size().height);
  const [x, setX] = useState<number>(shape.position().x);
  const [y, setY] = useState<number>(shape.position().y);
  const labelInputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    // Automatically focus on the label input field when the component is mounted
    if (labelInputRef.current) {
      labelInputRef.current.focus();
    }
  }, []);

  const handleSave = () => {
    shape.attr({
      body: { fill: fillColor },
      label: { text: label },
    });
    shape.resize(width, height);
    shape.position(x, y);
    onClose();
  };

  return (
    <div
      style={{
        position: 'absolute',
        top: '50%',
        left: '50%',
        transform: 'translate(-50%, -50%)',
        backgroundColor: '#f9f9f9', // Light gray background for better contrast
        padding: '20px',
        boxShadow: '0 4px 12px rgba(0, 0, 0, 0.3)', // Darker shadow for depth
        borderRadius: '10px', // Slightly more rounded corners
        zIndex: 1000,
        border: '1px solid #ddd', // Subtle border for definition
      }}
    >
      <h3 style={{ color: '#333', marginBottom: '15px' }}>Edit Shape</h3> {/* Darker text color */}
      <div style={{ marginBottom: '10px' }}>
        <label style={{ color: '#555' }}>Fill Color:</label> {/* Subtle label color */}
        <input
          type="color"
          value={fillColor}
          onChange={(e) => setFillColor(e.target.value)}
          style={{
            marginLeft: '10px',
            border: '1px solid #ccc',
            borderRadius: '4px',
          }}
        />
      </div>
      <div style={{ marginBottom: '10px' }}>
        <label style={{ color: '#555' }}>Label:</label>
        <input
          ref={labelInputRef}
          type="text"
          value={label}
          onChange={(e) => setLabel(e.target.value)}
          style={{
            marginLeft: '10px',
            border: '1px solid #ccc',
            borderRadius: '4px',
            padding: '5px',
          }}
        />
      </div>
      <div style={{ marginBottom: '10px' }}>
        <label style={{ color: '#555' }}>Width:</label>
        <input
          type="number"
          value={width}
          onChange={(e) => setWidth(Number(e.target.value))}
          style={{
            marginLeft: '10px',
            width: '60px',
            border: '1px solid #ccc',
            borderRadius: '4px',
            padding: '5px',
          }}
        />
      </div>
      <div style={{ marginBottom: '10px' }}>
        <label style={{ color: '#555' }}>Height:</label>
        <input
          type="number"
          value={height}
          onChange={(e) => setHeight(Number(e.target.value))}
          style={{
            marginLeft: '10px',
            width: '60px',
            border: '1px solid #ccc',
            borderRadius: '4px',
            padding: '5px',
          }}
        />
      </div>
      <div style={{ marginBottom: '10px' }}>
        <label style={{ color: '#555' }}>X Position:</label>
        <input
          type="number"
          value={x}
          onChange={(e) => setX(Number(e.target.value))}
          style={{
            marginLeft: '10px',
            width: '60px',
            border: '1px solid #ccc',
            borderRadius: '4px',
            padding: '5px',
          }}
        />
      </div>
      <div style={{ marginBottom: '10px' }}>
        <label style={{ color: '#555' }}>Y Position:</label>
        <input
          type="number"
          value={y}
          onChange={(e) => setY(Number(e.target.value))}
          style={{
            marginLeft: '10px',
            width: '60px',
            border: '1px solid #ccc',
            borderRadius: '4px',
            padding: '5px',
          }}
        />
      </div>
      <button
        onClick={handleSave}
        style={{
          marginRight: '10px',
          backgroundColor: '#4CAF50', // Green save button
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
        onClick={onClose}
        style={{
          backgroundColor: '#f44336', // Red cancel button
          color: '#fff',
          border: 'none',
          borderRadius: '4px',
          padding: '8px 12px',
          cursor: 'pointer',
        }}
      >
        Cancel
      </button>
    </div>
  );
};

export default ShapeEditor;
