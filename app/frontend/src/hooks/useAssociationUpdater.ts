import {
    AddAttributeToAssociation,
    RemoveAttributeFromAssociation,
    SetAttrContentComponent,
    SetAttrSizeComponent,
    SetAttrStyleComponent,
    SetAttrFontComponent,
    SetLayerComponent
} from "../../wailsjs/go/umlproject/UMLProject";
import { CanvasProps, AssociationProps } from "../utils/Props";
import { createSetSingleValue, createSetDoubleValue, createSetTripleValue } from "./updater";

export function useAssociationUpdater(
    selectedAssociation: AssociationProps | null,
    backendData: CanvasProps | null,
    reloadBackendData: () => void
) {
    // Create bound helper functions using factory functions
    const setSingleValue = createSetSingleValue(reloadBackendData);
    const setDoubleValue = createSetDoubleValue(reloadBackendData);
    const setTripleValue = createSetTripleValue(reloadBackendData);

    const handleAddAttributeToAssociation = (ratio: number, content: string) => {
        if (!selectedAssociation || !backendData || !backendData.associations) return;
        setDoubleValue(
            AddAttributeToAssociation,
            ratio, content,
            "Association attribute added",
            "Error adding association attribute"
        );
    };

    const handleUpdateAssociationProperty = (property: string, value: any) => {

        if (!selectedAssociation || !backendData || !backendData.associations) return;

        if (property.includes('.')) {
            const [parentProp, childProp] = property.split('.');
            if (parentProp.startsWith('attributes')) {
                const matches = parentProp.match(/attributes:(\d+)/);
                if (matches && matches.length === 2) {
                    const index = parseInt(matches[1]);
                    if (childProp === 'content') {
                        setTripleValue(
                            SetAttrContentComponent,
                            0, index, value,  // section is ignored for associations
                            "Association attribute content changed",
                            "Error editing association attribute content"
                        );
                    }
                    if (childProp === 'fontSize') {
                        setTripleValue(
                            SetAttrSizeComponent,
                            0, index, value,  // section is ignored for associations
                            "Association attribute fontSize changed",
                            "Error editing association attribute fontSize"
                        );
                    }
                    if (childProp === 'fontStyleB' || childProp === 'fontStyleI' || childProp === 'fontStyleU') {
                        setTripleValue(
                            SetAttrStyleComponent,
                            0, index, value,  // section is ignored for associations
                            "Association attribute font style changed",
                            "Error editing association attribute font style"
                        );
                    }
                    if (childProp === 'fontFile') {
                        setTripleValue(
                            SetAttrFontComponent,
                            0, index, value,  // section is ignored for associations
                            "Association attribute font changed",
                            "Error editing association attribute font"
                        );
                    }
                    if (childProp === 'ratio') {
                        // Note: ratio might not have a generic component function, keeping the association-specific approach
                        console.warn("Ratio setting for associations not yet implemented with generic functions");
                    }
                    if (childProp === 'delete') {
                        setSingleValue(
                            (index) => RemoveAttributeFromAssociation(index),
                            index,
                            "Association attribute deleted",
                            "Error deleting association attribute"
                        );
                    }
                }
            }
        } else {
            if (property === "layer") {
                setSingleValue(
                    SetLayerComponent,
                    value,
                    "Association layer changed",
                    "Error editing association layer"
                );
            }
        }
    };

    return { handleUpdateAssociationProperty, handleAddAttributeToAssociation };
}
