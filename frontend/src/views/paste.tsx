import React, {useEffect, useState} from 'react';
import {useParams} from 'react-router-dom';
import axios from 'axios';
import CodeMirror from "@uiw/react-codemirror";
import {python} from "@codemirror/lang-python";
import { Button } from 'react-bootstrap';


interface PasteData {
    id: string;
    paste: string;
    title: string;
    userID: string;
}

interface LintResult {
    message: string;
    line: number;
}

export const PastePage: React.FC = () => {
    const {id} = useParams<{ id: string }>();
    const apiURL = `http://localhost:8080/paste/${id}`;
    const [pasteData, setPasteData] = useState<PasteData | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get(apiURL);
                setPasteData(response.data.data);
                setValue(response.data.data.paste)
                console.log(response)
            } catch (error) {
                console.error(error);
            }
        };

        fetchData().then();
    }, [apiURL, id]);
    // eslint-disable-next-line react-hooks/rules-of-hooks
    const [response, setResponse] = useState<LintResult[]>([])
    const [value, setValue] = React.useState(pasteData?.paste);
    // eslint-disable-next-line react-hooks/rules-of-hooks
    const onChange = React.useCallback((value: string) => {
        setValue(value)
        console.log('value:', value);
    }, []);

    const handleSubmit = async () => {
        try {
            const res = await fetch('http://localhost:8080/lint', {
                method: 'POST',
                headers: {
                    'Content-Type': 'text/plain'
                },
                body: value
            });
            const data = await res.json();
            console.log(data)
            setResponse(data)

        } catch (error) {
            console.error(error);
        }
    };
    let cm = <CodeMirror
        className={"codeMirror__textarea"}
        theme={"dark"}
        value={value}
        height="400px"
        extensions={[python()]}
        onChange={onChange}
    />;

    return (
        <>
            <div>
                <h1>
                    Python linter
                </h1>
                {cm}
                <Button as="button" variant="primary" onClick={handleSubmit}>Lint code</Button>
                <div>{response.map(item => (
                    <li>Error code {item.message} in line {item.line}</li>
                ))}</div>
            </div>
        </>
    );
};
