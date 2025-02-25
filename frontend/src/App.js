import React, { useState, useEffect } from 'react';

function App() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    fetch('http://localhost:5000/api/hello')
      .then(response => response.json())
      .then(data => setMessage(data.text))
      .catch(err => console.error('Error fetching message:', err));
  }, []);

  return (
    <div className="App">
      <h1>{message ? message : 'Loading...'}</h1>
    </div>
  );
}

export default App;
