import React, { useState } from "react";
import { GadgetProps } from "./Gadget";

interface CreateGadgetPopupProps {
  onCreate: (gadget: GadgetProps) => void;
  onCancel: () => void;
}

const CreateGadgetPopup: React.FC<CreateGadgetPopupProps> = ({ onCreate, onCancel }) => {
    const [form, setForm] = useState<GadgetProps>({
    gadgetType: "Class",
    x: 100,
    y: 100,
    layer: 0,
    height: 90,
    width: 88,
    color: "#cccccc",
    header: "New Gadget",
    header_atrributes: {
        content: "Header",
        height: 22,
        width: 80,
        fontSize: 12,
        fontStyle: 0,
        fontFile: "Inkfree.ttf",
    },
    attributes: [],
    methods: [],
    });

  const handleInputChange = (field: string, value: any) => {
    setForm({ ...form, [field]: value });
  };

  const handleCreate = () => {
    const gadget: GadgetProps = {
      ...form,
      color: form.color,
      header_atrributes: form.header_atrributes,
      attributes: form.attributes,
      methods: form.methods,
    };
    onCreate(gadget);
  };

  const addAttribute = () => {
    const newAttr = {
      content: "Attribute",
      height: 22,
      width: 80,
      fontSize: 12,
      fontStyle: 0,
      fontFile: "Inkfree.ttf",
    };
    setForm({ ...form, attributes: [...form.attributes, newAttr] });
  };

  const addMethod = () => {
    const newMethod = {
      content: "Method",
      height: 22,
      width: 80,
      fontSize: 12,
      fontStyle: 0,
      fontFile: "Inkfree.ttf",
    };
    setForm({ ...form, methods: [...form.methods, newMethod] });
  };

  return (
    <div style={styles.overlay}>
      <div style={styles.popup}>
        <h2>Create Gadget</h2>

        <label>Gadget Type:</label>
        <input value={form.gadgetType} onChange={(e) => handleInputChange("gadgetType", e.target.value)} />

        <label>X:</label>
        <input type="number" value={form.x} onChange={(e) => handleInputChange("x", Number(e.target.value))} />

        <label>Y:</label>
        <input type="number" value={form.y} onChange={(e) => handleInputChange("y", Number(e.target.value))} />

        <label>Layer:</label>
        <input type="number" value={form.layer} onChange={(e) => handleInputChange("layer", Number(e.target.value))} />

        <label>Width:</label>
        <input type="number" value={form.width} onChange={(e) => handleInputChange("width", Number(e.target.value))} />

        <label>Height:</label>
        <input type="number" value={form.height} onChange={(e) => handleInputChange("height", Number(e.target.value))} />

        <label>Color:</label>
        <input type="color" value={form.color} onChange={(e) => handleInputChange("color", e.target.value)} />

        <label>Header:</label>
        <input value={form.header} onChange={(e) => handleInputChange("header", e.target.value)} />

        <h3>Attributes</h3>
        <button onClick={addAttribute}>+ Add Attribute</button>
        <ul>
          {form.attributes.map((attr, i) => (
            <li key={i}>{attr.content}</li>
          ))}
        </ul>

        <h3>Methods</h3>
        <button onClick={addMethod}>+ Add Method</button>
        <ul>
          {form.methods.map((m, i) => (
            <li key={i}>{m.content}</li>
          ))}
        </ul>

        <div style={{ marginTop: "20px" }}>
          <button onClick={handleCreate}>Create</button>
          <button onClick={onCancel} style={{ marginLeft: "10px" }}>
            Cancel
          </button>
        </div>
      </div>
    </div>
  );
};

const styles: {
  overlay: React.CSSProperties;
  popup: React.CSSProperties;
} = {
  overlay: {
    position: "fixed",
    top: 0,
    left: 0,
    width: "100vw",
    height: "100vh",
    backgroundColor: "rgba(0,0,0,0.5)",
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    zIndex: 999,
  },
  popup: {
    backgroundColor: "#fff",
    padding: "20px",
    borderRadius: "8px",
    width: "400px",
    maxHeight: "80vh",
    overflowY: "auto",
  },
};


export default CreateGadgetPopup;
