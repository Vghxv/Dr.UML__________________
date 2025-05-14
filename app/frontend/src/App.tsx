import React, {useEffect, useState} from "react";
import "./App.css";
import {offBackendEvent, onBackendEvent, ToPoint} from "./utils/wailsBridge";

import {AddGadget, GetCurrentDiagramName, GetDrawData} from "../wailsjs/go/umlproject/UMLProject";


import {CanvasProps} from "./utils/Props";
import DrawingCanvas from "./components/Canvas";
// import mockData from './assets/mock/gadget';
import {GadgetPopup} from "./components/CreateGadgetPopup";

const App: React.FC = () => {
    const [diagramName, setDiagramName] = useState<string | null>(null);
    const [backendData, setBackendData] = useState<CanvasProps | null>(null);
    const [showPopup, setShowPopup] = useState(false);

    const handleGetDiagramName = async () => {
        try {
            const name = await GetCurrentDiagramName();
            setDiagramName(name);
        } catch (error) {
            console.error("Error fetching diagram name:", error);
        }
    };

    // TODO: remove this when CreateGadgetPopup is done
    const handleAddGadget = async () => {
        try {
            // Generate random positions between 50 and 500
            const randomX = Math.floor(Math.random() * 450) + 50;
            const randomY = Math.floor(Math.random() * 450) + 50;
            // TODO: import gadget type from backend
            await AddGadget(1, ToPoint(randomX, randomY), 0, 0x00FF00, "sample header");
        } catch (error) {
            console.error("Error adding gadget:", error);
        }
    };

    const loadCanvasData = async () => {
        try {
            const diagram = await GetDrawData();
            const canvasData: CanvasProps = {
                margin: diagram.margin,
                color: diagram.color,
                lineWidth: diagram.lineWidth,
                gadgets: diagram.gadgets.map(gadget => ({
                    gadgetType: gadget.gadgetType.toString(),
                    x: gadget.x,
                    y: gadget.y,
                    layer: gadget.layer,
                    height: gadget.height,
                    width: gadget.width,
                    color: gadget.color,
                    isSelected: gadget.isSelected,
                    attributes: gadget.attributes
                }))
            };
            setBackendData(canvasData);
        } catch (error) {
            console.error("Error loading canvas data:", error);
        }
    };

    useEffect(() => {
        // Load canvas data when the component mounts
        loadCanvasData().then(r => console.log("Loaded canvas data:", r));

        onBackendEvent("backend-event", (result) => {
            console.log("Received data from backend:", result);
            // Convert the received data to CanvasProps format
            if (result) {
                const canvasData: CanvasProps = {
                    margin: result.margin,
                    color: result.color,
                    lineWidth: result.lineWidth,
                    gadgets: result.gadgets?.map((gadget: {
                        gadgetType: { toString: () => any; };
                        x: any;
                        y: any;
                        layer: any;
                        height: any;
                        width: any;
                        color: any;
                        isSelected: any;
                        attributes: any;
                    }) => ({
                        gadgetType: gadget.gadgetType.toString(),
                        x: gadget.x,
                        y: gadget.y,
                        layer: gadget.layer,
                        height: gadget.height,
                        width: gadget.width,
                        color: gadget.color,
                        isSelected: gadget.isSelected,
                        attributes: gadget.attributes
                    }))
                };
                setBackendData(canvasData);
            }
        });

        return () => {
            offBackendEvent("backend-event");
        };
    }, []);

    return (
        <div>
            <div className="section">
                <h1 style={{fontFamily: "Inkfree"}}>Dr.UML</h1>

                <div style={{marginBottom: "10px"}}>
                    <button className="btn" onClick={handleGetDiagramName}>
                        Get Diagram Name
                    </button>
                    {diagramName && <p>Diagram Name: {diagramName}</p>}
                </div>

                <div style={{marginBottom: "10px"}}>
                    <button className="btn" onClick={handleAddGadget}>
                        Add New Gadget
                    </button>
                </div>
            </div>

            <div style={{marginBottom: "10px"}}>
                <button className="btn" onClick={() => setShowPopup(true)}>
                    + Create Gadget (Popup)
                </button>
            </div>
            <div>
                {showPopup && (
                    <GadgetPopup
                        isOpen={showPopup}
                        onCreate={(gadget: {
                            id: number;
                            name: string;
                            position: { x: number; y: number };
                            color: string
                        }) => {
                            console.log("New Gadget Created:", gadget);
                            setShowPopup(false);
                        }}
                        onClose={() => setShowPopup(false)}
                    />
                )}
            </div>
            <DrawingCanvas backendData={backendData}/>
        </div>

    );
};

export default App;
