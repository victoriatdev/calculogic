import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
} from "react-router-dom";

import { ROUTES } from "./route";

const routes = ROUTES.map((route) => {
  return (
    <Route
      key={route.path}
      path={route.path}
      element={route.element}
    />
  );
});

export const router = createBrowserRouter(createRoutesFromElements(routes));
