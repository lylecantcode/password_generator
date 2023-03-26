import React from 'react';
import logo from './assets/airelogic_logo.svg';
import './App.css';
import NewBugDialog from './components/new_bug'
import GeneratePasswordDialog from './components/password_gen'
import AllBugs from './components/all_bugs'
import Box from '@mui/material/Box';


function App() {
  return (
    <div className="App">
        {/* <img src={logo} className="App-logo" alt="logo" /> */}
      <Box component="div">
        <GeneratePasswordDialog/>
      </Box>
    </div>

  );
}

export default App;
