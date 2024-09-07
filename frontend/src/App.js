import React, { useState } from 'react';
import { AnalysisForm } from './components/AnalysisForm';
import { ResultsDisplay } from './components/ResultsDisplay';
import { Auth } from './components/Auth';

function App() {
    const [results, setResults] = useState(null);
    const [isLoggedIn, setIsLoggedIn] = useState(false);

    const handleLogin = () => {
        setIsLoggedIn(true);
    };

    return (
        <div className="container mx-auto px-4 py-8">
            <h1 className="text-3xl font-bold mb-8">SEO Analyzer Tool</h1>
            {isLoggedIn ? (
                <>
                    <AnalysisForm onAnalysisComplete={setResults} />
                    <ResultsDisplay results={results} />
                </>
            ) : (
                <Auth onLogin={handleLogin} />
            )}
        </div>
    );
}

export default App;
