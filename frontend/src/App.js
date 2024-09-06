import React, { useState } from 'react';
import { AnalysisForm } from './components/AnalysisForm';
import { ResultsDisplay } from './components/ResultsDisplay';

function App() {
    const [results, setResults] = useState(null);

    return (
        <div className="container mx-auto px-4 py-8">
            <h1 className="text-3xl font-bold mb-8">SEO Analyzer Tool</h1>
            <AnalysisForm onAnalysisComplete={setResults} />
            <ResultsDisplay results={results} />
        </div>
    );
}

export default App;
