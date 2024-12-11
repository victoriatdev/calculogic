import Home from "../components/Home";
import AboutPage from "../components/AboutPage";
import { IRoute } from "./IRoute";
import Layout from "../components/Layout";

export const ROUTES: IRoute[] = [
  {
    path: "/",
    element: <Layout />,
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
