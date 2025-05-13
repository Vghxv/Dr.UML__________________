export { };
// WIP
// import React, { useState } from "react";
// import { GadgetProps } from "../utils/Props"; // 請確保這個 interface 定義存在

// const GadgetEditorPopup: React.FC = () => {
//   const [form, setForm] = useState<GadgetProps>({
//     gadgetType: "Class",
//     x: 100,
//     y: 100,
//     layer: 0,
//     height: 90,
//     width: 88,
//     color: "#cccccc",
//     header: "New Gadget",
//     header_atrributes: {
//       content: "Header",
//       height: 22,
//       width: 80,
//       fontSize: 12,
//       fontStyle: 0,
//       fontFile: "Inkfree.ttf",
//     },
//     attributes: [],
//     methods: [],
//   });

//   const handleInputChange = (field: string, value: any) => {
//     setForm({ ...form, [field]: value });
//   };

//   const handleHeaderChange = (field: string, value: any) => {
//     setForm({
//       ...form,
//       header_atrributes: { ...form.header_atrributes, [field]: value },
//     });
//   };

//   const addAttribute = () => {
//     const newAttr = {
//       content: "Attribute",
//       height: 22,
//       width: 80,
//       fontSize: 12,
//       fontStyle: 0,
//       fontFile: "Inkfree.ttf",
//     };
//     setForm({ ...form, attributes: [...form.attributes, newAttr] });
//   };

//   const addMethod = () => {
//     const newMethod = {
//       content: "Method",
//       height: 22,
//       width: 80,
//       fontSize: 12,
//       fontStyle: 0,
//       fontFile: "Inkfree.ttf",
//     };
//     setForm({ ...form, methods: [...form.methods, newMethod] });
//   };

//   return (
//     <div style={styles.overlay}>
//       <div style={styles.popup}>
//         <h2 style={styles.sectionTitle}>📦 Gadget Editor</h2>

//         <div style={styles.section}>
//           <label>Gadget Type:</label>
//           <input value={form.gadgetType} onChange={(e) => handleInputChange("gadgetType", e.target.value)} />

//           <div style={styles.row}>
//             <label>X:</label>
//             <input type="number" value={form.x} onChange={(e) => handleInputChange("x", Number(e.target.value))} />

//             <label>Y:</label>
//             <input type="number" value={form.y} onChange={(e) => handleInputChange("y", Number(e.target.value))} />
//           </div>

//           <div style={styles.row}>
//             <label>Layer:</label>
//             <input type="number" value={form.layer} onChange={(e) => handleInputChange("layer", Number(e.target.value))} />

//             <label>Color:</label>
//             <input type="color" value={form.color} onChange={(e) => handleInputChange("color", e.target.value)} />
//           </div>

//           <div style={styles.row}>
//             <label>Width:</label>
//             <input type="number" value={form.width} onChange={(e) => handleInputChange("width", Number(e.target.value))} />
//             <label>Height:</label>
//             <input type="number" value={form.height} onChange={(e) => handleInputChange("height", Number(e.target.value))} />
//           </div>

//           <label>Header Text:</label>
//           <input value={form.header} onChange={(e) => handleInputChange("header", e.target.value)} />
//         </div>

//         <div style={styles.section}>
//           <h3>🎨 Header Attributes</h3>
//           <label>Content:</label>
//           <input value={form.header_atrributes.content} onChange={(e) => handleHeaderChange("content", e.target.value)} />
//           <div style={styles.row}>
//             <label>Height:</label>
//             <input type="number" value={form.header_atrributes.height} onChange={(e) => handleHeaderChange("height", Number(e.target.value))} />
//             <label>Width:</label>
//             <input type="number" value={form.header_atrributes.width} onChange={(e) => handleHeaderChange("width", Number(e.target.value))} />
//           </div>
//           <label>Font Size:</label>
//           <input type="number" value={form.header_atrributes.fontSize} onChange={(e) => handleHeaderChange("fontSize", Number(e.target.value))} />
//           <label>Font Style:</label>
//           <input type="number" value={form.header_atrributes.fontStyle} onChange={(e) => handleHeaderChange("fontStyle", Number(e.target.value))} />
//           <label>Font File:</label>
//           <input value={form.header_atrributes.fontFile} onChange={(e) => handleHeaderChange("fontFile", e.target.value)} />
//         </div>

//         <div style={styles.section}>
//           <h3>📋 Attributes</h3>
//           <button onClick={addAttribute}>+ Add Attribute</button>
//           <ul>
//             {form.attributes.map((attr, i) => (
//               <li key={i}>{attr.content}</li>
//             ))}
//           </ul>
//         </div>

//         <div style={styles.section}>
//           <h3>🔧 Methods</h3>
//           <button onClick={addMethod}>+ Add Method</button>
//           <ul>
//             {form.methods.map((m, i) => (
//               <li key={i}>{m.content}</li>
//             ))}
//           </ul>
//         </div>
//       </div>
//     </div>
//   );
// };

// const styles: { [key: string]: React.CSSProperties } = {
//   overlay: {
//     position: "fixed",
//     top: 0,
//     left: 0,
//     width: "100vw",
//     height: "100vh",
//     backgroundColor: "rgba(0, 0, 0, 0.5)",
//     display: "flex",
//     justifyContent: "center",
//     alignItems: "center",
//     zIndex: 1000,
//   },
//   popup: {
//     backgroundColor: "#ffffff",
//     padding: "24px",
//     borderRadius: "10px",
//     width: "460px",
//     maxHeight: "90vh",
//     overflowY: "auto",
//     boxShadow: "0 4px 12px rgba(0,0,0,0.2)",
//   },
//   section: {
//     marginBottom: "20px",
//     borderBottom: "1px solid #ddd",
//     paddingBottom: "10px",
//   },
//   sectionTitle: {
//     marginBottom: "20px",
//     fontSize: "22px",
//     color: "#333",
//   },
//   row: {
//     display: "flex",
//     gap: "10px",
//     alignItems: "center",
//     marginBottom: "10px",
//   },
// };

// export default GadgetEditorPopup;
// const CreateComponentPopup: React.FC<GadgetProps> = ({ onClose }) => {


// }