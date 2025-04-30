import { useCallback } from 'react';
import { dia } from '@joint/core';
import { parseBackendGadget } from '../utils/createGadget';

export function useParseBackendGadget() {
    const parseGadget = useCallback((json: string): dia.Element | null => {
        try {
            return parseBackendGadget(json);
        } catch (error) {
            console.error('Failed to parse backend gadget:', error);
            return null;
        }
    }, []);

    return parseGadget;
}
