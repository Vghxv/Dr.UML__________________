import React from "react";
import { AssociationProps, GadgetProps } from "../utils/Props";
import GadgetPropertiesPanel from "./GadgetPropertiesPanel";
import AssociationPropertiesPanel from "./AssociationPropertiesPanel";

interface ComponentPropertiesPanelProps {
    selectedComponent: GadgetProps | AssociationProps | null;
    updateGadgetProperty: (property: string, value: any) => void;
    updateAssociationProperty: (property: string, value: any) => void;
    addAttributeToGadget: (section: number, content: string) => void;
    addAttributeToAssociation: (ratio: number, content: string) => void;
}

const ComponentPropertiesPanel: React.FC<ComponentPropertiesPanelProps> = ({
    selectedComponent,
    updateGadgetProperty,
    updateAssociationProperty,
    addAttributeToGadget,
    addAttributeToAssociation
}) => {
    if (!selectedComponent) return null;
    return (

        (selectedComponent as GadgetProps).gadgetType !== undefined
            ? (
                
                <GadgetPropertiesPanel
                    selectedGadget={selectedComponent as GadgetProps}
                    updateGadgetProperty={updateGadgetProperty}
                    addAttributeToGadget={addAttributeToGadget}
                />
            ) : (
                <AssociationPropertiesPanel
                    selectedAssociation={selectedComponent as AssociationProps}
                    updateAssociationProperty={updateAssociationProperty}
                    addAttributeToAssociation={addAttributeToAssociation}
                />
            )

    )
};

export default ComponentPropertiesPanel;
