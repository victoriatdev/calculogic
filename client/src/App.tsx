import { RouterProvider } from "react-router-dom";
import { router } from "./constants/router";
import Navbar from "./components/Navbar";

export default function App() {
  return (
    <>
      <div className="App">
        <RouterProvider router={router} />
      </div>
    </>
  );
}
