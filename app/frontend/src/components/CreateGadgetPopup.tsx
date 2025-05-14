import React, { useState } from "react";

interface GadgetPopupProps {
    isOpen: boolean;
    onCreate: (gadget: any) => void;
    onClose: () => void;
}

export const GadgetPopup: React.FC<GadgetPopupProps> = ({ isOpen, onCreate, onClose }) => {
    const [formData, setFormData] = useState({
        id: 0,
        name: "",
        x: 100,
        y: 100,
        color: "#FF0000",
    });

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: name === "x" || name === "y" || name === "id" ? parseInt(value) : value });
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        const { id, name, x, y, color } = formData;
        onCreate({ id, name, position: { x, y }, color });
    };

    if (!isOpen) return null;

    return (
        <div className="fixed inset-0 bg-black bg-opacity-40 backdrop-blur-sm flex justify-center items-center z-50">
            <div className="bg-white rounded-2xl shadow-2xl p-6 w-full max-w-md">
                <h2 className="text-2xl font-semibold text-center mb-6 text-gray-800">Create Gadget</h2>
                <form onSubmit={handleSubmit} className="space-y-4">
                    <div>
                        <label className="block text-sm font-medium text-gray-700 mb-1">ID</label>
                        <input
                            type="number"
                            name="id"
                            value={formData.id}
                            onChange={handleChange}
                            className="w-full border border-gray-300 rounded-lg p-2"
                        />
                    </div>
                    <div>
                        <label className="block text-sm font-medium text-gray-700 mb-1">Name</label>
                        <input
                            type="text"
                            name="name"
                            value={formData.name}
                            onChange={handleChange}
                            className="w-full border border-gray-300 rounded-lg p-2"
                        />
                    </div>
                    <div className="flex space-x-4">
                        <div className="flex-1">
                            <label className="block text-sm font-medium text-gray-700 mb-1">X</label>
                            <input
                                type="number"
                                name="x"
                                value={formData.x}
                                onChange={handleChange}
                                className="w-full border border-gray-300 rounded-lg p-2"
                            />
                        </div>
                        <div className="flex-1">
                            <label className="block text-sm font-medium text-gray-700 mb-1">Y</label>
                            <input
                                type="number"
                                name="y"
                                value={formData.y}
                                onChange={handleChange}
                                className="w-full border border-gray-300 rounded-lg p-2"
                            />
                        </div>
                    </div>
                    <div>
                        <label className="block text-sm font-medium text-gray-700 mb-1">Color</label>
                        <input
                            type="color"
                            name="color"
                            value={formData.color}
                            onChange={handleChange}
                            className="w-full h-10 p-1"
                        />
                    </div>
                    <div className="flex justify-end space-x-3 pt-4">
                        <button
                            type="button"
                            onClick={onClose}
                            className="px-4 py-2 rounded-lg bg-gray-300 hover:bg-gray-400 text-gray-800"
                        >
                            Cancel
                        </button>
                        <button
                            type="submit"
                            className="px-4 py-2 rounded-lg bg-yellow-500 hover:bg-yellow-600 text-white font-semibold"
                        >
                            Create
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
};
