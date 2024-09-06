import React from 'react';
import { Bar, Pie } from 'react-chartjs-2';
import { Chart as ChartJS, ArcElement, Tooltip, Legend, CategoryScale, LinearScale, BarElement, Title } from 'chart.js';

ChartJS.register(ArcElement, Tooltip, Legend, CategoryScale, LinearScale, BarElement, Title);

export function AnalysisCharts({ results }) {
  const countMetricsData = {
    labels: ['Word Count', 'H1 Count'],
    datasets: [
      {
        label: 'Count',
        data: [results.wordCount, results.h1Count],
        backgroundColor: ['rgba(75, 192, 192, 0.6)', 'rgba(255, 99, 132, 0.6)'],
      },
    ],
  };

  const seoScoreComposition = {
    labels: ['Content', 'Performance', 'Technical'],
    datasets: [
      {
        data: [
          results.wordCount > 300 ? 30 : 15,
          results.pageLoadSpeed < 3 ? 40 : (results.pageLoadSpeed < 5 ? 20 : 10),
          (results.sslCertificate ? 10 : 0) + (results.mobileFriendly ? 10 : 0) + (results.responsiveDesign ? 10 : 0),
        ],
        backgroundColor: ['rgba(255, 206, 86, 0.6)', 'rgba(75, 192, 192, 0.6)', 'rgba(153, 102, 255, 0.6)'],
      },
    ],
  };

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mt-8">
      <div>
        <h3 className="text-xl font-semibold mb-2">Content Metrics</h3>
        <Bar
          data={countMetricsData}
          options={{
            scales: {
              y: {
                beginAtZero: true,
              },
            },
          }}
        />
      </div>
      <div>
        <h3 className="text-xl font-semibold mb-2">SEO Score Composition</h3>
        <Pie data={seoScoreComposition} />
      </div>
    </div>
  );
}
