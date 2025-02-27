import { Outlet } from "react-router-dom";
import Navbar from "./Navbar";

const Layout = () => {
  return (
    <>
      <Navbar />
      <main className="min-h-screen w-full flex items-center justify-center bg-(--flexoki-paper) dark:bg-(--flexoki-black)">
        <Outlet />
      </main>
    </>
  );
};

export default Layout;
