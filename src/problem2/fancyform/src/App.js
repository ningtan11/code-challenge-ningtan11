import { Routes , Route } from "react-router-dom";
import './App.css';
import FormPage from './pages/FormPage';
import NoPage from './pages/NoPage';
import IncompletePage from './pages/IncompletePage';

function App() {
  return (
      <Routes> 
          <Route path="/" element={<FormPage /> } /> 
          <Route path="prices" element={<IncompletePage />}>
            <Route path="token" element={<IncompletePage />}/>
            <Route path="nfts" element={<IncompletePage />}/>
          </Route>
          <Route path="app" element={<IncompletePage />}>
            <Route path="vote" element={<IncompletePage />}/>
            <Route path="pools" element={<IncompletePage />}/>
            <Route path="analytics" element={<IncompletePage />}/>
          </Route>
          <Route path="connect" element={<IncompletePage />}>
            <Route path="uniswap" element={<IncompletePage />}/>
            <Route path="metamask" element={<IncompletePage />}/>
            <Route path="walletconnect" element={<IncompletePage />}/>
            <Route path="coinbase" element={<IncompletePage />}/>
          </Route>
          <Route path="aboutUs" element={<IncompletePage />}/>
          <Route path="faq" element={<IncompletePage />}/>
          <Route path="*" element={<NoPage/> } />
       </Routes> 
  );
}

export default App;
