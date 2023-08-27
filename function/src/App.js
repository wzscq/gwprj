import {HashRouter,Routes,Route} from "react-router-dom";
import Function from './pages/Function';

import './App.css';

function App() {
  return (
    <div className="App">
      <HashRouter>
        <Routes>
          <Route path="/" exact={true} element={<Function/>} />
        </Routes>
      </HashRouter>
    </div>
  );
}

export default App;
