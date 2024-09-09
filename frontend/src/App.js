import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import { AnalysisForm } from './components/AnalysisForm';
import { ResultsDisplay } from './components/ResultsDisplay';
import { Auth } from './components/Auth';
import { UserDashboard } from './components/UserDashboard';
import { Navbar } from './components/Navbar';

function Home({ isLoggedIn, results, setResults }) {
  return (
    <div className="px-4 py-6 sm:px-0">
      {isLoggedIn ? (
        <>
          <AnalysisForm onAnalysisComplete={setResults} />
          {results && <ResultsDisplay results={results} />}
        </>
      ) : (
        <Navigate to="/login" replace />
      )}
    </div>
  );
}

function App() {
  const [results, setResults] = useState(null);
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  const handleLogin = () => {
    setIsLoggedIn(true);
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
    localStorage.removeItem('token');
  };

  return (
    <Router>
      <div className="min-h-screen bg-gray-100">
        <Navbar isLoggedIn={isLoggedIn} onLogout={handleLogout} />
        <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
          <Routes>
            <Route path="/" element={<Home isLoggedIn={isLoggedIn} results={results} setResults={setResults} />} />
            <Route path="/login" element={<Auth onLogin={handleLogin} />} />
            <Route path="/dashboard" element={isLoggedIn ? <UserDashboard /> : <Navigate to="/login" replace />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
}

export default App;
