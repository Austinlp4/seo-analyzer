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
        <div className="min-h-screen bg-gray-100">
            <header className="bg-indigo-600 shadow">
                <div className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
                    <h1 className="text-3xl font-bold text-white">SEO Analyzer Tool</h1>
                </div>
            </header>
            <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
                {isLoggedIn ? (
                    <div className="px-4 py-6 sm:px-0">
                        <AnalysisForm onAnalysisComplete={setResults} />
                        {results && <ResultsDisplay results={results} />}
                    </div>
                ) : (
                    <div className="px-4 py-6 sm:px-0">
                        <Auth onLogin={handleLogin} />
                    </div>
                )}
            </main>
        </div>
    );
}

export default App;
