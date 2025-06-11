import { useState } from "react";

export function usePopupState() {
    const [showPopup, setShowPopup] = useState(false);
    const open = () => setShowPopup(true);
    const close = () => setShowPopup(false);
    return { showPopup, open, close };
}

export function useAssPopupState() {
    const [showAssPopup, setShowAssPopup] = useState(false);
    const [assStartPoint, setAssStartPoint] = useState<{ x: number, y: number } | null>(null);
    const [assEndPoint, setAssEndPoint] = useState<{ x: number, y: number } | null>(null);
    const open = () => setShowAssPopup(true);
    const close = () => {
        setShowAssPopup(false);
        setAssStartPoint(null);
        setAssEndPoint(null);
    };
    return { showAssPopup, open, close, assStartPoint, setAssStartPoint, assEndPoint, setAssEndPoint };
}
