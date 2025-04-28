import React from 'react';
import {DndProvider} from 'react-dnd';
import {HTML5Backend} from 'react-dnd-html5-backend';
import './App.css';
import { dia, shapes } from '@joint/core';
import Canvas from './components/Canvas';

const namespace = shapes;

const graph = new dia.Graph({}, { cellNamespace: namespace });

const paper = new dia.Paper({
    el: document.getElementById('paper'),
    model: graph,
    width: 300,
    height: 300,
    background: { color: '#F5F5F5' },
    cellViewNamespace: namespace
});
// Example usage of paper to avoid unused variable error
console.log(paper);

const App: React.FC = () => {
    return (
        <DndProvider backend={HTML5Backend}>
            <div className="App">
                <header className="App-header">
                    <h1>Dr.UML</h1>
                </header>
                <main className="App-main">
                    <div id="paper" style={{ width: '100%', height: '100vh' }}></div>
                    <Canvas graph={graph} />
                </main>
            </div>
        </DndProvider>
    );
};

export default App;
