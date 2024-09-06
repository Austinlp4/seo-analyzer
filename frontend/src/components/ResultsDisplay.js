import React from 'react';
import { AnalysisCharts } from './AnalysisCharts';

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
        <p><strong>Page Load Speed:</strong> {results.pageLoadSpeed.toFixed(2)} seconds</p>
        <p><strong>Mobile Friendly:</strong> {results.mobileFriendly ? 'Yes' : 'No'}</p>
        <p><strong>Responsive Design:</strong> {results.responsiveDesign ? 'Yes' : 'No'}</p>
        <p><strong>SSL Certificate:</strong> {results.sslCertificate ? 'Yes' : 'No'}</p>
        <p><strong>Meta Robots Content:</strong> {results.metaRobotsContent || 'Not specified'}</p>
      </div>
      <AnalysisCharts results={results} />
    </div>
  );
}
