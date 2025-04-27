class SessionManager {
    private static instance: SessionManager;
    private sessionData: Record<string, any>;

    private constructor() {
        this.sessionData = {};
    }

    static getInstance(): SessionManager {
        if (!SessionManager.instance) {
            SessionManager.instance = new SessionManager();
        }
        return SessionManager.instance;
    }

    setItem(key: string, value: any): void {
        this.sessionData[key] = value;
    }

    getItem<T>(key: string): T | null {
        return this.sessionData[key] ?? null;
    }

    removeItem(key: string): void {
        delete this.sessionData[key];
    }

    clear(): void {
        this.sessionData = {};
    }
}

export default SessionManager;
