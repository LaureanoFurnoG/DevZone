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

function App() {

  return (
    <>
      <Routes>
        <Route path="/" element={<MainLayout />}>
          <Route path="createpost" element={<CreatePost />} />
          <Route path="releasenotes" element={<ReleaseNotes />} />
          <Route path="libraries" element={<Libraries />} />
          <Route path="dependencies" element={<Dependencies />} />
          <Route path="frameworks" element={<Frameworks />} />
          <Route path="backend" element={<Backend />} />
          <Route path="authentication" element={<Authentication />} />
        </Route>
      </Routes>
    </>
  )
}

export default App
