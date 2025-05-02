import React, { useState, useEffect } from "react";
import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";
import "./App.css";
import { dia } from "@joint/core";
import Gadget from "./components/Gadget";
import Association from "./components/Association";
import Canvas from "./components/Canvas";
import ChatRoom from "./components/ChatRoom";
import { parseBackendGadget } from "./utils/createGadget";
import { EventsOn, EventsOff, EventsOnce } from "../wailsjs/runtime";

interface WindowWithGo extends Window {
  go: {
    umlproject: {
      UMLProject: {
        GetCurrentDiagramName(): Promise<string>;
        AddNewDiagram(diagramType: number, name: string): Promise<void>;
        SelectDiagram(name: string): Promise<void>;
        AddGadget(
          gadgetType: number,
          point: { x: number; y: number }
        ): Promise<void>;
      };
    };
  };
}

declare var window: WindowWithGo;

const App: React.FC = () => {
  const [graph] = useState(new dia.Graph()); // Create a new JointJS graph instance

  const handleCreateAssociation = (association: dia.Link) => {
    graph.addCell(association); // Add the association to the graph
  };

  const [diagramName, setDiagramName] = useState<any>(null);

  const handleGetDiagramName = async () => {
    try {
      const name =
        await window.go.umlproject.UMLProject.GetCurrentDiagramName();
      setDiagramName(name);
      console.log("diagram name:", name);
    } catch (error) {
      console.error("Error calling Go function:", error);
    }
  };

  const handleAddGadget = async () => {
    try {
      await window.go.umlproject.UMLProject.AddGadget(1, { x: 100, y: 100 });
      console.log("handleAddGadget");
    } catch (error) {
      console.error("Error calling Go function:", error);
    }
  };

  useEffect(() => {
    // Register the event listener
    EventsOn("backend-event", (result) => {
      // setCallbackResult(result);
      const components = result["components"]["gadgets"];
      console.log("components", components);
      for (const gadget_data of components) {
        console.log("gadget data in for", gadget_data);
        const gadget = parseBackendGadget(gadget_data);
        if (gadget) {
          graph.addCell(gadget);
          console.log("gadget added", gadget);
        }
      }
      console.log("=======================end=======================");
    });

    // Clean up the event listener when the component unmounts
    return () => {
      EventsOff("backend-event");
    };
  }, ["backend-event"]);

  return (
    <DndProvider backend={HTML5Backend}>
      <div className="section">
        <h1>Dr.UML</h1>
        <div style={{ marginBottom: "10px" }}>
          <button className="btn" onClick={handleGetDiagramName}>
            handleGetDiagramName
          </button>
          {<p>Diagram Name: {diagramName}</p>}
        </div>
      </div>
      <div
        className="App"
        style={{
          display: "flex",
          flexDirection: "row", // Layout with left and right sections
          justifyContent: "space-between",
          alignItems: "flex-start",
          height: "100vh",
          backgroundColor: "#121212",
          padding: "20px",
        }}
      >
        {/* Left Section: Gadget Palette and Association Tool */}
        <div
          style={{
            width: "300px", // Fixed width for the left section
            marginRight: "20px", // Add spacing between left section and Canvas
            display: "flex",
            flexDirection: "column",
            gap: "20px", // Add spacing between components
          }}
        >
          <div>
            <h2 style={{ color: "#ffffff", marginBottom: "10px" }}>Gadget Palette</h2>
            <div style={{ display: "flex", gap: "10px", flexWrap: "wrap" }}>
              <Gadget
                point={{ x: 200, y: 200 }}
                type="Class"
                layer={1}
                name="Class Gadget"
                onDrop={handleAddGadget}
              />
            </div>
          </div>

          <div>
            <h2 style={{ color: "#ffffff", marginBottom: "10px" }}>Association Tool</h2>
            <Association
              source={{ x: 100, y: 100 }}
              target={{ x: 300, y: 300 }}
              layer={1}
              style={{ stroke: "#FF5733", strokeWidth: 3 }}
              marker={{
                type: "path",
                d: "M 10 -5 0 0 10 5 Z",
                fill: "#FF5733",
              }}
              onCreate={handleCreateAssociation}
            />
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
          <Canvas graph={graph} />
        </div>

        {/* Right Section: ChatRoom */}
        <div
          style={{
            width: "300px", // Fixed width for ChatRoom
            marginLeft: "20px", // Add spacing between Canvas and ChatRoom
          }}
        >
          <ChatRoom />
        </div>
      </div>
    </DndProvider>
  );
};

export default App;
