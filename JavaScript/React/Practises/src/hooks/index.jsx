import { Link, Outlet } from 'react-router-dom';
import '../App.css';

function HooksLayout() {
  return (
    <div className='w-full bg-white relative'>
      <nav className='flex items-center justify-center w-full space-x-8 py-4 bg-gray-100 fixed'>
        <Link to="/usestate">useState</Link>
        <Link to="/useeffect">useEffect</Link>
        <Link to="/useref">useRef</Link>
        <Link to="/usecontext">useContext</Link>
        <Link to="/usememo">useMemo</Link>
        <Link to="/usecallback">useCallback</Link>
      </nav>

      <main className='container mx-auto p-8'>
        <Outlet />
      </main>
    </div>
  );
}  

export default HooksLayout;