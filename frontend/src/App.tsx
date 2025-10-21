import {useState} from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import type {ApiResponse} from "./common.ts";

function App() {
  const [count, setCount] = useState(0)
  const [currentTime, setCurrentTime] = useState('?????')

    const fetchTime  = async () => {
        try {
            const response = await fetch('/api/time');
            const result: ApiResponse<string> = await response.json();
            if (result.data) {
                setCurrentTime(result.data);
            }
        } catch (error) {
            console.error('Error fetching users:', error);
        }
    }

    // useEffect(() => {
    //     fetchTime();
    // }, []);
    
  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
        <div className="card">
            <button onClick={() => { fetchTime() }}>
                {currentTime}
            </button>
            <p>
                Click this button ↑ to get current time
            </p>
        </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
