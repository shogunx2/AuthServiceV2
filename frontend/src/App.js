import React, { useState } from "react";

const API_URL = "http://localhost:8080";

function App() {
  // Shared state
  const [userId, setUserId] = useState("");
  const [password, setPassword] = useState("");
  const [apiKey, setApiKey] = useState("");
  const [apiKeyValid, setApiKeyValid] = useState(false);
  const [result, setResult] = useState("");

  // Handlers for each action
  const handleAdd = async (e) => {
    console.log("Adding user:", userId);
    e.preventDefault();
    const res = await fetch(`${API_URL}/add`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        user_id: userId,
        password,
        api_key: apiKey,
        api_key_valid: apiKeyValid,
      }),
    });
    console.log("Response:", res);
    setResult(await res.text());
  };

  const handleRemove = async (e) => {
    e.preventDefault();
    const res = await fetch(`${API_URL}/remove`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ user_id: userId }),
    });
    setResult(await res.text());
  };

  const handleAuthenticate = async (e) => {
    e.preventDefault();
    const res = await fetch(`${API_URL}/authenticate`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ user_id: userId, password }),
    });
    setResult(await res.text());
  };

  const handleUpdatePassword = async (e) => {
    e.preventDefault();
    const res = await fetch(`${API_URL}/UpdatePassword`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ user_id: userId, password }),
    });
    setResult(await res.text());
  };

  const handleUpdateApiKey = async (e) => {
    e.preventDefault();
    const res = await fetch(`${API_URL}/UpdateAPIkey`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ user_id: userId, api_key: apiKey, api_key_valid: apiKeyValid }),
    });
    setResult(await res.text());
  };

  return (
    <div style={{ maxWidth: 500, margin: "auto" }}>
      <h2>AuthService Frontend</h2>
      <div>
        <label>User ID: </label>
        <input value={userId} onChange={e => setUserId(e.target.value)} />
      </div>
      <div>
        <label>Password: </label>
        <input value={password} onChange={e => setPassword(e.target.value)} type="password" />
      </div>
      <div>
        <label>API Key: </label>
        <input value={apiKey} onChange={e => setApiKey(e.target.value)} />
      </div>
      <div>
        <label>API Key Valid: </label>
        <input type="checkbox" checked={apiKeyValid} onChange={e => setApiKeyValid(e.target.checked)} />
      </div>
      <div style={{ margin: "10px 0" }}>
        <button onClick={handleAdd}>Add</button>
        <button onClick={handleRemove}>Remove</button>
        <button onClick={handleAuthenticate}>Authenticate</button>
        <button onClick={handleUpdatePassword}>Update Password</button>
        <button onClick={handleUpdateApiKey}>Update API Key</button>
      </div>
      <div>
        <strong>Result:</strong>
        <pre>{result}</pre>
      </div>
    </div>
  );
}

export default App;