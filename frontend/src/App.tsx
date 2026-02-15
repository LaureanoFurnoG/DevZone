import './App.css'
import { Routes, Route } from "react-router-dom";
import MainLayout from './MainLayout';
import CreatePost from './pages/CreatePost/CreatePost';
import ReleaseNotes from './pages/ReleaseNotes/ReleaseNotes';
import Libraries from './pages/Libraries/Libraries';
import Dependencies from './pages/Dependencies/Dependencies';
import Backend from './pages/Backend/Backend';
import Authentication from './pages/Authentication/Authentication';
import Frameworks from './pages/Frameworks/Frameworks';
import { useAuth } from './Auth/useAuth'
import { Navigate } from 'react-router-dom'
import { ConfigProvider, theme } from "antd"
import Home from './pages/Home/Home';

const ProtectedRoute = ({ children }: { children: React.ReactNode }) => {
    const { isAuthenticated, isLoading } = useAuth()
    
    if (isLoading) return <div>Loading...</div>
    if (!isAuthenticated) return <Navigate to="/" replace />
    
    return <>{children}</>
}

function App() {
    return (
        <>
          <ConfigProvider
            theme={{
              algorithm: theme.darkAlgorithm,
            }}
          >
            <Routes>
                <Route path="/" element={<MainLayout />}>
                    <Route index element={<Navigate to="/home" replace />} />
                    <Route path="home" element={<Home />} />
                    <Route path="createpost" element={<ProtectedRoute><CreatePost /></ProtectedRoute>} />
                    <Route path="releasenotes" element={<ReleaseNotes />} />
                    <Route path="libraries" element={<Libraries />} />
                    <Route path="dependencies" element={<Dependencies />} />
                    <Route path="frameworks" element={<Frameworks />} />
                    <Route path="backend" element={<Backend />} />
                    <Route path="authentication" element={<Authentication />} />
                </Route>
            </Routes>
          </ConfigProvider>
        </>
    )
}

export default App
