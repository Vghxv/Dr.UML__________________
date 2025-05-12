import React, { useState, useEffect } from "react";
import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";
import "./App.css";
import Canvas from "./components/Canvas";
import {
    getCurrentDiagramName,
    addGadget,
    onBackendEvent,
    offBackendEvent,
} from "./utils/wailsBridge";
import { BackendCanvasProps } from "./utils/createCanvas";
import mockData from './assets/mock/gadget';
import CreateGadgetPopup from "./components/CreateGadgetPopup";

const App: React.FC = () => {
    const [diagramName, setDiagramName] = useState<string | null>(null);
    const [backendData, setBackendData] = useState<BackendCanvasProps | null>(null);
    const [showPopup, setShowPopup] = useState(false);

    const handleGetDiagramName = async () => {
        try {
            const name = await getCurrentDiagramName();
            setDiagramName(name);
        } catch (error) {
            console.error("Error fetching diagram name:", error);
        }
    };

    const handleAddGadget = async () => {
        try {
            // Generate random positions between 50 and 500
            const randomX = Math.floor(Math.random() * 450) + 50;
            const randomY = Math.floor(Math.random() * 450) + 50;
            await addGadget(1, { x: randomX, y: randomY }, 0, 0x0000FF);
        } catch (error) {
            console.error("Error adding gadget:", error);
        }
    };

    useEffect(() => {
        onBackendEvent("backend-event", (result) => {
            console.log("Received data from backend:", result);
            setBackendData(result);
        });

        return () => {
            offBackendEvent("backend-event");
        };
    }, []);

    return (
        <DndProvider backend={HTML5Backend}>
            <div className="section">
                <h1 style={{ fontFamily: "Inkfree" }}>Dr.UML</h1>

                <div style={{ marginBottom: "10px" }}>
                    <button className="btn" onClick={handleGetDiagramName}>
                        Get Diagram Name
                    </button>
                    {diagramName && <p>Diagram Name: {diagramName}</p>}
                </div>

                <div style={{ marginBottom: "10px" }}>
                    <button className="btn" onClick={handleAddGadget}>
                        Load Gadget From Backend
                    </button>
                </div>
            </div>

            <div style={{ marginBottom: "10px" }}>
                <button className="btn" onClick={() => setShowPopup(true)}>
                    + Create Gadget (Popup)
                </button>
            </div>

            {showPopup && (
                <CreateGadgetPopup
                    onCreate={(gadget) => {
                        // 這裡可以將 gadget 傳給後端或存在 local state
                        console.log("New Gadget Created:", gadget);
                        setShowPopup(false); // 關閉 popup
                    }}
                    onCancel={() => setShowPopup(false)}
                />
            )}


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
                {backendData && <Canvas backendData={backendData} />}
            </div>
        </DndProvider>
    );
};

export default App;
