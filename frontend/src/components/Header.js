import { Navbar, Container, Nav } from "react-bootstrap";
import "../styles/css/header.css";

const Header = () => {
  return (
    <>
      <Navbar bg="dark" expand="lg" data-bs-theme="dark">
        <Container>
          <Navbar.Brand href="/">WebWatch</Navbar.Brand>
        </Container>
      </Navbar>
    </>
  );
};
export default Header;
