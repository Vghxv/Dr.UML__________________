import React from "react";

interface ToolbarProps {
    onGetDiagramName: () => void;
    onAddGadget: () => void;
    onShowPopup: () => void;
    onAddAss: () => void;
    diagramName?: string | null;
}

const Toolbar: React.FC<ToolbarProps> = ({
    onGetDiagramName,
    onAddGadget,
    onShowPopup,
    onAddAss,
    diagramName
}) => (
    <div
        style={{
            display: 'flex',
            alignItems: 'center',
            gap: '18px',
            margin: '24px 0 18px 0',
            padding: '14px 32px',
            background: '#23272f',
            borderRadius: '12px',
            boxShadow: '0 2px 10px rgba(0,0,0,0.13)',
            minHeight: 60,
        }}
    >
        <button
            onClick={onGetDiagramName}
            style={{
                background: '#4f8cff',
                color: 'white',
                border: 'none',
                borderRadius: '7px',
                padding: '10px 20px',
                fontWeight: 600,
                fontSize: '15px',
                cursor: 'pointer',
                boxShadow: '0 1px 4px rgba(79,140,255,0.12)',
                transition: 'background 0.2s',
            }}
            onMouseOver={e => (e.currentTarget.style.background = '#357ae8')}
            onMouseOut={e => (e.currentTarget.style.background = '#4f8cff')}
        >
            Get Diagram Name
        </button>
        {diagramName && (
            <span style={{
                color: '#fff',
                fontWeight: 500,
                fontSize: '15px',
                marginLeft: 4,
                marginRight: 8,
                letterSpacing: 0.5,
            }}>
                Diagram Name: {diagramName}
            </span>
        )}
        <button
            onClick={onAddGadget}
            style={{
                background: '#43c59e',
                color: 'white',
                border: 'none',
                borderRadius: '7px',
                padding: '10px 20px',
                fontWeight: 600,
                fontSize: '15px',
                cursor: 'pointer',
                boxShadow: '0 1px 4px rgba(67,197,158,0.12)',
                transition: 'background 0.2s',
            }}
            onMouseOver={e => (e.currentTarget.style.background = '#2ea97b')}
            onMouseOut={e => (e.currentTarget.style.background = '#43c59e')}
        >
            Add New Gadget
        </button>
        <button
            onClick={onShowPopup}
            style={{
                background: '#f7b32b',
                color: '#23272f',
                border: 'none',
                borderRadius: '7px',
                padding: '10px 20px',
                fontWeight: 600,
                fontSize: '15px',
                cursor: 'pointer',
                boxShadow: '0 1px 4px rgba(247,179,43,0.12)',
                transition: 'background 0.2s',
            }}
            onMouseOver={e => (e.currentTarget.style.background = '#e09e1b')}
            onMouseOut={e => (e.currentTarget.style.background = '#f7b32b')}
        >
            + Create Gadget (Popup)
        </button>
        <button
            onClick={onAddAss}
            style={{
                background: '#e94f37',
                color: 'white',
                border: 'none',
                borderRadius: '7px',
                padding: '10px 20px',
                fontWeight: 600,
                fontSize: '15px',
                cursor: 'pointer',
                boxShadow: '0 1px 4px rgba(233,79,55,0.12)',
                transition: 'background 0.2s',
            }}
            onMouseOver={e => (e.currentTarget.style.background = '#c0392b')}
            onMouseOut={e => (e.currentTarget.style.background = '#e94f37')}
        >
            Add Association
        </button>
    </div>
);

export default Toolbar;
