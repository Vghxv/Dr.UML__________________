import React from 'react';
import { attribute } from '../../wailsjs/go/models';

interface FontStyleButtonsProps {
    fontStyle: number;
    onStyleChange: (newStyle: number, styleType: 'B' | 'I' | 'U') => void;
}

export const FontStyleButtons: React.FC<FontStyleButtonsProps> = ({
    fontStyle,
    onStyleChange
}) => {
    const handleStyleToggle = (styleFlag: number, styleType: 'B' | 'I' | 'U') => {
        const isActive = (fontStyle & styleFlag) !== 0;
        let newStyle = fontStyle;
        if (isActive) {
            newStyle &= ~styleFlag;
        } else {
            newStyle |= styleFlag;
        }
        onStyleChange(newStyle, styleType);
    };

    return (
        <div className="flex space-x-2">
            <button
                type="button"
                onClick={() => handleStyleToggle(attribute.Textstyle.Bold, 'B')}
                className={`px-3 py-2 border rounded-md ${(fontStyle & attribute.Textstyle.Bold) !== 0
                    ? 'bg-blue-500 text-white'
                    : 'bg-white text-gray-700 border-gray-300'
                } hover:bg-blue-600 hover:text-white font-bold`}
            >
                B
            </button>
            <button
                type="button"
                onClick={() => handleStyleToggle(attribute.Textstyle.Italic, 'I')}
                className={`px-3 py-2 border rounded-md ${(fontStyle & attribute.Textstyle.Italic) !== 0
                    ? 'bg-blue-500 text-white'
                    : 'bg-white text-gray-700 border-gray-300'
                } hover:bg-blue-600 hover:text-white italic`}
            >
                I
            </button>
            <button
                type="button"
                onClick={() => handleStyleToggle(attribute.Textstyle.Underline, 'U')}
                className={`px-3 py-2 border rounded-md ${(fontStyle & attribute.Textstyle.Underline) !== 0
                    ? 'bg-blue-500 text-white'
                    : 'bg-white text-gray-700 border-gray-300'
                } hover:bg-blue-600 hover:text-white underline`}
            >
                U
            </button>
        </div>
    );
};
