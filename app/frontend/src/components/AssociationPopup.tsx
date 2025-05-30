import React, { useState } from "react";
import { component } from "../../wailsjs/go/models";

interface AssociationPopupProps {
    isOpen: boolean;
    startPoint: { x: number, y: number };
    endPoint: { x: number, y: number };
    onAdd: (assType: number) => void;
    onClose: () => void;
}

const AssociationPopup: React.FC<AssociationPopupProps> = ({
    isOpen, startPoint, endPoint, onAdd, onClose
}) => {
    const [assType, setAssType] = useState<number>(component.AssociationType.Dependency);

    if (!isOpen) return null;

    return (
        <div className="fixed inset-0 bg-black bg-opacity-40 flex justify-center items-center z-50">
            <div className="bg-white rounded-2xl shadow-2xl p-6 w-full max-w-sm">
                <h2 className="text-xl font-semibold mb-4 text-gray-800">Add Association</h2>
                <div className="mb-3">
                    <div className="text-gray-700 text-sm mb-1">Start Point: ({startPoint.x}, {startPoint.y})</div>
                    <div className="text-gray-700 text-sm mb-1">End Point: ({endPoint.x}, {endPoint.y})</div>
                </div>
                <div className="mb-4">
                    <label className="block text-sm font-medium text-gray-700 mb-1">Association Type</label>
                    <select
                        value={assType}
                        onChange={e => setAssType(Number(e.target.value))}
                        className="w-full border border-gray-300 rounded-lg p-2 text-black"
                    >
                        <option value={component.AssociationType.Dependency}>Dependency</option>
                        <option value={component.AssociationType.Composition}>Composition</option>
                        <option value={component.AssociationType.Extension}>Inheritance</option>
                        <option value={component.AssociationType.Implementation}>Realization</option>
                    </select>
                </div>
                <div className="flex justify-end gap-3 pt-2">
                    <button
                        type="button"
                        onClick={onClose}
                        className="px-4 py-2 rounded-lg bg-gray-300 text-gray-800 hover:bg-gray-400"
                    >
                        Cancel
                    </button>
                    <button
                        type="button"
                        onClick={() => onAdd(assType)}
                        className="px-4 py-2 rounded-lg bg-blue-500 text-white font-semibold hover:bg-blue-600"
                    >
                        Add
                    </button>
                </div>
            </div>
        </div>
    );
};

export default AssociationPopup;
