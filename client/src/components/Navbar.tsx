import { NavLink } from "react-router-dom";
import { FaGitlab } from "react-icons/fa";
import DarkModeToggle from "./DarkModeToggle";

interface NavLinkType {
  name: string,
  path: string
}

const navLinks : NavLinkType[] = [
    {
      name: "Sequent Calculus",
      path: "/sequent-calculus",
    },
    {
      name: "Natural Deduction",
      path: "/natural-deduction",
    },
    {
      name: "Custom",
      path: "/custom",
    },
    {
      name: "Glossary",
      path: "/glossary",
    },
  ];

const Navbar = () => {
  return (
      <header className="w-full shadow-sm px-8 border-b border-(--color-tx-normal) h-(--navbar-height) flex items-center bg-(--color-bg-primary)">
        <nav className="flex justify-between items-center w-full text-(--color-tx-normal)">
          <NavLink to="/" className="font-bold">
            Home
          </NavLink>
          <div>
            <ul className="flex items-center gap-8 list-none">
              {navLinks.map((link) => (
                <li key={link.name}>
                  <NavLink to={link.path} className={({isActive}) => isActive ? 'text-(--color-bl)' : 'text-(--color-tx-normal) hover:text-(--color-bl-hover)'}>
                    {link.name}
                  </NavLink>
                </li>
              ))}
            </ul> 
          </div>
          <div className="flex gap-8">
            <DarkModeToggle />
            <a target="_blank" href="https://git.cs.bham.ac.uk/waughamt/vrt911-project/">
              <FaGitlab />
            </a>
          </div>
       </nav>
      </header>
  );
};

export default Navbar;
