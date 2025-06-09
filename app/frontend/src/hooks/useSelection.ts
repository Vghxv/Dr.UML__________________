import { useMemo, useRef } from "react";
import { AssociationProps, GadgetProps } from "../utils/Props";

export function useSelection(
    components: (GadgetProps | AssociationProps)[] | undefined
) {
    const previousSelectionRef = useRef<string>('');
    
    const selectedComponents = useMemo(
        () => (components ? components.filter((g) => g.isSelected) : []),
        [components]
    );
    
    const selectedComponent =
        selectedComponents.length === 1 ? selectedComponents[0] : null;
    const selectedComponentCount = selectedComponents.length;
    
    // Only log when the selection actually changes
    // Create a simple hash of selected components for comparison
    const currentSelectionHash = useMemo(() => {
        return selectedComponents
            .map(comp => {
                if ('gadgetType' in comp) {
                    return `g-${comp.x}-${comp.y}-${comp.gadgetType}`;
                } else {
                    return `a-${comp.startX}-${comp.startY}-${comp.endX}-${comp.endY}-${comp.assType}`;
                }
            })
            .sort()
            .join('|');
    }, [selectedComponents]);
    
    // Update the ref and log only when there's a real change
    if (currentSelectionHash !== previousSelectionRef.current) {
        previousSelectionRef.current = currentSelectionHash;
        console.log("Selected components:", selectedComponents);
    }
    
    return {
        selectedComponentCount,
        selectedComponent,
    };
}
