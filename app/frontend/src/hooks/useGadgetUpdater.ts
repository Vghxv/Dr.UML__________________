import {
    AddAttributeToGadget,
    SetAttrContentGadget,
    SetAttrSizeGadget,
    SetAttrStyleGadget,
    SetColorGadget,
    SetPointGadget,
    SetSetLayerGadget
} from "../../wailsjs/go/umlproject/UMLProject";
import { ToPoint } from "../utils/wailsBridge";
import { CanvasProps, GadgetProps } from "../utils/Props";

export function useGadgetUpdater(
    selectedGadget: GadgetProps | null,
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

    // Helper function to handle setting a value with three parameters
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

    const handleAddAttributeToGadget = (section: number, content: string) => {
        if (!selectedGadget || !backendData || !backendData.gadgets) return;
        setSingleValue(
            (val) => AddAttributeToGadget(section, val),
            content,
            "Attribute added",
            "Error adding attribute"
        );
    };

    const handleUpdateGadgetProperty = (property: string, value: any) => {
        if (!selectedGadget || !backendData || !backendData.gadgets) return;
        if (property.includes('.')) {
            const [parentProp, childProp] = property.split('.');
            if (parentProp.startsWith('attributes')) {
                const matches = parentProp.match(/attributes(\d+):(\d+)/);
                if (matches && matches.length === 3) {
                    const i = parseInt(matches[1]);
                    const j = parseInt(matches[2]);
                    if (childProp === 'content') {
                        setTripleValue(
                            SetAttrContentGadget,
                            i, j, value,
                            "Gadget content changed",
                            "Error editing gadget content"
                        );
                    }
                    if (childProp === 'fontSize') {
                        setTripleValue(
                            SetAttrSizeGadget,
                            i, j, value,
                            "Gadget fontSize changed",
                            "Error editing gadget fontSize"
                        );
                    }
                    if (childProp === 'fontStyle') {
                        setTripleValue(
                            SetAttrStyleGadget,
                            i, j, value,
                            "Gadget fontStyle changed",
                            "Error editing gadget fontStyle"
                        );
                    }
                }
            }
        } else {
            if (property === "x") {
                setSingleValue(
                    (val) => SetPointGadget(ToPoint(val, selectedGadget.y)),
                    value,
                    "Gadget moved",
                    "Error editing gadget"
                );
            }
            if (property === "y") {
                setSingleValue(
                    (val) => SetPointGadget(ToPoint(selectedGadget.x, val)),
                    value,
                    "Gadget moved",
                    "Error editing gadget"
                );
            }
            if (property === "layer") {
                setSingleValue(
                    SetSetLayerGadget,
                    value,
                    "layer changed",
                    "Error editing gadget"
                );
            }
            if (property === "color") {
                setSingleValue(
                    SetColorGadget,
                    value,
                    "color changed",
                    "Error editing gadget"
                );
            }
        }
    };

    return { handleUpdateGadgetProperty, handleAddAttributeToGadget };
}
