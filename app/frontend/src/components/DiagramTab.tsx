import React, {useState} from 'react';
import Canvas from './Canvas';
import {dia} from '@joint/core';

interface DiagramTabProps {
    diagrams: { id: string; name: string; graph: dia.Graph }[];
    onAddDiagram: () => void;
    onRemoveDiagram: (id: string) => void;
}

const DiagramTab: React.FC<DiagramTabProps> = ({diagrams, onAddDiagram, onRemoveDiagram}) => {
    const [activeDiagramId, setActiveDiagramId] = useState<string>(diagrams[0]?.id || '');

    const handleTabClick = (id: string) => {
        setActiveDiagramId(id);
    };

    return (
        <div style={{display: 'flex', flexDirection: 'column', height: '100%'}}>
            <div
                style={{
                    display: 'flex',
                    borderBottom: '1px solid #ddd',
                    backgroundColor: '#f9f9f9',
                    padding: '5px',
                }}
            >
                {diagrams.map((diagram) => (
                    <div
                        key={diagram.id}
                        onClick={() => handleTabClick(diagram.id)}
                        style={{
                            padding: '10px 15px',
                            cursor: 'pointer',
                            backgroundColor: activeDiagramId === diagram.id ? '#e6e6fa' : '#fff',
                            border: activeDiagramId === diagram.id ? '1px solid #ccc' : '1px solid transparent',
                            borderRadius: '4px 4px 0 0',
                            marginRight: '5px',
                        }}
                    >
                        {diagram.name}
                        <button
                            onClick={(e) => {
                                e.stopPropagation();
                                onRemoveDiagram(diagram.id);
                            }}
                            style={{
                                marginLeft: '10px',
                                backgroundColor: '#f44336',
                                color: '#fff',
                                border: 'none',
                                borderRadius: '50%',
                                width: '20px',
                                height: '20px',
                                cursor: 'pointer',
                                fontSize: '12px',
                                lineHeight: '20px',
                                textAlign: 'center',
                            }}
                        >
                            ×
                        </button>
                    </div>
                ))}
                <button
                    onClick={onAddDiagram}
                    style={{
                        padding: '10px 15px',
                        cursor: 'pointer',
                        backgroundColor: '#4CAF50',
                        color: '#fff',
                        border: 'none',
                        borderRadius: '4px',
                        marginLeft: 'auto',
                    }}
                >
                    + Add Diagram
                </button>
            </div>
            <div style={{flex: 1, position: 'relative'}}>
                {diagrams.map((diagram) =>
                    diagram.id === activeDiagramId ? (
                        <Canvas key={diagram.id} graph={diagram.graph}/>
                    ) : null
                )}
            </div>
        </div>
    );
};

export default DiagramTab;
