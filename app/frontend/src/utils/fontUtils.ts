type FontFile = {
    default: string;
};

let fontFiles: Record<string, FontFile>;
fontFiles = import.meta.glob<FontFile>('../assets/fonts/*.woff2', { eager: true });

export const getFontOptions = () => {
    return Object.keys(fontFiles).map(path => {
        const filename = path.split(/[/\\]/).pop() || '';
        return filename.replace('.woff2', '');
    });
};
