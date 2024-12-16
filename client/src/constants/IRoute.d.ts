import { ReactElement, ReactNode } from "react";
import { RouteObject } from "react-router-dom";

interface IRoute {
  path: string;
  element: ReactElement;
  errorElement?: ReactElement;
  children: RouteObject[];
}
