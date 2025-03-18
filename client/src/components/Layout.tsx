import { Outlet } from "react-router-dom";
import Navbar from "./Navbar";

const Layout = () => {
  return (
    <>
      <Navbar />
      <main className="min-h-[calc(100vh-80px)] w-full flex items-center justify-center bg-(--flexoki-paper) dark:bg-(--flexoki-black)">
        <Outlet />
      </main>
    </>
  );
};

export default Layout;
