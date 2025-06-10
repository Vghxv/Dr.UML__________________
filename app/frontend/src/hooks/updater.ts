export const createSetSingleValue = (reloadBackendData: () => void) => {
    return (
        apiFunction: (value: any) => Promise<any>,
        value: any,
        successMessage: string,
        errorPrefix: string
    ) => {
        apiFunction(value)
            .then(() => {
                console.log(successMessage);
                reloadBackendData();
            })
            .catch((error) => {
                console.error(`${errorPrefix}:`, error);
            });
    };
};

export const createSetDoubleValue = (reloadBackendData: () => void) => {
    return (
        apiFunction: (param1: any, param2: any) => Promise<any>,
        param1: any,
        param2: any,
        successMessage: string,
        errorPrefix: string
    ) => {
        apiFunction(param1, param2)
            .then(() => {
                console.log(successMessage);
                reloadBackendData();
            })
            .catch((error) => {
                console.error(`${errorPrefix}:`, error);
            });
    };
};

export const createSetTripleValue = (reloadBackendData: () => void) => {
    return (
        apiFunction: (param1: any, param2: any, param3: any) => Promise<any>,
        param1: any,
        param2: any,
        param3: any,
        successMessage: string,
        errorPrefix: string
    ) => {
        apiFunction(param1, param2, param3)
            .then(() => {
                console.log(successMessage);
                reloadBackendData();
            })
            .catch((error) => {
                console.error(`${errorPrefix}:`, error);
            });
    };
};
