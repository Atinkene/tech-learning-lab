import { Routes, Route } from 'react-router-dom';
import HooksLayout from './hooks/index.jsx';
import ReactUseState from './hooks/ReactUseState/index.jsx';
import UseEffect from './hooks/UseEffect/index.jsx';
import UseContext from './hooks/UseContext/index.jsx';
import UseCallback from './hooks/UseCallback/index.jsx';

function App() {
  return (
    <Routes>
      <Route path="/" element={<HooksLayout />}>
        <Route index element={<ReactUseState />} />
        <Route path="usestate" element={<ReactUseState />} />
        <Route path="useeffect" element={<UseEffect />} />
        <Route path="usecontext" element={<UseContext />} />
        <Route path="usecallback" element={<UseCallback />} />
      </Route>
    </Routes>
  );
}

export default App;