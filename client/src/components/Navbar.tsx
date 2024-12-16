import { Link } from "react-router-dom";
import "../styles/Navbar.css";

const Navbar = () => {
  return (
    <header>
      <nav>
        <div className="title">
          <h2>
            <Link to="/" className="site-title">
              Home
            </Link>
          </h2>
        </div>
        <div className="internal-links">
          <Link to="/about">About</Link>
          <Link to="/about">About</Link>
          <Link to="/about">About</Link>
          <Link to="/about">About</Link>
        </div>
        <div className="social-links">
          <Link to="/about">About The Sites</Link>
        </div>
      </nav>
    </header>
  );
};

export default Navbar;
