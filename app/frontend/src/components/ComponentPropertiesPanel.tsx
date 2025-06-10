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
        <div 
            className="absolute right-0 top-0 w-[320px] h-full shadow-2xl border-l-2 border-[#1C1C1C] overflow-y-auto"
            style={{
                background: 'linear-gradient(180deg, #333333 0%, #1C1C1C 50%, #333333 100%)',
                boxShadow: 'inset 1px 0 0 rgba(242, 242, 240, 0.1), inset -1px 0 0 rgba(0, 0, 0, 0.3)'
            }}
        >
            {/* Industrial texture overlay */}
            <div 
                className="absolute inset-0 opacity-5"
                style={{
                    backgroundImage: `url("data:image/svg+xml,%3Csvg width='40' height='40' viewBox='0 0 40 40' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='%23F2F2F0' fill-opacity='0.1'%3E%3Cpath d='M20 20c0 11.046-8.954 20-20 20v-40c11.046 0 20 8.954 20 20zM0 0h40v40H0V0z'/%3E%3C/g%3E%3C/svg%3E")`,
                }}
            />
            
            <div className="relative z-10">
                {(selectedComponent as GadgetProps).gadgetType !== undefined ? (
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
                )}
            </div>
        </div>
    );
};

export default ComponentPropertiesPanel;
