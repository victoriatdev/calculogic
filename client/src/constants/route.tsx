import Home from "../components/Home";
import AboutPage from "../components/AboutPage";
import ErrorPage from "../components/ErrorPage";
import { IRoute } from "./IRoute";
import Layout from "../components/Layout";

export const ROUTES: any = [
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
];
