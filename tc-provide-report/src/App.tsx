import { MemoryRouter as Router, Routes, Route } from 'react-router-dom';
import { Toaster } from 'react-hot-toast';
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import './App.css'
import { AuthPage } from './pages/auth/Auth.page';
import { ReportPage } from './pages/report/Report.page';

function App() {
  return (
    <>
    <Router>
      <Routes>
        <Route path="/" element={<AuthPage />} />
        <Route path="/report" element={<ReportPage />} />
      </Routes>
    </Router>
    <Toaster />
  </>
  )
}

export default App
