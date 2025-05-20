import { useMemo } from "react";
import { GadgetProps } from "../utils/Props";

export function useSelection(gadgets: GadgetProps[] | undefined) {
    const selectedGadgets = useMemo(
        () => (gadgets ? gadgets.filter(g => g.isSelected) : []),
        [gadgets]
    );
    const selectedGadget = selectedGadgets.length === 1 ? selectedGadgets[0] : null;
    const selectedGadgetCount = selectedGadgets.length;

    return {
        selectedGadgetCount,
        selectedGadget,
        selectedGadgets
    };
}
