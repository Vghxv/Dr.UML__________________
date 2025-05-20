import React, {useEffect, useState} from "react";
import {offBackendEvent, onBackendEvent, ToPoint} from "./utils/wailsBridge";

import {AddGadget, GetCurrentDiagramName, GetDrawData} from "../wailsjs/go/umlproject/UMLProject";
import { mockAssociation, mockSelfAssociation, mockHorizontalAssociation, mockVerticalAssociation, mockSelfAssociationLeft, mockSelfAssociationUp} from "./assets/mock/ass";

import {CanvasProps} from "./utils/Props";
import DrawingCanvas from "./components/Canvas";
// import mockData from './assets/mock/gadget';
import {GadgetPopup} from "./components/CreateGadgetPopup";
import Toolbar from "./components/Toolbar";
import { createAss } from "./utils/createAssociation";

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

    // 新增 addAss handler
    const handleAddAss = async () => {
        // 直接將 mockAssociation 加入 backendData.Association
        setBackendData(prev => {
            if (!prev) return prev;
            const newAss = prev.Association ? [...prev.Association, mockSelfAssociationUp] : [mockSelfAssociationUp];
            return { ...prev, Association: newAss };
        });
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
            <h1 style={{fontFamily: "Inkfree"}}>Dr.UML</h1>
            <Toolbar
                onGetDiagramName={handleGetDiagramName}
                onShowPopup={() => setShowPopup(true)}
                onAddAss={handleAddAss}
                diagramName={diagramName}
            />
            {showPopup && (
                <GadgetPopup
                    isOpen={showPopup}
                    onCreate={(gadget) => {
                        console.log("New Gadget Created:", gadget);
                        setShowPopup(false);
                    }}
                    onClose={() => setShowPopup(false)}
                />
            )}
            <DrawingCanvas backendData={backendData}/>
        </div>
    );
};

export default App;
