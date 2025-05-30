/// <reference types="vite/client" />

// Extend ImportMeta interface to include Vite-specific APIs
interface ImportMeta {
    readonly glob: <T = any>(pattern: string, options?: { eager?: boolean }) => Record<string, T>;
}
