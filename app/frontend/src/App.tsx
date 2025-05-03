import React, { useState, useEffect } from "react";
import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";
import "./App.css";
import { dia } from "@joint/core";
import Canvas from "./components/Canvas";
import { getCurrentDiagramName, addGadget, onBackendEvent, offBackendEvent } from "./utils/wailsBridge";

const App: React.FC = () => {
    const [graph] = useState(new dia.Graph()); // Create a new JointJS graph instance
    const [diagramName, setDiagramName] = useState<string | null>(null);
    const [backendData, setBackendData] = useState<any>(null);
  
    const handleGetDiagramName = async () => {
        try {
            const name = await getCurrentDiagramName();
            setDiagramName(name);
            console.log("Diagram name:", name);
        } catch (error) {
            console.error("Error fetching diagram name:", error);
        }
    };

    const handleAddGadget = async () => {
        try {
            await addGadget(1, { x: 100, y: 100 });
        } catch (error) {
            console.error("Error adding gadget:", error);
        }
    };

    useEffect(() => {
        // Register the event listener
        onBackendEvent("backend-event", (result) => {
            console.log("Received data from backend:", result);
            const components = result["gadgets"];
            console.log("Components:", components);
            setBackendData(components);
        });

        // Clean up the event listener when the component unmounts
        return () => {
            offBackendEvent("backend-event");
        };
    }, [graph]);

    return (
        <DndProvider backend={HTML5Backend}>
            <div className="section">
                <h1>Dr.UML</h1>
                <div style={{ marginBottom: "10px" }}>
                    <button className="btn" onClick={handleGetDiagramName}>
                        Get Diagram Name
                    </button>
                    {<p>Diagram Name: {diagramName}</p>}
                </div>
                <div style={{ marginBottom: "10px" }}>
                    <button className="btn" onClick={handleAddGadget}>
                        Load Gadget From Backend
                    </button>
                </div>
            </div>


            {/* Center Section: Canvas */}
            <div
                style={{
                    flex: 1,
                    display: "flex",
                    flexDirection: "column",
                    justifyContent: "center",
                    alignItems: "center",
                }}
            >
                <Canvas graph={graph} backendData={backendData} />
            </div>
        </DndProvider>
    );
};

export default App;
