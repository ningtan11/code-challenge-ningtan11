import React from 'react';
import './page.css';
import SideNavbar from '../components/SideNavBar/SideNavbar';

const NoPage = () => {
  return (
    <div style={{ display: "flex" }}>
      <div>
        <SideNavbar />
      </div>

      <div class='page-content'>
        <h1>Page does not exist</h1>
      </div>
    </div>
  );
};

export default NoPage;
