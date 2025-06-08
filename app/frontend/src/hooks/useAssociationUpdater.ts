import {
    AddAttributeToAssociation,
    RemoveAttributeFromAssociation,
    SetAttrContentComponent,
    SetAttrSizeComponent,
    SetAttrStyleComponent,
    SetAttrFontComponent,
    SetLayerComponent
} from "../../wailsjs/go/umlproject/UMLProject";
import {CanvasProps, AssociationProps} from "../utils/Props";

export function useAssociationUpdater(
    selectedAssociation: AssociationProps | null,
    backendData: CanvasProps | null,
    reloadBackendData: () => void
) {
    // Helper function to handle setting a single value with a promise
    const setSingleValue = (
        apiFunction: (value: any) => Promise<any>,
        value: any,
        successMessage: string,
        errorPrefix: string
    ) => {
        apiFunction(value).then(
            () => {
                console.log(successMessage);
                reloadBackendData();
            }
        ).catch((error) => {
            console.error(`${errorPrefix}:`, error);
        });
    };

    // Helper function to handle setting a value with two parameters
    const setDoubleValue = (
        apiFunction: (param1: any, param2: any) => Promise<any>,
        param1: any,
        param2: any,
        successMessage: string,
        errorPrefix: string
    ) => {
        apiFunction(param1, param2).then(
            () => {
                console.log(successMessage);
                reloadBackendData();
            }
        ).catch((error) => {
            console.error(`${errorPrefix}:`, error);
        });
    };    // Helper function to handle setting a value with three parameters (componentId, index, value)
    const setTripleValue = (
        apiFunction: (param1: any, param2: any, param3: any) => Promise<any>,
        param1: any,
        param2: any,
        param3: any,
        successMessage: string,
        errorPrefix: string
    ) => {
        apiFunction(param1, param2, param3).then(
            () => {
                console.log(successMessage);
                reloadBackendData();
            }
        ).catch((error) => {
            console.error(`${errorPrefix}:`, error);
        });
    };

    const handleAddAttributeToAssociation = (ratio: number, content: string) => {
        if (!selectedAssociation || !backendData) return;
        setDoubleValue(
            AddAttributeToAssociation,
            ratio, content,
            "Association attribute added",
            "Error adding association attribute"
        );
    };

    const handleUpdateAssociationProperty = (property: string, value: any) => {
        if (!selectedAssociation || !backendData) return;
        
        if (property.includes('.')) {
            const [parentProp, childProp] = property.split('.');
            if (parentProp.startsWith('attributes')) {
                const matches = parentProp.match(/attributes:(\d+)/);
                if (matches && matches.length === 2) {
                    const index = parseInt(matches[1]);                    if (childProp === 'content') {
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
        } else {            // Handle direct properties like layer, etc.
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

    return {handleUpdateAssociationProperty, handleAddAttributeToAssociation};
}
