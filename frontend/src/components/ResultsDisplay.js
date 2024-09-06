import React from 'react';

export function ResultsDisplay({ results }) {
  if (!results) return null;

  return (
    <div className="mt-8">
      <h2 className="text-2xl font-bold mb-4">Analysis Results</h2>
      <div className="bg-gray-100 p-4 rounded">
        <p><strong>URL:</strong> {results.url}</p>
        <p><strong>Title:</strong> {results.title}</p>
        <p><strong>Description:</strong> {results.description}</p>
        <p><strong>Status Code:</strong> {results.statusCode}</p>
        <p><strong>H1 Count:</strong> {results.h1Count}</p>
        <p><strong>Word Count:</strong> {results.wordCount}</p>
      </div>
    </div>
  );
}
