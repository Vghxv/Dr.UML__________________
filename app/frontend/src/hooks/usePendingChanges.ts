import { useState, useEffect, useCallback } from 'react';

interface UsePendingChangesOptions {
    onPropertyUpdate: (property: string, value: any) => void;
    dependencyArray: any[];
    formatProperty?: (property: string) => string;
}

export const usePendingChanges = ({
    onPropertyUpdate,
    dependencyArray,
    formatProperty
}: UsePendingChangesOptions) => {
    // Local state to track pending changes for each input
    const [pendingChanges, setPendingChanges] = useState<Record<string, any>>({});

    // Clear pending changes when selection changes
    useEffect(() => {
        setPendingChanges({});
    }, dependencyArray);

    // Handle input changes locally without calling API
    const handleInputChange = useCallback((property: string, value: any) => {
        setPendingChanges(prev => ({
            ...prev,
            [property]: value
        }));
    }, []);

    // Handle Enter key press to commit changes
    const handleKeyPress = useCallback((e: React.KeyboardEvent, property: string) => {
        if (e.key === 'Enter') {
            const value = pendingChanges[property];
            if (value !== undefined) {
                const finalProperty = formatProperty ? formatProperty(property) : property;
                onPropertyUpdate(finalProperty, value);
                setPendingChanges(prev => {
                    const updated = { ...prev };
                    delete updated[property];
                    return updated;
                });
            }
        }
    }, [pendingChanges, onPropertyUpdate, formatProperty]);

    // Handle blur for select/color inputs
    const handleBlur = useCallback((property: string) => {
        const value = pendingChanges[property];
        if (value !== undefined) {
            const finalProperty = formatProperty ? formatProperty(property) : property;
            onPropertyUpdate(finalProperty, value);
            setPendingChanges(prev => {
                const updated = { ...prev };
                delete updated[property];
                return updated;
            });
        }
    }, [pendingChanges, onPropertyUpdate, formatProperty]);

    // Get current value (pending change or original value)
    const getValue = useCallback((property: string, originalValue: any) => {
        return pendingChanges[property] !== undefined ? pendingChanges[property] : originalValue;
    }, [pendingChanges]);

    return {
        pendingChanges,
        handleInputChange,
        handleKeyPress,
        handleBlur,
        getValue
    };
};
