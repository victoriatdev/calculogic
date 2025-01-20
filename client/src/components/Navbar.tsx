import { Link } from "react-router-dom";
import "../styles/Navbar.css";

const Navbar = () => {
  const navbarLinks = [
    {
      name: "Home",
      redirect: "/",
    },
    {
      name: "Sequent Calculus",
      redirect: "/sequent-calculus",
    },
    {
      name: "Natural Deduction",
      redirect: "/natural-deduction",
    },
    {
      name: "Custom",
      redirect: "/custom",
    },
    {
      name: "About",
      redirect: "/about",
    },
  ];

  const navbarLinkItems = navbarLinks.map((link) => {
    <Link to={link.redirect}>{link.name}</Link>;
  });

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
        <div className="internal-links">{navbarLinkItems}</div>
        <div className="social-links">
          <Link to="/about">About The Sites</Link>
        </div>
      </nav>
    </header>
  );
};

export default Navbar;
