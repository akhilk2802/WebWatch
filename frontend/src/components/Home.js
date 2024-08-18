import { Container, Button } from "react-bootstrap";
import { Link } from "react-router-dom";
import "../styles/css/home.css";

const Home = () => {
  return (
    <Container className="home-main">
      <h2>This is real time analytics site </h2>
      <h6> Under Construction</h6>
      <Link to="/demo">
        <Button>Demo</Button>
      </Link>
    </Container>
  );
};
export default Home;
