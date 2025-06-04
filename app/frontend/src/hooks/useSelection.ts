import {useMemo} from "react";
import {AssociationProps, GadgetProps} from "../utils/Props";

export function useSelection(componenets: (GadgetProps | AssociationProps)[] | undefined) {
    const selectedComponents = useMemo(
        () => (componenets ? componenets.filter(g => g.isSelected) : []),
        [componenets]
    );
    const selectedComponent = selectedComponents.length === 1 ? selectedComponents[0] : null;
    const selectedComponentCount = selectedComponents.length;

    return {
        selectedComponentCount,
        selectedComponent,
        selectedComponents
    };
}
