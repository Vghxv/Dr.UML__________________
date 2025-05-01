import React from 'react';
import { useParseBackendGadget } from '../hooks/useParseBackendGadget';

const ExampleComponent: React.FC = () => {
    const parseGadget = useParseBackendGadget();

    const handleParse = () => {
        const backendJson = `{
            "gadgetType": "Class",
            "x": 100,
            "y": 200,
            "layer": 1,
            "height": 120,
            "width": 200,
            "color": 16777215,
            "attributes": [
                [{"content": "id: Int", "height": 20, "width": 100, "fontSize": 12, "fontStyle": 0, "fontFile": ""}],
                [{"content": "+getId(): Int", "height": 20, "width": 100, "fontSize": 12, "fontStyle": 0, "fontFile": ""}]
            ]
        }`;

        const gadget = parseGadget(backendJson);
        if (gadget) {
            console.log('Parsed gadget:', gadget);
        }
    };

    return <button onClick={handleParse}>Parse Gadget</button>;
};

export default ExampleComponent;
