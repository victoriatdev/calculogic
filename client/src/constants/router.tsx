import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
} from "react-router-dom";

import Layout from "../components/Layout";
import Home from "../components/Home";
import ErrorPage from "../components/ErrorPage";
import AboutPage from "../components/AboutPage";

import { ROUTES } from "./route";

const routes = ROUTES.map((route) => {
  return (
    <Route
      key={route.path}
      path={route.path}
      errorElement={route.errorElement}
      element={route.element}
      children={route.children}
    />
  );
});

export const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: "",
        element: <Home />,
      },
      {
        path: "about",
        element: <AboutPage />,
      },
    ],
  },
]);
//export const router = createBrowserRouter(createRoutesFromElements(routes));
