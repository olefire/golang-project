import React, {useEffect, useState} from 'react';
import {useParams, useNavigate} from 'react-router-dom';
import CodeMirror from "@uiw/react-codemirror";
import {python} from "@codemirror/lang-python";
import {Button, Col, Container, Row} from 'react-bootstrap';
import {CreatePaste, GetPasteById, PasteData} from '../services/paste'
import {LintPaste, LintResult} from '../services/lint'

import '../styles/PastePage.css';


export const PastePage: React.FC = () => {
    const navigate = useNavigate();
    const [lintError, setLintError] = useState<string | null>(null);
    const [saveError, setSaveError] = useState<string | null>(null);
    const [fetchError, setFetchError] = useState<string | null>(null);


    const {id} = useParams();
    const [pasteData, setPasteData] = useState<PasteData>({
        id: "",
        paste: "",
        title: "example",
        userID: "",
    });
    const [lintResults, setLintResults] = useState<LintResult[]>([])
    const [pasteContent, setPasteContent] = React.useState(pasteData.paste);

    const onChange = React.useCallback(setPasteContent, [setPasteContent]);

    const handleSubmit = async () => {
        if (pasteContent === "") {
            setLintError("Can't lint empty paste");
            return;
        }
        await LintPaste(pasteContent).then((resp) => {
                if (resp.error !== "") {
                    setLintError(resp.error);
                    return;
                }
                setLintError(null);
                setLintResults(resp.result);
                console.log(lintResults)
            }
        )
    };

    const handleSave = async () => {
        // TODO: setPasteContent doesn't mutate pasteData.paste
        if (pasteContent === "") {
            setSaveError(`Can't save empty paste`);
        }
        pasteData.paste = pasteContent
        await CreatePaste(pasteData).then((resp) => {
                if (resp.error !== "") {
                    setSaveError(`Error saving the paste: ${resp.error}`);
                    return;
                }
                setSaveError(null);
                const pasteId = resp.result;
                console.log(`created paste with id=${pasteId}`);
                navigate(`/paste/${pasteId}`)
            }
        )
    };

    useEffect(() => {
        // There is no id for `/` route
        // Fetch paste only for `/paste/:id` route
        if (id != null) {
            GetPasteById(id).then((resp) => {
                if (resp.error !== "") {
                    setFetchError(`Error fetching the paste: ${resp.error}`)
                    return;
                }
                setFetchError(null)
                setPasteData(resp.result);
                setPasteContent(resp.result.paste)
            })
        }
    }, [id]);

    return (
        <Container className="paste-container">
            <Row>
                <Col>
                    <h1 className="page-title">Python Linter</h1>
                    {fetchError && <div className="error-message">{fetchError}</div>}
                    <CodeMirror
                        className="codeMirror__textarea"
                        theme={"dark"}
                        value={pasteContent}
                        height="700px"
                        extensions={[python()]}
                        onChange={onChange}
                    />
                    <div className="button-container">
                        <Button
                            as="button"
                            variant="primary"
                            onClick={handleSubmit}
                            className="lint-button"
                        >
                            Lint Code
                        </Button>
                        <Button
                            as="button"
                            variant="success"
                            onClick={handleSave}
                            className="save-button"
                        >
                            Save Paste
                        </Button>
                    </div>
                    {lintError && <div className="error-message">{lintError}</div>}
                    {saveError && <div className="error-message">{saveError}</div>}
                </Col>
                <Col>
                    <div className="response-container">
                        {lintResults.map((item, index) => (
                            <div key={index} className="error-message">
                                Error code {item.message} in line {item.line}
                            </div>
                        ))}
                    </div>
                </Col>
            </Row>
        </Container>
    );
};
