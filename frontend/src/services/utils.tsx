export interface Result<T> {
    error: string;
    result: T;
}

export function Ok<T>(data : T) : Result<T> {
    return {
        error: "",
        result: data
    }
}

export function Fail<T>(error : string) : Result<T> {
    return {
        error: error,
        result: {} as T
    }
}

