import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Line } from 'react-chartjs-2';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
);

export function UserDashboard() {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [seoPerformanceData, setSeoPerformanceData] = useState(null);
  const [pageSpeedData, setPageSpeedData] = useState(null);
  const [contentMetricsData, setContentMetricsData] = useState(null);

  useEffect(() => {
    const fetchAnalyses = async () => {
      try {
        const response = await axios.get('/api/user/analyses', {
          headers: { 
            Authorization: `Bearer ${localStorage.getItem('token')}`,
            'Content-Type': 'application/json'
          }
        });
        const data = response.data;
        prepareChartData(data);
        setLoading(false);
      } catch (err) {
        setError('Failed to fetch analyses');
        setLoading(false);
      }
    };

    fetchAnalyses();
  }, []);

  const prepareChartData = (data) => {
    const dates = data.map(analysis => new Date(analysis.CreatedAt).toLocaleDateString());
    const seoScores = data.map(analysis => analysis.SEOScore);
    const pageLoadSpeeds = data.map(analysis => analysis.PageLoadSpeed);
    const firstContentfulPaints = data.map(analysis => analysis.FirstContentfulPaint);
    const timeToInteractives = data.map(analysis => analysis.TimeToInteractive);
    const wordCounts = data.map(analysis => analysis.WordCount);
    const h1Counts = data.map(analysis => analysis.H1Count);

    setSeoPerformanceData({
      labels: dates,
      datasets: [{
        label: 'SEO Score',
        data: seoScores,
        borderColor: 'rgb(75, 192, 192)',
        tension: 0.1
      }]
    });

    setPageSpeedData({
      labels: dates,
      datasets: [
        {
          label: 'Page Load Speed (s)',
          data: pageLoadSpeeds,
          borderColor: 'rgb(255, 99, 132)',
          tension: 0.1
        },
        {
          label: 'First Contentful Paint (s)',
          data: firstContentfulPaints,
          borderColor: 'rgb(54, 162, 235)',
          tension: 0.1
        },
        {
          label: 'Time to Interactive (s)',
          data: timeToInteractives,
          borderColor: 'rgb(255, 206, 86)',
          tension: 0.1
        }
      ]
    });

    setContentMetricsData({
      labels: dates,
      datasets: [
        {
          label: 'Word Count',
          data: wordCounts,
          borderColor: 'rgb(153, 102, 255)',
          tension: 0.1
        },
        {
          label: 'H1 Count',
          data: h1Counts,
          borderColor: 'rgb(75, 192, 192)',
          tension: 0.1
        }
      ]
    });
  };

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div className="mt-8">
      <h2 className="text-2xl font-bold mb-4">Your SEO Performance Over Time</h2>
      
      {seoPerformanceData && (
        <div className="mb-8">
          <h3 className="text-xl font-bold mb-2">Overall SEO Score Trend</h3>
          <Line data={seoPerformanceData} />
        </div>
      )}

      {pageSpeedData && (
        <div className="mb-8">
          <h3 className="text-xl font-bold mb-2">Page Speed Metrics</h3>
          <Line data={pageSpeedData} />
        </div>
      )}

      {contentMetricsData && (
        <div className="mb-8">
          <h3 className="text-xl font-bold mb-2">Content Metrics</h3>
          <Line data={contentMetricsData} />
        </div>
      )}
    </div>
  );
}
