import Editor from '@monaco-editor/react';
import {useEffect, useState} from "react";

export default function CodeEditor(props) {
    const { snippetCode, snippetLanguage, loading } = props;

    useEffect(() => {
        setCode(snippetCode);
        setLanguage(snippetLanguage);

    }, [snippetCode, snippetLanguage, loading]);

    const [code, setCode] = useState('');
    const [language, setLanguage] = useState('');

    const handleEditorDidMount = (editor, monaco) => {
        monaco.languages.typescript.javascriptDefaults.setDiagnosticsOptions({
            noSemanticValidation: false,
            noSyntaxValidation: false
        });

        monaco.languages.typescript.javascriptDefaults.setCompilerOptions({
            target: monaco.languages.typescript.ScriptTarget.ES6,
            allowNonTsExtensions: true
        });
    }

    return (
        <>
            <Editor
            height="70vh"
            width="100%"
            language={language}
            value={code}
            onChange={(value) => setCode(value)}
            theme='vs-dark'
            onMount={handleEditorDidMount}
            loading={loading}
            options={{
                fontSize: 16,
                wordWrap: 'on',
                scrollBehavior: 'smooth',
                scrollBeyondLastLine: true,
            }}/>
        </>
    )
}