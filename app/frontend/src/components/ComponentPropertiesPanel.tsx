import React from "react";
import { AssociationProps, GadgetProps } from "../utils/Props";
import GadgetPropertiesPanel from "./GadgetPropertiesPanel";
import AssociationPropertiesPanel from "./AssociationPropertiesPanel";

interface ComponentPropertiesPanelProps {
    selectedComponent: GadgetProps | AssociationProps | null;
    updateComponentProperty: (property: string, value: any) => void;
    addAttributeToComponent: (section: number, content: string) => void;
}

const ComponentPropertiesPanel: React.FC<ComponentPropertiesPanelProps> = ({
    selectedComponent,
    updateComponentProperty,
    addAttributeToComponent
}) => {
    if (!selectedComponent) return null;

    // Type guard for GadgetProps
    function isGadgetProps(obj: any): obj is GadgetProps {
        return obj && typeof obj.gadgetType === "string";
    }

    return (
        <div
            style={{
                position: "fixed", // fixed
                right: 0,
                top: 70, // lower than the top menu
                width: 320,
                height: "calc(100vh - 70px)",
                background: "#f3f4f6",
                boxShadow: "-2px 0 8px rgba(0,0,0,0.08)",
                overflowY: "auto",
                zIndex: 60
            }}
        >
            {isGadgetProps(selectedComponent) ? (
                <GadgetPropertiesPanel
                    selectedGadget={selectedComponent}
                    updateGadgetProperty={updateComponentProperty}
                    addAttributeToGadget={addAttributeToComponent}
                />
            ) : (
                <AssociationPropertiesPanel
                    selectedAssociation={selectedComponent}
                    updateAssociationProperty={updateComponentProperty}
                />
            )}
        </div>
    );
};

export default ComponentPropertiesPanel;
