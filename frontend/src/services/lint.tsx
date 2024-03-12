import axios from "axios";
import {Fail, Ok, Result} from "./utils";

export interface LintResult {
    message: string;
    line: number;
}

const api = "http://localhost:8080"


const IsValid = (result: any[]): result is LintResult[] => {
    return (result as LintResult[]) !== undefined;
}

export const LintPaste = async (paste: string): Promise<Result<LintResult[]>> => {
    try {
        const response = await axios.post(`${api}/lint`, paste);
        const data = response.data;
        if (response.status === 200 && IsValid(data)) {
            console.log(`LintPaste ok`)
            return Ok(data);
        }
        console.error(`LintPaste request failed with status ${response.status}, response ${response.statusText}`);
        return Ok([]);
    } catch (error) {
        // @ts-ignore
        return Fail(error.response.data);
    }
}
