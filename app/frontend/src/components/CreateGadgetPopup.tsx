import React, {useState} from "react";
import {AddGadget} from "../../wailsjs/go/umlproject/UMLProject";
import {ToPoint} from "../utils/wailsBridge";
import {component} from "../../wailsjs/go/models";

export interface GadgetPopupProps {
    isOpen: boolean;
    onCreate: (gadget: any) => void;
    onClose: () => void;
}

export const GadgetPopup: React.FC<GadgetPopupProps> = ({isOpen, onCreate, onClose}) => {
    // TODO: import type from backend
    const [formData, setFormData] = useState({
        gtype: 1,
        x: 0,
        y: 0,
        layer: 0,
        color: "#0000FF",
        header: "sample header",
    });
    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const {name, value} = e.target;
        setFormData({
            ...formData,
            [name]: name === "x" || name === "y" || name === "layer" || name === "gtype" ? parseInt(value) : value
        });
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        const {gtype, x, y, layer, color, header} = formData;
        onCreate({gtype, position: {x, y}, layer, color: color, header});
        AddGadget(gtype, ToPoint(x, y), layer, color, header).then(
            (res) => {
                console.log(res)
            },
            (err) => {
                console.log(err)
            }
        )
    };

    if (!isOpen) return null;

    return (
        <div className="fixed inset-0 bg-black bg-opacity-40 backdrop-blur-sm flex justify-center items-center z-50">
            <div className="bg-white rounded-2xl shadow-2xl p-6 w-full max-w-md">
                <h2 className="text-2xl font-semibold text-center mb-6 text-gray-800">Create Gadget</h2>
                <form onSubmit={handleSubmit} className="mb-4">
                    <div className="mb-4">
                        <label className="block text-sm font-medium text-gray-700 mb-1">Gadget type</label>
                        <select
                            name="gtype"
                            value={formData.gtype}
                            onChange={handleChange}
                            className="w-full border border-gray-300 rounded-lg p-2 focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 text-black"
                        >
                            <option value={component.GadgetType.Class}>Class</option>
                        </select>
                    </div>
                    <div className="mb-4">
                        <label className="block text-sm font-medium text-gray-700 mb-1">Name</label>
                        <input
                            type="text"
                            name="header"
                            value={formData.header}
                            onChange={handleChange}
                            className="w-full border border-gray-300 rounded-lg p-2 focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 text-black"
                        />
                    </div>
                    <div className="flex gap-4 mb-4">
                        <div className="flex-1">
                            <label className="block text-sm font-medium text-gray-700 mb-1">X</label>
                            <input
                                type="number"
                                name="x"
                                value={formData.x}
                                onChange={handleChange}
                                className="w-full border border-gray-300 rounded-lg p-2 focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 text-black"
                            />
                        </div>
                        <div className="flex-1">
                            <label className="block text-sm font-medium text-gray-700 mb-1">Y</label>
                            <input
                                type="number"
                                name="y"
                                value={formData.y}
                                onChange={handleChange}
                                className="w-full border border-gray-300 rounded-lg p-2 focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 text-black"
                            />
                        </div>
                    </div>
                    <div className="mb-4">
                        <label className="block text-sm font-medium text-gray-700 mb-1">Color</label>
                        <input
                            type="color"
                            name="color"
                            value={formData.color}
                            onChange={handleChange}
                            className="w-full h-10 p-1 rounded border border-gray-300 text-black"
                        />
                    </div>
                    <div className="flex justify-end gap-3 pt-4">
                        <button
                            type="button"
                            onClick={onClose}
                            className="px-4 py-2 rounded-lg bg-gray-300 text-gray-800 hover:bg-gray-400 transition-colors cursor-pointer"
                        >
                            Cancel
                        </button>
                        <button
                            type="submit"
                            className="px-4 py-2 rounded-lg bg-yellow-500 text-white font-semibold hover:bg-yellow-600 transition-colors cursor-pointer"
                        >
                            Create
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
};
