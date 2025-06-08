import {
    AddAttributeToGadget,
    RemoveAttributeFromGadget,
    SetAttrContentComponent,
    SetAttrFontComponent,
    SetAttrSizeComponent,
    SetAttrStyleComponent,
    SetColorComponent,
    SetPointComponent,
    SetLayerComponent,
} from "../../wailsjs/go/umlproject/UMLProject";
import { ToPoint } from "../utils/wailsBridge";
import { CanvasProps, GadgetProps } from "../utils/Props";
import {
    createSetSingleValue,
    createSetDoubleValue,
    createSetTripleValue,
} from "./updater";
export function useGadgetUpdater(
    selectedGadget: GadgetProps | null,
    backendData: CanvasProps | null,
    reloadBackendData: () => void
) {
    // Create bound helper functions using factory functions
    const setSingleValue = createSetSingleValue(reloadBackendData);
    const setDoubleValue = createSetDoubleValue(reloadBackendData);
    const setTripleValue = createSetTripleValue(reloadBackendData);

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
        if (property.includes(".")) {
            const [parentProp, childProp] = property.split(".");
            if (parentProp.startsWith("attributes")) {
                const matches = parentProp.match(/attributes(\d+):(\d+)/);
                if (matches && matches.length === 3) {
                    const i = parseInt(matches[1]);
                    const j = parseInt(matches[2]);
                    if (childProp === "content") {
                        setTripleValue(
                            SetAttrContentComponent,
                            i,
                            j,
                            value,
                            "Gadget content changed",
                            "Error editing gadget content"
                        );
                    }
                    if (childProp === "fontSize") {
                        setTripleValue(
                            SetAttrSizeComponent,
                            i,
                            j,
                            value,
                            "Gadget fontSize changed",
                            "Error editing gadget fontSize"
                        );
                    }
                    if (childProp === "fontStyleB") {
                        setTripleValue(
                            SetAttrStyleComponent,
                            i,
                            j,
                            value,
                            "Gadget font style (bold) changed",
                            "Error editing gadget font style (bold)"
                        );
                    }
                    if (childProp === "fontStyleI") {
                        setTripleValue(
                            SetAttrStyleComponent,
                            i,
                            j,
                            value,
                            "Gadget font style (italic) changed",
                            "Error editing gadget font style (italic)"
                        );
                    }
                    if (childProp === "fontStyleU") {
                        setTripleValue(
                            SetAttrStyleComponent,
                            i,
                            j,
                            value,
                            "Gadget font style (underline) changed",
                            "Error editing gadget font style (underline)"
                        );
                    }
                    if (childProp === "fontFile") {
                        setTripleValue(
                            SetAttrFontComponent,
                            i,
                            j,
                            value,
                            "Gadget font changed",
                            "Error editing gadget font"
                        );
                    }
                    if (childProp === "delete") {
                        setDoubleValue(
                            RemoveAttributeFromGadget,
                            i,
                            j,
                            "Attribute deleted",
                            "Error deleting attribute"
                        );
                    }
                }
            }
        } else {
            if (property === "x") {
                setSingleValue(
                    (val) => SetPointComponent(ToPoint(val, selectedGadget.y)),
                    value,
                    "Gadget moved",
                    "Error editing gadget"
                );
            }
            if (property === "y") {
                setSingleValue(
                    (val) => SetPointComponent(ToPoint(selectedGadget.x, val)),
                    value,
                    "Gadget moved",
                    "Error editing gadget"
                );
            }
            if (property === "layer") {
                setSingleValue(
                    SetLayerComponent,
                    value,
                    "layer changed",
                    "Error editing gadget"
                );
            }
            if (property === "color") {
                setSingleValue(
                    SetColorComponent,
                    value,
                    "color changed",
                    "Error editing gadget"
                );
            }
        }
    };

    return { handleUpdateGadgetProperty, handleAddAttributeToGadget };
}
