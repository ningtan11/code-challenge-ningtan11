import React from 'react';
import './page.css';
import SideNavbar from '../components/SideNavBar/SideNavbar';

function IncompletePage() {
  return (
    <div style={{ display: "flex" }}>
      <div>
        <SideNavbar />
      </div>

      <div class='page-content'>
        <h1>Not implemented Yet!</h1>
      </div>
    </div>
  )
}

export default IncompletePage
