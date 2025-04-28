import React from 'react';
import {DndProvider} from 'react-dnd';
import {HTML5Backend} from 'react-dnd-html5-backend';
import './App.css';
import { dia, shapes } from '@joint/core';
import Canvas from './components/Canvas';

const graph = new dia.Graph();
const App: React.FC = () => {
    // Create a new JointJS graph instance

    return (
        <DndProvider backend={HTML5Backend}>
            <div
                className="App"
                style={{
                    display: 'flex',
                    flexDirection: 'column', // Stack header and canvas vertically
                    justifyContent: 'center',
                    alignItems: 'center',
                    height: '100vh', // Full viewport height
                    backgroundColor: '#121212', // Dark background
                    padding: '20px' // Padding around the content
                }}
            >
                <h1
                    style={{
                        color: '#ffffff', // White text
                        marginBottom: '20px', // Space below the header
                        fontFamily: 'Arial, sans-serif', // Clean font
                        fontSize: '2rem' // Large font size
                    }}
                >
                    Dr.UML
                </h1>
                {/*
                    Memoize the graph instance to avoid recreating it on every render.
                */}
                <Canvas graph={graph} />
            </div>
        </DndProvider>
    );
};

export default App;
