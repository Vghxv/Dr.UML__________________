import React, { useState, useEffect } from 'react';
import { DndProvider } from 'react-dnd';
import { HTML5Backend } from 'react-dnd-html5-backend';
import './App.css';
import { dia } from '@joint/core';
import Canvas from './components/Canvas';
import Gadget from './components/Gadget';
import Association from './components/Association';
import { useParseBackendGadget } from './hooks/useParseBackendGadget';

const App: React.FC = () => {
    const [graph] = useState(new dia.Graph()); // Create a new JointJS graph instance
    const parseBackendGadget = useParseBackendGadget();

    const handleDrop = (gadget: dia.Element) => {
        if (gadget instanceof dia.Element) {
            graph.addCell(gadget); // Add the gadget to the graph
        } else {
            console.error('Invalid gadget type. Must be an instance of dia.Element.');
        }
    };

    const handleCreateAssociation = (association: dia.Link) => {
        graph.addCell(association); // Add the association to the graph
    };

    useEffect(() => {
        // Example backend JSON string
        const backendJson = `{
            "gadgetType": "Class",
            "x": 100,
            "y": 100,
            "layer": 1,
            "height": 120,
            "width": 200,
            "color": 16777215,
            "attributes": [
                [{"content": "id: Int", "height": 20, "width": 100, "fontSize": 12, "fontStyle": 0, "fontFile": ""}],
                [{"content": "+getId(): Int", "height": 20, "width": 100, "fontSize": 12, "fontStyle": 0, "fontFile": ""}]
            ]
        }`;

        const gadget = parseBackendGadget(backendJson);
        if (gadget) {
            graph.addCell(gadget); // Add the parsed gadget to the graph
        }
    }, [graph, parseBackendGadget]);

    return (
        <DndProvider backend={HTML5Backend}>
            <div
                className="App"
                style={{
                    display: 'flex',
                    flexDirection: 'column',
                    justifyContent: 'center',
                    alignItems: 'center',
                    height: '100vh',
                    backgroundColor: '#121212',
                    padding: '20px',
                }}
            >
                <h1
                    style={{
                        color: '#ffffff',
                        marginBottom: '20px',
                        fontFamily: 'Arial, sans-serif',
                        fontSize: '2rem',
                    }}
                >
                    Dr.UML
                </h1>
                <h1>Gadget Palette</h1>
                <div style={{ display: 'flex', gap: '10px', marginBottom: '20px' }}>
                    <Gadget
                        point={{ x: 200, y: 200 }}
                        type="Class"
                        layer={1}
                        name="Class Gadget"
                        onDrop={handleDrop}
                    />
                </div>

                <h1>Association Tool</h1>
                <Association
                    source={{ x: 100, y: 100 }}
                    target={{ x: 300, y: 300 }}
                    layer={1}
                    style={{ stroke: '#FF5733', strokeWidth: 3 }}
                    marker={{
                        type: 'path',
                        d: 'M 10 -5 0 0 10 5 Z',
                        fill: '#FF5733',
                    }}
                    onCreate={handleCreateAssociation}
                />

                <Canvas graph={graph} />
            </div>
        </DndProvider>
    );
};

export default App;
