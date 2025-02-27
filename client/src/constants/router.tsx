import {
  createBrowserRouter,
} from "react-router-dom";

import Layout from "../components/Layout";
import Home from "../components/Home";
import ErrorPage from "../components/ErrorPage";
import AboutPage from "../components/AboutPage";
import GlossaryPage from "../components/GlossaryPage";
import NaturalDeduction from "../components/NaturalDeduction";
import SequentCalculus from "../components/SequentCalculus";
import CustomRuleset from "../components/CustomRuleset";

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
      {
        path: "glossary",
        element: <GlossaryPage />,
      },
      {
        path: "natural-deduction",
        element: <NaturalDeduction />,
      },
      {
        path: "sequent-calculus",
        element: <SequentCalculus />,
      },
      {
        path: "custom",
        element: <CustomRuleset />,
      }
    ],
  },
]);
