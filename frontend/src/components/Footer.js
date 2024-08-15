import { Navbar, Container, Nav } from "react-bootstrap";
import "../styles/css/footer.css";

const Footer = () => {
  return (
    <footer className="footer">
      <p>&copy; {new Date().getFullYear()} WebWatch. All rights reserved.</p>
    </footer>
  );
};
export default Footer;
