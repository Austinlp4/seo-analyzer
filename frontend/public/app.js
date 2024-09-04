const { useState, useEffect } = React;

function App() {
    const [page, setPage] = useState('home');

    useEffect(() => {
        const path = window.location.pathname;
        if (path === '/') {
            setPage('home');
        } else if (path === '/analyze') {
            setPage('analyze');
        }
    }, []);

    return (
        <div className="container mx-auto px-4">
            {page === 'home' && <Home />}
            {page === 'analyze' && <Analyze />}
        </div>
    );
}

function Home() {
    return (
        <div className="text-center mt-10">
            <h1 className="text-4xl font-bold mb-4">SEO Analyzer Tool</h1>
            <a href="/analyze" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                Start Analysis
            </a>
        </div>
    );
}

function Analyze() {
    const [url, setUrl] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        console.log('Analyzing URL:', url);
    };

    return (
        <div className="mt-10">
            <h1 className="text-3xl font-bold mb-4">Analyze Website</h1>
            <form onSubmit={handleSubmit} className="flex flex-col items-center">
                <input
                    type="url"
                    value={url}
                    onChange={(e) => setUrl(e.target.value)}
                    placeholder="Enter website URL"
                    required
                    className="w-full max-w-md px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 mb-4"
                />
                <button type="submit" className="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded">
                    Analyze
                </button>
            </form>
        </div>
    );
}

ReactDOM.render(<App />, document.getElementById('root'));
