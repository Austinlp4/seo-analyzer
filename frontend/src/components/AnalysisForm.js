import React, { useState } from 'react';
import axios from 'axios';

export function AnalysisForm({ onAnalysisComplete }) {
  const [url, setUrl] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');

    if (!url) {
      setError('Please enter a URL');
      setIsLoading(false);
      return;
    }

    try {
      const response = await axios.post('/api/analyze', { url });
      onAnalysisComplete(response.data);
    } catch (err) {
      if (err.response) {
        setError(err.response.data || 'An error occurred while analyzing the URL.');
      } else if (err.request) {
        setError('No response received from the server. Please try again.');
      } else {
        setError('An unexpected error occurred. Please try again.');
      }
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <input
        type="url"
        value={url}
        onChange={(e) => setUrl(e.target.value)}
        placeholder="Enter URL to analyze"
        required
        className="w-full px-4 py-2 border rounded"
      />
      <button
        type="submit"
        disabled={isLoading}
        className="w-full px-4 py-2 text-white bg-blue-500 rounded hover:bg-blue-600 disabled:bg-blue-300"
      >
        {isLoading ? 'Analyzing...' : 'Analyze'}
      </button>
      {error && <p className="text-red-500">{error}</p>}
    </form>
  );
}
