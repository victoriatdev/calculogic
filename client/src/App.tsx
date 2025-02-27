import { RouterProvider } from "react-router-dom";
import { router } from "./constants/router";

export default function App() {
  return (
    <main>
      <RouterProvider router={router} />
    </main>
  );
}
