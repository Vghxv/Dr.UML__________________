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
        <div style={{
                position: 'fixed',
                top: 0,
                right: 0,
                bottom: 0,
                left: 0,
                backgroundColor: 'rgba(0, 0, 0, 0.4)',
                backdropFilter: 'blur(4px)',
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                zIndex: 50
            }}>
            <div style={{
                backgroundColor: 'white',
                borderRadius: '1rem',
                boxShadow: '0 25px 50px -12px rgba(0, 0, 0, 0.25)',
                padding: '1.5rem',
                width: '100%',
                maxWidth: '28rem'
            }}>
                <h2 style={{
                    fontSize: '1.5rem',
                    fontWeight: 600,
                    textAlign: 'center',
                    marginBottom: '1.5rem',
                    color: '#2d3748'
                }}>Create Gadget</h2>
                <form onSubmit={handleSubmit} style={{ marginBottom: '1rem' }}>
                    <div style={{ marginBottom: '1rem' }}>
                        <label style={{ 
                            display: 'block', 
                            fontSize: '0.875rem', 
                            fontWeight: 500, 
                            color: '#4a5568', 
                            marginBottom: '0.25rem' 
                        }}>ID</label>
                        <input
                            type="number"
                            name="id"
                            value={formData.id}
                            onChange={handleChange}
                            style={{
                                width: '100%',
                                border: '1px solid #d2d6dc',
                                borderRadius: '0.5rem',
                                padding: '0.5rem'
                            }}
                        />
                    </div>
                    <div style={{ marginBottom: '1rem' }}>
                        <label style={{ 
                            display: 'block', 
                            fontSize: '0.875rem', 
                            fontWeight: 500, 
                            color: '#4a5568', 
                            marginBottom: '0.25rem' 
                        }}>Name</label>
                        <input
                            type="text"
                            name="name"
                            value={formData.name}
                            onChange={handleChange}
                            style={{
                                width: '100%',
                                border: '1px solid #d2d6dc',
                                borderRadius: '0.5rem',
                                padding: '0.5rem'
                            }}
                        />
                    </div>
                    <div style={{ 
                        display: 'flex', 
                        gap: '1rem',
                        marginBottom: '1rem'
                    }}>
                        <div style={{ flex: 1 }}>
                            <label style={{ 
                                display: 'block', 
                                fontSize: '0.875rem', 
                                fontWeight: 500, 
                                color: '#4a5568', 
                                marginBottom: '0.25rem' 
                            }}>X</label>
                            <input
                                type="number"
                                name="x"
                                value={formData.x}
                                onChange={handleChange}
                                style={{
                                    width: '100%',
                                    border: '1px solid #d2d6dc',
                                    borderRadius: '0.5rem',
                                    padding: '0.5rem'
                                }}
                            />
                        </div>
                        <div style={{ flex: 1 }}>
                            <label style={{ 
                                display: 'block', 
                                fontSize: '0.875rem', 
                                fontWeight: 500, 
                                color: '#4a5568', 
                                marginBottom: '0.25rem' 
                            }}>Y</label>
                            <input
                                type="number"
                                name="y"
                                value={formData.y}
                                onChange={handleChange}
                                style={{
                                    width: '100%',
                                    border: '1px solid #d2d6dc',
                                    borderRadius: '0.5rem',
                                    padding: '0.5rem'
                                }}
                            />
                        </div>
                    </div>
                    <div style={{ marginBottom: '1rem' }}>
                        <label style={{ 
                            display: 'block', 
                            fontSize: '0.875rem', 
                            fontWeight: 500, 
                            color: '#4a5568', 
                            marginBottom: '0.25rem' 
                        }}>Color</label>
                        <input
                            type="color"
                            name="color"
                            value={formData.color}
                            onChange={handleChange}
                            style={{
                                width: '100%',
                                height: '2.5rem',
                                padding: '0.25rem'
                            }}
                        />
                    </div>
                    <div style={{ 
                        display: 'flex', 
                        justifyContent: 'flex-end', 
                        gap: '0.75rem', 
                        paddingTop: '1rem' 
                    }}>
                        <button
                            type="button"
                            onClick={onClose}
                            style={{
                                padding: '0.5rem 1rem',
                                borderRadius: '0.5rem',
                                backgroundColor: '#d1d5db',
                                color: '#1f2937',
                                border: 'none',
                                cursor: 'pointer'
                            }}
                            onMouseOver={(e) => e.currentTarget.style.backgroundColor = '#9ca3af'}
                            onMouseOut={(e) => e.currentTarget.style.backgroundColor = '#d1d5db'}
                        >
                            Cancel
                        </button>
                        <button
                            type="submit"
                            style={{
                                padding: '0.5rem 1rem',
                                borderRadius: '0.5rem',
                                backgroundColor: '#eab308',
                                color: 'white',
                                fontWeight: 600,
                                border: 'none',
                                cursor: 'pointer'
                            }}
                            onMouseOver={(e) => e.currentTarget.style.backgroundColor = '#ca8a04'}
                            onMouseOut={(e) => e.currentTarget.style.backgroundColor = '#eab308'}
                        >
                            Create
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
};
