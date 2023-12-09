import React, {useState} from 'react';
import CodeMirror from "@uiw/react-codemirror";
import {python} from "@codemirror/lang-python";

interface LintResult {
    type: string;
    module: string;
    obj: string;
    line: number;
    column: number;
    endLine: number;
    endColumn: number;
    path: string;
    symbol: string;
    message: string;
    'message-id': string;
}

export const TestLintPage: React.FC = () => {
    const [value, setValue] = React.useState("print('hello goyland')");
    const [response, setResponse] = useState<LintResult[]>([])

    const onChange = React.useCallback((value: string) => {
        console.log('value:', value);
        setValue(value)
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
                <button onClick={handleSubmit}>Submit</button>
                <div>{response.map(item => (
                    <li key={item.module}>Error {item["message-id"]} in {item.line}:{item.column}: {item.message}</li>
                ))}</div>
            </div>
        </>
    );
};

