import SideNavbar from '../components/SideNavBar/SideNavbar';
import Form from '../components/Form/Form';
import './page.css';

function FormPage() {


  return (
    <div style={{ display: "flex" }}>
      <div>
        <SideNavbar />
      </div>

      <div class='page-content'>
        <Form />
      </div>
    </div>
  );
}

export default FormPage;
