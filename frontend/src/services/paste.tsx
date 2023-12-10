import axios from "axios";
import {Fail, Ok, Result} from "./utils";

export interface PasteData {
    id: string;
    paste: string;
    title: string;
    userID: string;
}


const api = "http://localhost:8080"


export const IsValid = (data: any): data is PasteData => {
    return (data as PasteData) !== undefined;
}

export const CreatePaste = async (paste: PasteData): Promise<Result<string>> => {
    try {
        if (!IsValid(paste)) {
            console.error(`Attempted to create paste with incorrect payload: ${paste}`)
            return Fail(`Attempted to create paste with incorrect payload: ${paste}`)
        }

        const response = await axios.post(`${api}/paste`, paste);
        const data = response.data.id
        if (response.status === 200) {
            console.log(`CreatePaste ok`)
            console.log(Ok(data))
            console.log(data)
            return Ok(data)
        }
        console.error(`CreatePaste request failed with status ${response.status}, response ${response.statusText}`)
        return Ok(data);
    } catch (error) {
        // @ts-ignore
        return Fail(error.response.data);
    }
}

export const GetPasteById = async (id: string): Promise<Result<PasteData>> => {
    try {
        const response = await axios.get(`${api}/paste/${id}`);
        const data = response.data.data;
        if (response.status === 200 && IsValid(data)) {
            console.log(`GetPasteById ok`);
            return Ok(data)
        }
        console.error(`GetPasteById request failed with status ${response.status}, response ${response.statusText}`);
        return response.data;
    } catch (error) {
        // @ts-ignore
        return Fail(error.response.data);
    }
}


