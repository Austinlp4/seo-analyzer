import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import axios from 'axios';
import AnalysisForm from './AnalysisForm';

jest.mock('axios');

describe('AnalysisForm', () => {
  test('renders input and submit button', () => {
    render(<AnalysisForm />);
    expect(screen.getByPlaceholderText('Enter URL to analyze')).toBeInTheDocument();
    expect(screen.getByRole('button', { name: /analyze/i })).toBeInTheDocument();
  });

  test('handles form submission', async () => {
    axios.post.mockResolvedValue({ data: { url: 'https://example.com' } });
    render(<AnalysisForm onAnalysisComplete={() => {}} />);
    
    fireEvent.change(screen.getByPlaceholderText('Enter URL to analyze'), {
      target: { value: 'https://example.com' },
    });
    fireEvent.click(screen.getByRole('button', { name: /analyze/i }));

    await waitFor(() => {
      expect(axios.post).toHaveBeenCalledWith('/api/analyze', { url: 'https://example.com' });
    });
  });

  test('displays error message for invalid URL', async () => {
    render(<AnalysisForm onAnalysisComplete={() => {}} />);
    
    fireEvent.change(screen.getByPlaceholderText('Enter URL to analyze'), {
      target: { value: 'invalid-url' },
    });
    fireEvent.click(screen.getByRole('button', { name: /analyze/i }));

    await waitFor(() => {
      expect(screen.getByText(/please enter a valid url/i)).toBeInTheDocument();
    });
  });
});
