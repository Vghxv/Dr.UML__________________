import React, { useEffect, useRef, useState } from 'react';
import { useDrop } from 'react-dnd';
import { ItemTypes } from '../types';
import { dia, shapes } from '@joint/core';
import ShapeEditor from './ShapeEditor';
import SessionManager from '../utils/SessionManager';

const Canvas: React.FC<{ graph: dia.Graph }> = ({ graph }) => {
  const canvasRef = useRef<HTMLDivElement>(null);
  const paperRef = useRef<dia.Paper | null>(null);
  const [selectedShape, setSelectedShape] = useState<dia.Element | null>(null);

  useEffect(() => {
    const sessionManager = SessionManager.getInstance();
    const savedGraph = sessionManager.getItem<dia.Graph.Options>('graph');
    if (savedGraph) {
      graph.fromJSON(savedGraph); // Restore graph state from session
    }

    if (canvasRef.current && !paperRef.current) {
      paperRef.current = new dia.Paper({
        el: canvasRef.current,
        model: graph,
        width: canvasRef.current.offsetWidth,
        height: canvasRef.current.offsetHeight,
        background: { color: '#F5F5F5' },
        cellViewNamespace: shapes,
      });

      paperRef.current.on('element:pointerdown', (elementView) => {
        setSelectedShape(elementView.model);
      });
    }

    return () => {
      sessionManager.setItem('graph', graph.toJSON()); // Save graph state on unmount
    };
  }, [graph]);

  const [, drop] = useDrop({
    accept: ItemTypes.SHAPE,
    drop: (item: { type: string }, monitor) => {
      const offset = monitor.getClientOffset();
      if (offset && paperRef.current) {
        const { x, y } = offset;
        const localPoint = paperRef.current.clientToLocalPoint({ x, y });

        let element;
        if (item.type === 'Rectangle') {
          element = new shapes.standard.Rectangle();
          element.attr({
            body: { fill: '#3498db' },
          });
        } else if (item.type === 'Circle') {
          element = new shapes.standard.Circle();
          element.attr({
            body: { fill: '#e74c3c' },
          });
        }

        if (element) {
          element.position(localPoint.x, localPoint.y);
          element.resize(100, 40);
          element.addTo(graph);
        }
      }
    },
  });

  return (
    <div
      ref={(node) => {
        drop(node);
        (canvasRef.current as HTMLDivElement | null) = node;
      }}
      id="paper"
      style={{
        flex: 1,
        height: '100vh',
        backgroundColor: '#ffffff',
        position: 'relative',
        zIndex: 1,
        overflow: 'hidden',
        border: '1px solid #ddd',
      }}
    >
      {selectedShape && (
        <ShapeEditor
          shape={selectedShape}
          onClose={() => setSelectedShape(null)}
        />
      )}
    </div>
  );
};

export default Canvas;
