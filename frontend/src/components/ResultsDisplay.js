import React from 'react';
import { AnalysisCharts } from './AnalysisCharts';

export function ResultsDisplay({ results }) {
  if (!results) return null;

  const ResultItem = ({ label, value }) => (
    <div className="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
      <dt className="text-sm font-medium text-gray-500">{label}</dt>
      <dd className="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">{value}</dd>
    </div>
  );

  return (
    <div className="bg-white shadow overflow-hidden sm:rounded-lg">
      <div className="px-4 py-5 sm:px-6">
        <h3 className="text-lg leading-6 font-medium text-gray-900">Analysis Results</h3>
        <p className="mt-1 max-w-2xl text-sm text-gray-500">{results.url}</p>
      </div>
      <div className="border-t border-gray-200">
        <dl>
          <ResultItem label="Title" value={results.title} />
          <ResultItem label="Description" value={results.description} />
          <ResultItem label="Status Code" value={results.statusCode} />
          <ResultItem label="H1 Count" value={results.h1Count} />
          <ResultItem label="Word Count" value={results.wordCount} />
          <ResultItem label="Page Load Speed" value={`${results.pageLoadSpeed.toFixed(2)} seconds`} />
          <ResultItem label="Mobile Friendly" value={results.mobileFriendly ? 'Yes' : 'No'} />
          <ResultItem label="Responsive Design" value={results.responsiveDesign ? 'Yes' : 'No'} />
          <ResultItem label="SSL Certificate" value={results.sslCertificate ? 'Yes' : 'No'} />
          <ResultItem label="Meta Robots Content" value={results.metaRobotsContent || 'Not specified'} />
        </dl>
      </div>
      <AnalysisCharts results={results} />
    </div>
  );
}
