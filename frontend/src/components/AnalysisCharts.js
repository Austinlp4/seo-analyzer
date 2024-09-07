import React from 'react';
import { Bar, Doughnut, Radar } from 'react-chartjs-2';
import { Chart as ChartJS, ArcElement, Tooltip, Legend, CategoryScale, LinearScale, BarElement, Title, RadialLinearScale, PointElement, LineElement, Filler } from 'chart.js';

ChartJS.register(ArcElement, Tooltip, Legend, CategoryScale, LinearScale, BarElement, Title, RadialLinearScale, PointElement, LineElement, Filler);

export function AnalysisCharts({ results }) {
  const seoScoreData = {
    labels: ['Content', 'Performance', 'Technical'],
    datasets: [
      {
        data: [
          results.wordCount > 300 ? 30 : 15,
          results.pageLoadSpeed < 3 ? 40 : (results.pageLoadSpeed < 5 ? 20 : 10),
          (results.sslCertificate ? 10 : 0) + (results.mobileFriendly ? 10 : 0) + (results.responsiveDesign ? 10 : 0),
        ],
        backgroundColor: ['#FFA500', '#4CAF50', '#2196F3'],
      },
    ],
  };

  const performanceData = {
    labels: ['Page Load Speed', 'First Contentful Paint', 'Time to Interactive'],
    datasets: [
      {
        label: 'Performance Metrics (seconds)',
        data: [results.pageLoadSpeed, results.firstContentfulPaint, results.timeToInteractive],
        backgroundColor: 'rgba(75, 192, 192, 0.6)',
      },
    ],
  };

  const technicalSEOData = {
    labels: ['Mobile Friendly', 'Responsive Design', 'SSL Certificate', 'Meta Robots', 'H1 Usage'],
    datasets: [
      {
        label: 'Technical SEO',
        data: [
          results.mobileFriendly ? 100 : 0,
          results.responsiveDesign ? 100 : 0,
          results.sslCertificate ? 100 : 0,
          results.metaRobotsContent ? 100 : 0,
          results.h1Count > 0 ? 100 : 0,
        ],
        backgroundColor: 'rgba(54, 162, 235, 0.2)',
        borderColor: 'rgb(54, 162, 235)',
        pointBackgroundColor: 'rgb(54, 162, 235)',
        pointBorderColor: '#fff',
        pointHoverBackgroundColor: '#fff',
        pointHoverBorderColor: 'rgb(54, 162, 235)'
      }
    ]
  };

  return (
    <div className="mt-8 grid grid-cols-1 md:grid-cols-2 gap-8">
      <div className="bg-white p-4 rounded-lg shadow">
        <h3 className="text-lg font-semibold mb-4">SEO Score Composition</h3>
        <Doughnut data={seoScoreData} />
      </div>
      <div className="bg-white p-4 rounded-lg shadow">
        <h3 className="text-lg font-semibold mb-4">Performance Metrics</h3>
        <Bar data={performanceData} />
      </div>
      <div className="bg-white p-4 rounded-lg shadow md:col-span-2">
        <h3 className="text-lg font-semibold mb-4">Technical SEO Overview</h3>
        <Radar data={technicalSEOData} />
      </div>
    </div>
  );
}
