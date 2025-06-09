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

    if (isGadgetProps(selectedComponent)) {
        return (
            <GadgetPropertiesPanel
                selectedGadget={selectedComponent}
                updateGadgetProperty={updateComponentProperty}
                addAttributeToGadget={addAttributeToComponent}
            />
        );
    } else {
        return (
            <AssociationPropertiesPanel
                selectedAssociation={selectedComponent}
                updateAssociationProperty={updateComponentProperty}
            />
        );
    }
};

export default ComponentPropertiesPanel;
