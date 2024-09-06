import React from 'react';

export function ResultsDisplay({ results }) {
  if (!results) return null;

  return (
    <div className="mt-8">
      <h2 className="text-2xl font-bold mb-4">Analysis Results</h2>
      <pre className="bg-gray-100 p-4 rounded overflow-x-auto">
        {JSON.stringify(results, null, 2)}
      </pre>
    </div>
  );
}
